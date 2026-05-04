package orders

import (
	"net/http"

	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/pagination"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/codercollo/hotel-system/backend/pkg/validator"
	"github.com/go-chi/chi/v5"
)

// Handler holds HTTP handlers for the orders module.
type Handler struct{ svc *Service }

// NewHandler creates an orders Handler.
func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }

// List godoc — GET /api/v1/orders
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	p := pagination.Parse(r)
	claims := middleware.GetClaims(r.Context())

	var (
		orders []*Order
		total  int
		err    error
	)

	if claims.Role == "admin" {
		orders, total, err = h.svc.ListAll(r.Context(), p)
	} else {
		orders, total, err = h.svc.ListForUser(r.Context(), claims.UserID, p)
	}

	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONPaginated(w, orders, pagination.NewMeta(p, total))
}

// Get godoc — GET /api/v1/orders/{id}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	claims := middleware.GetClaims(r.Context())

	order, err := h.svc.Get(r.Context(), id, claims.UserID, claims.Role)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, order)
}

// Place godoc — POST /api/v1/orders
func (h *Handler) Place(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Items    []OrderItemInput `json:"items"    validate:"required,min=1,dive"`
		Notes    string           `json:"notes"`
		Metadata map[string]any   `json:"metadata"`
	}
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	userID := middleware.GetUserID(r.Context())
	order, err := h.svc.Place(r.Context(), CreateOrderInput{
		UserID:   userID,
		Items:    body.Items,
		Notes:    body.Notes,
		Metadata: body.Metadata,
	})
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONCreated(w, order)
}

// UpdateStatus godoc — PATCH /api/v1/orders/{id}/status  (admin)
func (h *Handler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	claims := middleware.GetClaims(r.Context())

	var body struct {
		Status string `json:"status" validate:"required"`
	}
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	order, err := h.svc.UpdateStatus(r.Context(), id, body.Status, claims.Role)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, order)
}

// Cancel godoc — PATCH /api/v1/orders/{id}/cancel
func (h *Handler) Cancel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	claims := middleware.GetClaims(r.Context())

	order, err := h.svc.Cancel(r.Context(), id, claims.UserID, claims.Role)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, order)
}
