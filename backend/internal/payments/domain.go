package payments

import "time"

// Payment represents a financial transaction for an order.
type Payment struct {
	ID             string         `json:"id"`
	OrderID        string         `json:"order_id"`
	Provider       string         `json:"provider"` // mpesa|stripe|flutterwave|paystack
	ProviderRef    *string        `json:"provider_ref,omitempty"`
	Amount         int64          `json:"amount"` // smallest currency unit
	Currency       string         `json:"currency"`
	Status         string         `json:"status"` // pending|processing|completed|failed|refunded
	WebhookPayload map[string]any `json:"webhook_payload,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

// InitiateRequest is the validated DTO for starting a payment.
type InitiateRequest struct {
	OrderID  string `json:"order_id"  validate:"required"`
	Provider string `json:"provider"  validate:"required,oneof=mpesa stripe flutterwave paystack"`
	Phone    string `json:"phone"` // used by mobile-money providers
}

// InitiateResponse is what the API returns after starting a payment.
type InitiateResponse struct {
	PaymentID    string `json:"payment_id"`
	CheckoutURL  string `json:"checkout_url,omitempty"`  // Stripe / Flutterwave redirect
	CheckoutCode string `json:"checkout_code,omitempty"` // M-Pesa STK push ref
	Status       string `json:"status"`
}

const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
	StatusFailed     = "failed"
	StatusRefunded   = "refunded"
)
