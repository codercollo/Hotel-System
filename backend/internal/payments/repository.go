package payments

import "context"

// Repository defines persistence for payments.
type Repository interface {
	FindByID(ctx context.Context, id string) (*Payment, error)
	FindByOrderID(ctx context.Context, orderID string) ([]*Payment, error)
	FindByProviderRef(ctx context.Context, providerRef string) (*Payment, error)
	Create(ctx context.Context, p *Payment) (*Payment, error)
	UpdateStatus(ctx context.Context, id, status string, providerRef *string) (*Payment, error)
	StoreWebhookPayload(ctx context.Context, id string, payload map[string]any) error
}
