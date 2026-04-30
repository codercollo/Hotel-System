package middleware

import (
	"context"
	"net/http"
	"strings"

	pkgjwt "github.com/codercollo/hotel-system/backend/pkg/jwt" // Changed internal -> pkg
	"github.com/codercollo/hotel-system/backend/pkg/response"
)

type authContextKey string

const (
	claimsKey authContextKey = "jwt_claims"
)

// Authenticate validates the Bearer token and attaches claims to the context.
// Routes behind this middleware require a valid JWT access token.
func Authenticate(jwtManager *pkgjwt.Manager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractBearerToken(r)
			if token == "" {
				response.Error(w, pkgjwt.ErrMissingToken)
				return
			}

			claims, err := jwtManager.Parse(token)
			if err != nil {
				response.Error(w, err)
				return
			}

			if claims.TokenType != pkgjwt.TokenTypeAccess {
				response.Error(w, pkgjwt.ErrWrongTokenType)
				return
			}

			ctx := context.WithValue(r.Context(), claimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetClaims retrieves the JWT claims from the request context.
// Returns nil if the context has no claims (unauthenticated route).
func GetClaims(ctx context.Context) *pkgjwt.Claims {
	c, _ := ctx.Value(claimsKey).(*pkgjwt.Claims)
	return c
}

// GetUserID is a convenience wrapper over GetClaims.
func GetUserID(ctx context.Context) string {
	if c := GetClaims(ctx); c != nil {
		return c.UserID
	}
	return ""
}

func extractBearerToken(r *http.Request) string {
	hdr := r.Header.Get("Authorization")
	if !strings.HasPrefix(hdr, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(hdr, "Bearer ")
}
