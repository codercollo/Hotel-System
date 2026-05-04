package auth

import (
	"context"
	"net/http"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/password"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/codercollo/hotel-system/backend/pkg/validator"
)

// UserCreator is satisfied by users.Service — avoids import cycle.
type UserCreator interface {
	Create(ctx context.Context, name, email, passwordHash, role string) (AuthUser, error)
}

// Handler holds HTTP handlers for the auth module.
type Handler struct {
	svc         *Service
	userCreator UserCreator
}

// NewHandler constructs an auth Handler.
// Signature: NewHandler(authSvc) — matches main.go.
// Call SetUserCreator(userSvc) after wiring users.Service.
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// SetUserCreator injects the user creation dependency after initial construction.
// Called in main.go after both authH and userH are initialised.
func (h *Handler) SetUserCreator(uc UserCreator) {
	h.userCreator = uc
}

// Register — POST /api/v1/auth/register
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var body RegisterInput
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	hash, err := password.Hash(body.Password)
	if err != nil {
		response.Error(w, apierror.ErrInternal)
		return
	}

	if h.userCreator == nil {
		response.Error(w, apierror.ErrInternal)
		return
	}

	user, err := h.userCreator.Create(r.Context(), body.Name, body.Email, hash, "guest")
	if err != nil {
		response.Error(w, err)
		return
	}

	pair, err := h.svc.IssueForUser(r.Context(), user, r.UserAgent(), r.RemoteAddr)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.JSONCreated(w, pair)
}

// Login — POST /api/v1/auth/login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var body LoginInput
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	pair, err := h.svc.Login(r.Context(), body, r.UserAgent(), r.RemoteAddr)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.JSON(w, pair)
}

// Refresh — POST /api/v1/auth/refresh
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var body RefreshInput
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	pair, err := h.svc.Refresh(r.Context(), body.RefreshToken, r.UserAgent(), r.RemoteAddr)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.JSON(w, pair)
}

// Logout — POST /api/v1/auth/logout
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var body RefreshInput
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	if err := h.svc.Logout(r.Context(), body.RefreshToken); err != nil {
		response.Error(w, err)
		return
	}

	response.NoContent(w)
}

// Me — GET /api/v1/auth/me  (requires Authenticate middleware upstream)
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	claims, err := h.svc.VerifyToken(extractBearer(r))
	if err != nil {
		response.Error(w, apierror.ErrUnauthorized)
		return
	}

	response.JSON(w, map[string]any{
		"user_id": claims.UserID,
		"email":   claims.Email,
		"role":    claims.Role,
	})
}

func extractBearer(r *http.Request) string {
	const prefix = "Bearer "
	h := r.Header.Get("Authorization")
	if len(h) > len(prefix) {
		return h[len(prefix):]
	}
	return ""
}
