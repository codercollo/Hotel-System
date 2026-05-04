package users

import (
	"net/http"

	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/pagination"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/codercollo/hotel-system/backend/pkg/validator"
	"github.com/go-chi/chi/v5"
)

// Handler holds HTTP handlers for the users module.
type Handler struct {
	svc *Service
}

// NewHandler constructs a users Handler.
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// List godoc — GET /api/v1/users  (admin only)
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	p := pagination.Parse(r)
	users, total, err := h.svc.List(r.Context(), p)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONPaginated(w, users, pagination.NewMeta(p, total))
}

// Get godoc — GET /api/v1/users/{id}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Regular users may only fetch their own profile.
	claims := middleware.GetClaims(r.Context())
	if claims.Role != "admin" && claims.UserID != id {
		response.Error(w, apierror.ErrForbidden)
		return
	}

	user, err := h.svc.Get(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, user)
}

// Update godoc — PATCH /api/v1/users/{id}
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	claims := middleware.GetClaims(r.Context())
	if claims.Role != "admin" && claims.UserID != id {
		response.Error(w, apierror.ErrForbidden)
		return
	}

	var body struct {
		Name     *string        `json:"name"`
		Email    *string        `json:"email"`
		IsActive *bool          `json:"is_active"`
		Metadata map[string]any `json:"metadata"`
	}
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	user, err := h.svc.Update(r.Context(), id, UpdateUserInput{
		Name:     body.Name,
		Email:    body.Email,
		IsActive: body.IsActive,
		Metadata: body.Metadata,
	})
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, user)
}

// Delete godoc — DELETE /api/v1/users/{id}  (admin only)
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		response.Error(w, err)
		return
	}
	response.NoContent(w)
}
