package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	chi "github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/codercollo/hotel-system/backend/config"
	"github.com/codercollo/hotel-system/backend/database"
	"github.com/codercollo/hotel-system/backend/events"
	evhandlers "github.com/codercollo/hotel-system/backend/events/handlers"
	"github.com/codercollo/hotel-system/backend/internal/admin"
	"github.com/codercollo/hotel-system/backend/internal/auth"
	"github.com/codercollo/hotel-system/backend/internal/items"
	"github.com/codercollo/hotel-system/backend/internal/notifications"
	notifchannels "github.com/codercollo/hotel-system/backend/internal/notifications/channels"
	"github.com/codercollo/hotel-system/backend/internal/orders"
	"github.com/codercollo/hotel-system/backend/internal/payments"
	"github.com/codercollo/hotel-system/backend/internal/payments/providers"
	"github.com/codercollo/hotel-system/backend/internal/roles"
	"github.com/codercollo/hotel-system/backend/internal/search"
	"github.com/codercollo/hotel-system/backend/internal/uploads"
	uploadstorage "github.com/codercollo/hotel-system/backend/internal/uploads/storage"
	"github.com/codercollo/hotel-system/backend/internal/users"
	ws "github.com/codercollo/hotel-system/backend/internal/websocket"
	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/cache"
	"github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

// authUserAdapter lives in main (Composition Root) — safely imports both
// auth and users packages without creating an import cycle.
type authUserAdapter struct {
	repo users.Repository
}

