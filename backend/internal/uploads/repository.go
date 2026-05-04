package uploads

import "context"

// Repository defines persistence for upload metadata.
type Repository interface {
	Create(ctx context.Context, u *Upload) (*Upload, error)
	FindByID(ctx context.Context, id string) (*Upload, error)
	ListForUser(ctx context.Context, userID string) ([]*Upload, error)
	Delete(ctx context.Context, id string) error
}
