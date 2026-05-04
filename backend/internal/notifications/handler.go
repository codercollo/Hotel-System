package notifications

import (
	"net/http"

	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/codercollo/hotel-system/backend/pkg/pagination"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/go-chi/chi/v5"
)

// Handler holds HTTP handlers for the notifications module.
type Handler struct{ svc *Service }

// NewHandler creates a notifications Handler.
func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }

// List godoc — GET /api/v1/notifications
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	p := pagination.Parse(r)

	notes, total, err := h.svc.List(r.Context(), userID, p)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONPaginated(w, notes, pagination.NewMeta(p, total))
}

// MarkRead godoc — PATCH /api/v1/notifications/{id}/read
func (h *Handler) MarkRead(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := middleware.GetUserID(r.Context())

	if err := h.svc.MarkRead(r.Context(), id, userID); err != nil {
		response.Error(w, err)
		return
	}
	response.NoContent(w)
}

// MarkAllRead godoc — PATCH /api/v1/notifications/read-all
func (h *Handler) MarkAllRead(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	if err := h.svc.MarkAllRead(r.Context(), userID); err != nil {
		response.Error(w, err)
		return
	}
	response.NoContent(w)
}

// RegisterRoutes mounts notification endpoints. All require authentication.
func RegisterRoutes(r chi.Router, h *Handler, jwtManager *jwt.Manager) {
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate(jwtManager))
		r.Get("/notifications", h.List)
		r.Patch("/notifications/read-all", h.MarkAllRead)
		r.Patch("/notifications/{id}/read", h.MarkRead)
	})
}
