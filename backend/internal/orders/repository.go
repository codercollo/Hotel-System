package orders

import "context"

// Status Constants

// Repository defines persistence for orders.
type Repository interface {
	FindByID(ctx context.Context, id string) (*Order, error)
	ListForUser(ctx context.Context, userID string, limit, offset int) ([]*Order, int, error)
	ListAll(ctx context.Context, limit, offset int) ([]*Order, int, error)
	Create(ctx context.Context, order *Order, items []OrderItem) (*Order, error)
	UpdateStatus(ctx context.Context, id, status string) (*Order, error)
}
