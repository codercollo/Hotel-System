package providers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/codercollo/hotel-system/backend/pkg/crypto"
)

// StripeProvider implements Provider for Stripe PaymentIntents.
type StripeProvider struct {
	secretKey     string
	webhookSecret string
}

// NewStripeProvider constructs a Stripe provider.
func NewStripeProvider(secretKey, webhookSecret string) *StripeProvider {
	return &StripeProvider{secretKey: secretKey, webhookSecret: webhookSecret}
}

func (p *StripeProvider) Name() string { return "stripe" }

// Initiate creates a Stripe PaymentIntent and returns the client secret URL.
// Full implementation in Phase 5.
func (p *StripeProvider) Initiate(_ context.Context, params InitiateParams) (*InitiateResult, error) {
	return &InitiateResult{
		ProviderRef: fmt.Sprintf("pi_stub_%s", params.PaymentID),
		CheckoutURL: fmt.Sprintf("https://checkout.stripe.com/stub/%s", params.PaymentID),
		Status:      "pending",
	}, nil
}

// VerifyWebhook validates a Stripe webhook using the endpoint secret.
func (p *StripeProvider) VerifyWebhook(payload WebhookPayload) (string, string, error) {
	if !crypto.Verify(p.webhookSecret, string(payload.Body), payload.Signature) {
		return "", "", fmt.Errorf("stripe: invalid webhook signature")
	}

	var event struct {
		Type string `json:"type"`
		Data struct {
			Object struct {
				ID     string `json:"id"`
				Status string `json:"status"`
			} `json:"object"`
		} `json:"data"`
	}
	if err := json.Unmarshal(payload.Body, &event); err != nil {
		return "", "", fmt.Errorf("stripe: malformed webhook: %w", err)
	}

	status := "failed"
	switch event.Type {
	case "payment_intent.succeeded":
		status = "completed"
	case "payment_intent.processing":
		status = "processing"
	}

	return event.Data.Object.ID, status, nil
}
