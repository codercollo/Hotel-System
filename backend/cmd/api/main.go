// Package main is the entrypoint for the platform API server.
// It wires together configuration, infrastructure, and HTTP routing,
// then starts a gracefully-shutdownable HTTP server.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/codercollo/hotel-system/backend/config"
	"github.com/codercollo/hotel-system/backend/database"
	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/cache"
	"github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
	"github.com/codercollo/hotel-system/backend/pkg/response"
)

func main() {
	// ── Config ───────────────────────────────────────────────────────────────
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config error: %v\n", err)
		os.Exit(1)
	}

	// ── Logger ───────────────────────────────────────────────────────────────
	logger.Init(cfg.Log.Level, cfg.Log.Format)
	logger.Info("starting platform API")

	// ── Database ─────────────────────────────────────────────────────────────
	ctx := context.Background()
	pool, err := database.Init(ctx, cfg.Database)
	if err != nil {
		logger.Fatal(err, "database init failed")
	}
	defer pool.Close()
	logger.Info("database connected")

	// ── Cache (Redis) ────────────────────────────────────────────────────────
	redisCache, err := cache.NewRedis(cfg.Redis.URL)
	if err != nil {
		logger.Fatal(err, "redis init failed")
	}
	defer redisCache.Close()
	logger.Info("redis connected")

	// ── Event Broker ─────────────────────────────────────────────────────────
	var broker events.Broker
	if cfg.Events.Broker == "rabbitmq" {
		broker, err = events.NewRabbitMQBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
		if err != nil {
			logger.Fatal(err, "rabbitmq init failed")
		}
		logger.Info("rabbitmq connected")
	} else {
		broker = events.NewInMemoryBroker()
		logger.Info("using in-memory event broker")
	}
	defer broker.Close()

	// ── JWT Manager ──────────────────────────────────────────────────────────
	jwtManager := jwt.NewManager(cfg.JWT.Secret, cfg.JWT.Expiry, cfg.JWT.RefreshExpiry)

	// ── Router ───────────────────────────────────────────────────────────────
	r := chi.NewRouter()

	// Global middleware chain
	r.Use(middleware.RequestID)
	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)
	r.Use(middleware.CORS(middleware.CORSOptions{
		AllowedOrigins: cfg.CORS.Origins,
		MaxAge:         cfg.CORS.MaxAge,
	}))
	r.Use(chimiddleware.Compress(5))

	// ── Health endpoint ──────────────────────────────────────────────────────
	r.Get("/api/health", response.HealthHandler(pool))

	// ── API v1 routes (modules registered here in later phases) ─────────────
	r.Route("/api/v1", func(r chi.Router) {
		// Auth middleware available for protected sub-routes:
		// r.Group(func(r chi.Router) {
		//     r.Use(middleware.Authenticate(jwtManager))
		//     auth.RegisterRoutes(r, authSvc)
		// })

		// Placeholder to prove the router is wired correctly.
		r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"success":true,"data":"pong"}`)
		})
	})

	// Suppress "declared and not used" errors while modules are not yet wired.
	_ = jwtManager
	_ = broker
	_ = redisCache

	// ── HTTP Server ──────────────────────────────────────────────────────────
	srv := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine so we can listen for shutdown signals.
	serverErr := make(chan error, 1)
	go func() {
		logger.Info("HTTP server listening on " + cfg.Addr())
		serverErr <- srv.ListenAndServe()
	}()

	// ── Graceful Shutdown ────────────────────────────────────────────────────
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err, "server error")
		}
	case sig := <-quit:
		logger.Info("shutdown signal received: " + sig.String())
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error(err, "graceful shutdown failed")
		os.Exit(1)
	}

	logger.Info("server stopped cleanly")
}
