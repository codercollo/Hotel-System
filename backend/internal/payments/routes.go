package payments

import (
	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes mounts payments endpoints.
func RegisterRoutes(r chi.Router, h *Handler, jwtManager *jwt.Manager) {
	// Webhooks are unauthenticated (signature-verified inside handler)
	r.Post("/payments/webhook/{provider}", h.Webhook)

	// Protected payment routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate(jwtManager))
		r.Post("/payments/initiate", h.Initiate)
		r.Get("/payments/{id}", h.Get)
		r.Get("/orders/{orderId}/payments", h.ListForOrder)
	})
}
