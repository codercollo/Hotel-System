package items

import (
	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes mounts items endpoints onto the router.
func RegisterRoutes(r chi.Router, h *Handler, jwtManager *jwt.Manager) {
	// Public routes
	r.Get("/items", h.List)
	r.Get("/items/search", h.Search)
	r.Get("/items/{id}", h.Get)

	// Admin-only routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate(jwtManager))
		r.Use(middleware.RequireAdmin())
		r.Post("/items", h.Create)
		r.Patch("/items/{id}", h.Update)
		r.Delete("/items/{id}", h.Delete)
	})
}
