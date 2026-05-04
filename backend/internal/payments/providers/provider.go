// Package providers defines the payment provider interface and stubs.
// Each provider (M-Pesa, Stripe, Flutterwave, Paystack) implements this interface.
// Concrete implementations are completed in Phase 5.
package providers

import "context"

// InitiateParams carries the data each provider needs to start a transaction.
type InitiateParams struct {
	PaymentID string
	OrderID   string
	Amount    int64
	Currency  string
	Phone     string // mobile money
	Email     string // card/online
	Reference string // your internal ref
}

// InitiateResult is what the provider returns after successfully initiating payment.
type InitiateResult struct {
	ProviderRef  string // provider's transaction ID
	CheckoutURL  string // redirect URL (Stripe, Flutterwave)
	CheckoutCode string // STK push code (M-Pesa)
	Status       string // initial status
}

// WebhookPayload is the raw body + signature from a provider webhook.
type WebhookPayload struct {
	Body      []byte
	Signature string
	Provider  string
}

// Provider is the interface every payment gateway must satisfy.
type Provider interface {
	// Name returns the canonical provider identifier (e.g. "mpesa").
	Name() string

	// Initiate starts a payment and returns a result the caller stores.
	Initiate(ctx context.Context, params InitiateParams) (*InitiateResult, error)

	// VerifyWebhook validates the webhook signature and returns the payment status.
	VerifyWebhook(payload WebhookPayload) (providerRef, status string, err error)
}
