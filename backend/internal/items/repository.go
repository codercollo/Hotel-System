package items

import "context"

// Repository defines persistence for items.
type Repository interface {
	FindByID(ctx context.Context, id string) (*Item, error)
	List(ctx context.Context, filters ListFilters) ([]*Item, int, error)
	Create(ctx context.Context, input CreateItemInput) (*Item, error)
	Update(ctx context.Context, id string, input UpdateItemInput) (*Item, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string, limit, offset int) ([]*Item, int, error)
}

// ListFilters holds optional filter parameters for item listing.
type ListFilters struct {
	Status string
	Limit  int
	Offset int
}
