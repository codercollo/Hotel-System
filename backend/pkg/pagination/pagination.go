// Package pagination provides offset-based and cursor-based pagination helpers.
// All list endpoints use the Meta struct as the standard envelope metadata.
package pagination

import (
	"math"
	"net/http"
	"strconv"
)

const (
	defaultPage  = 1
	defaultLimit = 20
	maxLimit     = 100
)

// Params holds parsed query parameters for list requests.
type Params struct {
	Page   int
	Limit  int
	Offset int
}

// Meta is the pagination block included in paginated JSON responses.
type Meta struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

// Parse extracts ?page= and ?limit= from the request query string.
// Invalid or missing values fall back to defaults.
func Parse(r *http.Request) Params {
	q := r.URL.Query()

	page, err := strconv.Atoi(q.Get("page"))
	if err != nil || page < 1 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil || limit < 1 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	return Params{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}

// NewMeta constructs a Meta block from params and the total record count.
func NewMeta(p Params, total int) Meta {
	totalPages := int(math.Ceil(float64(total) / float64(p.Limit)))
	if totalPages < 1 {
		totalPages = 1
	}
	return Meta{
		Page:       p.Page,
		Limit:      p.Limit,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    p.Page < totalPages,
		HasPrev:    p.Page > 1,
	}
}
