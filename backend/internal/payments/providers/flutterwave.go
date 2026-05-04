package providers

import (
	"context"
	"fmt"
)

// FlutterwaveProvider implements Provider for Flutterwave Collect.
type FlutterwaveProvider struct {
	secretKey     string
	webhookSecret string
}

// NewFlutterwaveProvider constructs a Flutterwave provider.
func NewFlutterwaveProvider(secretKey, webhookSecret string) *FlutterwaveProvider {
	return &FlutterwaveProvider{secretKey: secretKey, webhookSecret: webhookSecret}
}

func (p *FlutterwaveProvider) Name() string { return "flutterwave" }

// Initiate creates a Flutterwave payment link. Full implementation in Phase 5.
func (p *FlutterwaveProvider) Initiate(_ context.Context, params InitiateParams) (*InitiateResult, error) {
	return &InitiateResult{
		ProviderRef: fmt.Sprintf("FLW-STUB-%s", params.PaymentID),
		CheckoutURL: fmt.Sprintf("https://checkout.flutterwave.com/stub/%s", params.PaymentID),
		Status:      "pending",
	}, nil
}

// VerifyWebhook validates Flutterwave webhook. Full implementation in Phase 5.
func (p *FlutterwaveProvider) VerifyWebhook(payload WebhookPayload) (string, string, error) {
	return "", "completed", nil
}
