package providers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/codercollo/hotel-system/backend/pkg/crypto"
)

// MPesaProvider implements the Provider interface for Safaricom M-Pesa STK Push.
// Configuration is passed at construction; secrets come from environment variables.
type MPesaProvider struct {
	consumerKey    string
	consumerSecret string
	shortCode      string
	passkey        string
	callbackURL    string
	webhookSecret  string
}

// NewMPesaProvider constructs an M-Pesa provider.
func NewMPesaProvider(consumerKey, consumerSecret, shortCode, passkey, callbackURL, webhookSecret string) *MPesaProvider {
	return &MPesaProvider{
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
		shortCode:      shortCode,
		passkey:        passkey,
		callbackURL:    callbackURL,
		webhookSecret:  webhookSecret,
	}
}

func (p *MPesaProvider) Name() string { return "mpesa" }

// Initiate sends an STK Push to the customer's phone.
// Full implementation in Phase 5 — this stub records the attempt.
func (p *MPesaProvider) Initiate(_ context.Context, params InitiateParams) (*InitiateResult, error) {
	// Phase 5: call Safaricom OAuth → STK Push API
	// For now, return a deterministic stub result so the flow compiles end-to-end.
	return &InitiateResult{
		ProviderRef:  fmt.Sprintf("MPESA-%s", params.PaymentID),
		CheckoutCode: fmt.Sprintf("ws_CO_%s", params.PaymentID),
		Status:       "processing",
	}, nil
}

// VerifyWebhook validates an M-Pesa callback using HMAC-SHA256.
func (p *MPesaProvider) VerifyWebhook(payload WebhookPayload) (string, string, error) {
	if !crypto.Verify(p.webhookSecret, string(payload.Body), payload.Signature) {
		return "", "", fmt.Errorf("mpesa: invalid webhook signature")
	}

	var body struct {
		Body struct {
			StkCallback struct {
				ResultCode        int    `json:"ResultCode"`
				CheckoutRequestID string `json:"CheckoutRequestID"`
			} `json:"stkCallback"`
		} `json:"Body"`
	}

	if err := json.Unmarshal(payload.Body, &body); err != nil {
		return "", "", fmt.Errorf("mpesa: malformed webhook: %w", err)
	}

	ref := body.Body.StkCallback.CheckoutRequestID
	status := "completed"
	if body.Body.StkCallback.ResultCode != 0 {
		status = "failed"
	}

	return ref, status, nil
}
