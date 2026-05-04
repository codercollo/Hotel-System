package providers

import (
	"context"
	"fmt"
)

// PaystackProvider implements Provider for Paystack Charge API.
type PaystackProvider struct {
	secretKey     string
	webhookSecret string
}

// NewPaystackProvider constructs a Paystack provider.
func NewPaystackProvider(secretKey, webhookSecret string) *PaystackProvider {
	return &PaystackProvider{secretKey: secretKey, webhookSecret: webhookSecret}
}

func (p *PaystackProvider) Name() string { return "paystack" }

// Initiate creates a Paystack transaction. Full implementation in Phase 5.
func (p *PaystackProvider) Initiate(_ context.Context, params InitiateParams) (*InitiateResult, error) {
	return &InitiateResult{
		ProviderRef: fmt.Sprintf("PSK-STUB-%s", params.PaymentID),
		CheckoutURL: fmt.Sprintf("https://checkout.paystack.com/stub/%s", params.PaymentID),
		Status:      "pending",
	}, nil
}

// VerifyWebhook validates a Paystack webhook. Full implementation in Phase 5.
func (p *PaystackProvider) VerifyWebhook(payload WebhookPayload) (string, string, error) {
	return "", "completed", nil
}
