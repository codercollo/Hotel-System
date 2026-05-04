package search

import "context"

// Repository is the swappable search backend interface.
// Implementations: repository_postgres.go (Phase 8 default), elasticsearch (Phase 8 upgrade).
type Repository interface {
	Search(ctx context.Context, query SearchQuery) ([]*SearchResult, int, error)
}