func (a *authUserAdapter) FindByEmail(ctx context.Context, email string) (auth.AuthUser, error) {
	u, err := a.repo.FindByEmail(ctx, email)
	if err != nil {
		return auth.AuthUser{}, err
	}
	return auth.AuthUser{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.Password,
		Role:         u.Role,
		IsActive:     u.IsActive,
	}, nil
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}
	logger.Init(cfg.Log.Level, cfg.Log.Format)
	logger.Info("platform API starting")

	ctx := context.Background()

	db, err := database.Init(ctx, cfg.Database)
	if err != nil {
		logger.Fatal(err, "database init")
	}
	defer db.Close()

	redisCache, err := cache.NewRedis(cfg.Redis.URL)
	if err != nil {
		logger.Fatal(err, "redis init")
	}
	defer redisCache.Close()

	var broker events.Broker
	if cfg.Events.Broker == "rabbitmq" {
		broker, err = events.NewRabbitMQBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
		if err != nil {
			logger.Fatal(err, "rabbitmq init")
		}
	} else {
		broker = events.NewInMemoryBroker()
		logger.Info("using in-memory event broker")
	}
	defer broker.Close()

	jwtManager := jwt.NewManager(cfg.JWT.Secret, cfg.JWT.Expiry, cfg.JWT.RefreshExpiry)

	// ── Repositories ─────────────────────────────────────────────────────────
	userRepo := users.NewPostgresRepository(db)
	roleRepo := roles.NewPostgresRepository(db)
	itemRepo := items.NewPostgresRepository(db)
	orderRepo := orders.NewPostgresRepository(db)
	paymentRepo := payments.NewPostgresRepository(db)
	notifRepo := notifications.NewPostgresRepository(db)
	uploadRepo := uploads.NewPostgresRepository(db)
	searchRepo := search.NewPostgresRepository(db)
	tokenRepo := auth.NewRedisTokenRepo(redisCache)

	// ── Services ─────────────────────────────────────────────────────────────
	authLookup := &authUserAdapter{repo: userRepo}
	authSvc := auth.NewService(authLookup, tokenRepo, jwtManager)
	userSvc := users.NewService(userRepo)
	roleSvc := roles.NewService(roleRepo)
	itemSvc := items.NewService(itemRepo, broker)
	orderSvc := orders.NewService(orderRepo, itemRepo, broker)

	paymentSvc := payments.NewService(paymentRepo, orderRepo, []providers.Provider{
		providers.NewMPesaProvider("", "", "", "", "", ""),
		providers.NewStripeProvider("", ""),
		providers.NewFlutterwaveProvider("", ""),
		providers.NewPaystackProvider("", ""),
	}, broker)

	notifSvc := notifications.NewService(notifRepo, []notifchannels.Channel{
		notifchannels.NewEmailChannel("localhost", "1025", "", "", "noreply@platform.local"),
		notifchannels.NewSMSChannel("", ""),
		notifchannels.NewPushChannel(""),
	})
	uploadSvc := uploads.NewService(uploadRepo, uploadstorage.NewLocalStorage("./uploads", cfg.App.Name), "local")
	searchSvc := search.NewService(searchRepo)
	adminSvc := admin.NewService(db)
	wsHub := ws.NewHub()

	// Seed default roles
	if err := roleSvc.SeedDefaults(ctx); err != nil {
		logger.Error(err, "role seed (non-fatal)")
	}

	// ── Event subscriptions ──────────────────────────────────────────────────
	// In main.go, update the event handler subscriptions:
	broker.Subscribe(events.TopicOrderConfirmed, evhandlers.OrderConfirmedWithWS(notifSvc, wsHub))
	broker.Subscribe(events.TopicPaymentCompleted, evhandlers.PaymentCompletedWithWS(notifSvc, wsHub))
	broker.Subscribe(events.TopicOrderPlaced, evhandlers.AuditLog())
	broker.Subscribe(events.TopicPaymentFailed, evhandlers.AuditLog())
	broker.Subscribe(events.TopicUserRegistered, evhandlers.AuditLog())

	// ── Handlers ─────────────────────────────────────────────────────────────
	authH := auth.NewHandler(authSvc)
	userH := users.NewHandler(userSvc)
	roleH := roles.NewHandler(roleSvc)
	itemH := items.NewHandler(itemSvc)
	orderH := orders.NewHandler(orderSvc)
	paymentH := payments.NewHandler(paymentSvc)
	notifH := notifications.NewHandler(notifSvc)
	uploadH := uploads.NewHandler(uploadSvc)
	searchH := search.NewHandler(searchSvc)
	adminH := admin.NewHandler(adminSvc)
	wsH := ws.NewHandler(wsHub, jwtManager)

	// Wire userSvc into authH so Register can create users.
	authH.SetUserCreator(userSvc)

	// ── Router ───────────────────────────────────────────────────────────────
	r := chi.NewRouter()

	// Global middleware — ORDER MATTERS.
	// RequestID and Recovery wrap everything including WS.
	r.Use(middleware.RequestID)
	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)
	r.Use(middleware.CORS(middleware.CORSOptions{
		AllowedOrigins: cfg.CORS.Origins,
		MaxAge:         cfg.CORS.MaxAge,
	}))
	// NOTE: Compress is applied ONLY to non-WS routes below.
	// Registering it globally wraps ResponseWriter in a compressor that
	// does not implement http.Hijacker, causing the WS upgrade to panic.

	r.Get("/api/health", healthHandler(db))

	// WebSocket route — must be outside the compress group.
	// The upgrade requires raw http.Hijacker access on the ResponseWriter.
	ws.RegisterRoutes(r, wsH)

	// All other API routes get compression.
	r.Group(func(r chi.Router) {
		r.Use(chimw.Compress(5))

		r.Route("/api/v1", func(r chi.Router) {
			auth.RegisterRoutes(r, authH, jwtManager)
			search.RegisterRoutes(r, searchH)
			items.RegisterRoutes(r, itemH, jwtManager)
			payments.RegisterRoutes(r, paymentH, jwtManager)

			r.Group(func(r chi.Router) {
				r.Use(middleware.Authenticate(jwtManager))
				users.RegisterRoutes(r, userH)
				roles.RegisterRoutes(r, roleH)
				orders.RegisterRoutes(r, orderH)
				notifications.RegisterRoutes(r, notifH, jwtManager)
				uploads.RegisterRoutes(r, uploadH, jwtManager)
				admin.RegisterRoutes(r, adminH, jwtManager)
			})
		})
	})

	_ = userH
	_ = roleH
	_ = orderH
	_ = notifH
	_ = uploadH
	_ = searchH
	_ = adminH
	_ = paymentH

	// ── HTTP server ──────────────────────────────────────────────────────────
	srv := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("listening on " + cfg.Addr())
		errCh <- srv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err, "server error")
		}
	case sig := <-quit:
		logger.Info("shutdown: " + sig.String())
	}

	shutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutCtx); err != nil {
		logger.Error(err, "graceful shutdown")
		os.Exit(1)
	}
	logger.Info("server stopped cleanly")
}

func healthHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := db.Ping(r.Context())
		status := "ok"
		if err != nil {
			status = "db_error"
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status": status,
			"time":   time.Now().UTC(),
		})
	}
}
