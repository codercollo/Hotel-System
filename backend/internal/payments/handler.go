package payments

import (
	"net/http"

	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/codercollo/hotel-system/backend/pkg/validator"
	"github.com/go-chi/chi/v5"
)

// Handler holds HTTP handlers for the payments module.
type Handler struct{ svc *Service }

// NewHandler creates a payments Handler.
func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }

// Initiate godoc — POST /api/v1/payments/initiate
func (h *Handler) Initiate(w http.ResponseWriter, r *http.Request) {
	var req InitiateRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		response.Error(w, err)
		return
	}

	userID := middleware.GetUserID(r.Context())
	res, err := h.svc.Initiate(r.Context(), req, userID)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONCreated(w, res)
}

// Webhook godoc — POST /api/v1/payments/webhook/{provider}
func (h *Handler) Webhook(w http.ResponseWriter, r *http.Request) {
	providerName := chi.URLParam(r, "provider")
	if err := h.svc.HandleWebhook(r.Context(), r, providerName); err != nil {
		response.Error(w, err)
		return
	}
	response.NoContent(w)
}

// Get godoc — GET /api/v1/payments/{id}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, err := h.svc.Get(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, p)
}

// ListForOrder godoc — GET /api/v1/orders/{orderId}/payments
func (h *Handler) ListForOrder(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "orderId")
	payments, err := h.svc.ListForOrder(r.Context(), orderID)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, payments)
}
