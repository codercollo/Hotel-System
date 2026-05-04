package users

import (
	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes mounts user endpoints. All routes require authentication.
func RegisterRoutes(r chi.Router, h *Handler) {
	r.Get("/users", h.List)
	r.Get("/users/{id}", h.Get)
	r.Patch("/users/{id}", h.Update)

	// Admin-only delete
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAdmin())
		r.Delete("/users/{id}", h.Delete)
	})
}
