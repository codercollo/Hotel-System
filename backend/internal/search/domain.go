// Package search provides a decoupled search abstraction.
// The SearchRepo interface allows swapping PostgreSQL tsvector for Elasticsearch
// in Phase 8 without touching any other module.
package search

import "time"

// SearchQuery is the validated search request.
type SearchQuery struct {
	Q      string
	Type   string // "items" | "" (all)
	Limit  int
	Offset int
}

// SearchResult is a single hit returned from a search.
type SearchResult struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       *int64         `json:"price,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	Score       float64        `json:"score"`
	CreatedAt   time.Time      `json:"created_at"`
}
