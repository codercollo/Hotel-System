package auth

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	pkgjwt "github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/codercollo/hotel-system/backend/pkg/response"
)

// RegisterRoutes mounts auth endpoints onto r.
// Signature: RegisterRoutes(r, handler, jwtManager) — matches main.go.
func RegisterRoutes(r chi.Router, h *Handler, jwtMgr *pkgjwt.Manager) {
	r.Route("/auth", func(r chi.Router) {
		// Public routes
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Post("/refresh", h.Refresh)

		// Protected routes — inline JWT guard so we don't need middleware pkg here
		r.Group(func(r chi.Router) {
			r.Use(jwtAuthMiddleware(jwtMgr))
			r.Post("/logout", h.Logout)
			r.Get("/me", h.Me)
		})
	})
}

// jwtAuthMiddleware is a minimal inline middleware for auth routes only.
// The main middleware.Authenticate is used for all other module groups.
func jwtAuthMiddleware(jwtMgr *pkgjwt.Manager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractBearerFromHeader(r.Header.Get("Authorization"))
			if token == "" {
				response.Error(w, apierror.ErrUnauthorized)
				return
			}
			if _, err := jwtMgr.Parse(token); err != nil {
				response.Error(w, apierror.ErrUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func extractBearerFromHeader(header string) string {
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}
	return ""
}
