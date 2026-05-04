package items

import (
	"net/http"

	"github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/pagination"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/codercollo/hotel-system/backend/pkg/validator"
	"github.com/go-chi/chi/v5"
)

// Handler holds HTTP handlers for the items module.
type Handler struct {
	svc *Service
}

// NewHandler creates an items Handler.
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// List godoc — GET /api/v1/items
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	p := pagination.Parse(r)
	status := r.URL.Query().Get("status")

	items, total, err := h.svc.List(r.Context(), p, status)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONPaginated(w, items, pagination.NewMeta(p, total))
}

// Get godoc — GET /api/v1/items/{id}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item, err := h.svc.Get(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, item)
}

// Create godoc — POST /api/v1/items  (admin)
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name        string         `json:"name"        validate:"required,min=2"`
		Description string         `json:"description"`
		Price       int64          `json:"price"       validate:"gte=0"`
		Currency    string         `json:"currency"`
		Stock       int            `json:"stock"       validate:"gte=0"`
		Images      []string       `json:"images"`
		Metadata    map[string]any `json:"metadata"`
	}
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	userID := middleware.GetUserID(r.Context())
	item, err := h.svc.Create(r.Context(), CreateItemInput{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Currency:    body.Currency,
		Stock:       body.Stock,
		Images:      body.Images,
		Metadata:    body.Metadata,
		CreatedBy:   &userID,
	})
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONCreated(w, item)
}

// Update godoc — PATCH /api/v1/items/{id}  (admin)
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var body struct {
		Name        *string        `json:"name"`
		Description *string        `json:"description"`
		Price       *int64         `json:"price"`
		Status      *string        `json:"status"`
		Stock       *int           `json:"stock"`
		Images      []string       `json:"images"`
		Metadata    map[string]any `json:"metadata"`
	}
	if err := validator.DecodeAndValidate(r, &body); err != nil {
		response.Error(w, err)
		return
	}

	item, err := h.svc.Update(r.Context(), id, UpdateItemInput{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Status:      body.Status,
		Stock:       body.Stock,
		Images:      body.Images,
		Metadata:    body.Metadata,
	})
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSON(w, item)
}

// Delete godoc — DELETE /api/v1/items/{id}  (admin)
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		response.Error(w, err)
		return
	}
	response.NoContent(w)
}

// Search godoc — GET /api/v1/items/search?q=...
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	p := pagination.Parse(r)

	items, total, err := h.svc.Search(r.Context(), q, p)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONPaginated(w, items, pagination.NewMeta(p, total))
}
