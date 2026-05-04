package notifications

import "context"

// Repository defines persistence for in-app notifications.
type Repository interface {
	Create(ctx context.Context, n *Notification) (*Notification, error)
	ListForUser(ctx context.Context, userID string, limit, offset int) ([]*Notification, int, error)
	MarkRead(ctx context.Context, id, userID string) error
	MarkAllRead(ctx context.Context, userID string) error
}
