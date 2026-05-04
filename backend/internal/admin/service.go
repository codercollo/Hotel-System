package admin

import (
	"context"
	"net/http"

	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/codercollo/hotel-system/backend/pkg/validator"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Stats is the admin dashboard aggregate.
type Stats struct {
	TotalUsers   int   `json:"total_users"`
	TotalItems   int   `json:"total_items"`
	TotalOrders  int   `json:"total_orders"`
	RevenueTotal int64 `json:"revenue_total"` // sum of completed order totals
}

// ─── Service ─────────────────────────────────────────────────────────────────

// Service provides platform-wide aggregations for the admin module.
type Service struct{ db *pgxpool.Pool }

// NewService creates an admin Service backed by direct DB queries.
func NewService(db *pgxpool.Pool) *Service { return &Service{db: db} }

// GetStats returns platform-wide statistics.
func (s *Service) GetStats(ctx context.Context) (*Stats, error) {
	stats := &Stats{}

	_ = s.db.QueryRow(ctx, `SELECT COUNT(*) FROM users WHERE deleted_at IS NULL`).Scan(&stats.TotalUsers)
	_ = s.db.QueryRow(ctx, `SELECT COUNT(*) FROM items WHERE deleted_at IS NULL`).Scan(&stats.TotalItems)
	_ = s.db.QueryRow(ctx, `SELECT COUNT(*) FROM orders WHERE deleted_at IS NULL`).Scan(&stats.TotalOrders)
	_ = s.db.QueryRow(ctx,
		`SELECT COALESCE(SUM(total),0) FROM orders WHERE status='completed' AND deleted_at IS NULL`,
	).Scan(&stats.RevenueTotal)

	return stats, nil
}

// ─── Handler ─────────────────────────────────────────────────────────────────

// Handler holds HTTP handlers for the admin module.
type Handler struct{ svc *Service }

// NewHandler creates an admin Handler.
func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }

// Stats godoc — GET /api/v1/admin/stats
func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.svc.GetStats(r.Context())
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, stats)
}

type ContactInput struct {
	Name    string `json:"name"    validate:"required"`
	Email   string `json:"email"   validate:"required,email"`
	Phone   string `json:"phone"`
	Subject string `json:"subject" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func (h *Handler) Contact(w http.ResponseWriter, r *http.Request) {
	var input ContactInput
	if err := validator.DecodeAndValidate(r, &input); err != nil {
		response.Error(w, err)
		return
	}
	// Phase 10: deliver via email channel
	// For now log it and return success
	logger.Log().Info().
		Str("name", input.Name).
		Str("email", input.Email).
		Str("subject", input.Subject).
		Msg("contact form submission")
	response.JSON(w, map[string]any{"received": true})
}

// ─── Routes ──────────────────────────────────────────────────────────────────

// RegisterRoutes mounts admin endpoints. All require admin role.
func RegisterRoutes(r chi.Router, h *Handler, jwtManager *jwt.Manager) {
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate(jwtManager))
		r.Use(middleware.RequireAdmin())
		r.Get("/admin/stats", h.Stats)
	})
}
