package roles

import (
	"context"
	"net/http"

	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/go-chi/chi/v5"
)

// ─── Service ─────────────────────────────────────────────────────────────────

// Service implements role management logic.
type Service struct{ repo Repository }

// NewService creates a roles Service.
func NewService(repo Repository) *Service { return &Service{repo: repo} }

func (s *Service) List(ctx context.Context) ([]*Role, error) { return s.repo.ListRoles(ctx) }

func (s *Service) HasPermission(ctx context.Context, roleName, perm string) (bool, error) {
	return s.repo.HasPermission(ctx, roleName, perm)
}

func (s *Service) SeedDefaults(ctx context.Context) error { return s.repo.SeedDefaults(ctx) }

// ─── Handler ─────────────────────────────────────────────────────────────────

// Handler holds HTTP handlers for the roles module.
type Handler struct{ svc *Service }

// NewHandler creates a roles Handler.
func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }

// List godoc — GET /api/v1/roles  (admin only)
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	roles, err := h.svc.List(r.Context())
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, roles)
}

// ─── Routes ──────────────────────────────────────────────────────────────────

// RegisterRoutes mounts roles endpoints onto the router.
func RegisterRoutes(r chi.Router, h *Handler) {
	r.Get("/roles", h.List)
}
