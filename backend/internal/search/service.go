package search

import (
	"context"
	"net/http"

	"github.com/codercollo/hotel-system/backend/pkg/pagination"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/go-chi/chi/v5"
)

// ─── Service ─────────────────────────────────────────────────────────────────

// Service wraps the search repository with business logic.
type Service struct{ repo Repository }

// NewService creates a search Service.
func NewService(repo Repository) *Service { return &Service{repo: repo} }

// Search performs a full-text search.
func (s *Service) Search(ctx context.Context, q string, p pagination.Params) ([]*SearchResult, int, error) {
	return s.repo.Search(ctx, SearchQuery{Q: q, Limit: p.Limit, Offset: p.Offset})
}

// ─── Handler ─────────────────────────────────────────────────────────────────

// Handler holds HTTP handlers for the search module.
type Handler struct{ svc *Service }

// NewHandler creates a search Handler.
func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }

// Search godoc — GET /api/v1/search?q=...
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	p := pagination.Parse(r)

	results, total, err := h.svc.Search(r.Context(), q, p)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONPaginated(w, results, pagination.NewMeta(p, total))
}

// ─── Routes ──────────────────────────────────────────────────────────────────

// RegisterRoutes mounts the search endpoint (public).
func RegisterRoutes(r chi.Router, h *Handler) {
	r.Get("/search", h.Search)
}
