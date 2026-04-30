package middleware

import (
	"net/http"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/response"
)

// RequireRole allows only users whose role matches one of the provided roles.
// Must be placed after Authenticate in the middleware chain.
func RequireRole(roles ...string) func(http.Handler) http.Handler {
	allowed := make(map[string]struct{}, len(roles))
	for _, r := range roles {
		allowed[r] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := GetClaims(r.Context())
			if claims == nil {
				response.Error(w, apierror.ErrUnauthorized)
				return
			}

			if _, ok := allowed[claims.Role]; !ok {
				response.Error(w, apierror.ErrForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAdmin is a convenience middleware that only allows users with role "admin".
func RequireAdmin() func(http.Handler) http.Handler {
	return RequireRole("admin")
}
