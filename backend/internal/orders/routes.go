package orders

import (
	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes mounts orders endpoints. All routes require authentication.
func RegisterRoutes(r chi.Router, h *Handler) {
	r.Get("/orders", h.List)
	r.Post("/orders", h.Place)
	r.Get("/orders/{id}", h.Get)
	r.Patch("/orders/{id}/cancel", h.Cancel)

	// Admin status management
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAdmin())
		r.Patch("/orders/{id}/status", h.UpdateStatus)
	})
}
