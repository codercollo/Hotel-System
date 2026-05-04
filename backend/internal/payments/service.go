package payments

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/internal/orders"
	"github.com/codercollo/hotel-system/backend/internal/payments/providers"
	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
)

// Service orchestrates payment flows across providers.
type Service struct {
	repo      Repository
	orderRepo orders.Repository
	providers map[string]providers.Provider
	broker    events.Broker
}

// NewService constructs a payments Service with a map of registered providers.
func NewService(
	repo Repository,
	orderRepo orders.Repository,
	providerList []providers.Provider,
	broker events.Broker,
) *Service {
	pm := make(map[string]providers.Provider, len(providerList))
	for _, p := range providerList {
		pm[p.Name()] = p
	}
	return &Service{
		repo:      repo,
		orderRepo: orderRepo,
		providers: pm,
		broker:    broker,
	}
}

// Initiate begins a payment for an order using the requested provider.
func (s *Service) Initiate(ctx context.Context, req InitiateRequest, userID string) (*InitiateResponse, error) {
	order, err := s.orderRepo.FindByID(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}
	if order.UserID != userID {
		return nil, apierror.ErrForbidden
	}

	// Prevent double-payment on already-completed orders.
	if order.Status == orders.StatusCompleted || order.Status == orders.StatusCancelled {
		return nil, apierror.BadRequest(fmt.Sprintf("order is already %s", order.Status))
	}

	provider, ok := s.providers[req.Provider]
	if !ok {
		return nil, apierror.BadRequest(fmt.Sprintf("unsupported provider: %s", req.Provider))
	}

	payment, err := s.repo.Create(ctx, &Payment{
		OrderID:  req.OrderID,
		Provider: req.Provider,
		Amount:   order.Total,
		Currency: order.Currency,
	})
	if err != nil {
		return nil, err
	}

	result, err := provider.Initiate(ctx, providers.InitiateParams{
		PaymentID: payment.ID,
		OrderID:   req.OrderID,
		Amount:    payment.Amount,
		Currency:  payment.Currency,
		Phone:     req.Phone,
	})
	if err != nil {
		_, _ = s.repo.UpdateStatus(ctx, payment.ID, StatusFailed, nil)
		return nil, apierror.Internal(err)
	}

	_, _ = s.repo.UpdateStatus(ctx, payment.ID, result.Status, &result.ProviderRef)

	ev, _ := events.NewEvent(idgen.NewULID(), events.TopicPaymentInitiated, payment)
	_ = s.broker.Publish(ctx, ev)

	return &InitiateResponse{
		PaymentID:    payment.ID,
		CheckoutURL:  result.CheckoutURL,
		CheckoutCode: result.CheckoutCode,
		Status:       result.Status,
	}, nil
}

// HandleWebhook processes an inbound provider webhook.
func (s *Service) HandleWebhook(ctx context.Context, r *http.Request, providerName string) error {
	provider, ok := s.providers[providerName]
	if !ok {
		return apierror.BadRequest("unknown provider: " + providerName)
	}

	body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		return apierror.BadRequest("failed to read webhook body")
	}

	// Support both generic and Stripe-specific signature headers.
	sig := r.Header.Get("X-Signature")
	if sig == "" {
		sig = r.Header.Get("Stripe-Signature")
	}

	providerRef, status, err := provider.VerifyWebhook(providers.WebhookPayload{
		Body:      body,
		Signature: sig,
		Provider:  providerName,
	})
	if err != nil {
		return apierror.BadRequest(err.Error())
	}

	// FIX: look up payment by providerRef, not by orderID.
	// The original code called FindByOrderID(ctx, providerRef) which was wrong:
	// providerRef is a provider transaction ID, not an order ID.
	payment, err := s.repo.FindByProviderRef(ctx, providerRef)
	if err != nil {
		// Idempotent — unknown ref means nothing to update.
		return nil
	}

	// Persist the raw webhook body for audit / debugging.
	var webhookBody map[string]any
	_ = json.Unmarshal(body, &webhookBody)
	_ = s.repo.StoreWebhookPayload(ctx, payment.ID, webhookBody)

	updated, err := s.repo.UpdateStatus(ctx, payment.ID, status, &providerRef)
	if err != nil {
		return err
	}

	// FIX: propagate payment outcome to the order so order.status stays in sync.
	if err := s.syncOrderStatus(ctx, updated); err != nil {
		// Log but don't fail the webhook — payment is already recorded.
		_ = err
	}

	// Publish the appropriate domain event.
	topic := events.TopicPaymentCompleted
	if status == StatusFailed {
		topic = events.TopicPaymentFailed
	}
	ev, _ := events.NewEvent(idgen.NewULID(), topic, updated)
	_ = s.broker.Publish(ctx, ev)

	return nil
}

// syncOrderStatus drives the order to its next status after a payment outcome.
//
// orders.StatusFailed does not exist in the domain — the closest state after a
// failed payment is leaving the order "pending" so the user can retry.
// Swap the StatusFailed case to orders.StatusCancelled if you prefer to close it.
func (s *Service) syncOrderStatus(ctx context.Context, p *Payment) error {
	var orderStatus string
	switch p.Status {
	case StatusCompleted:
		orderStatus = orders.StatusConfirmed
	case StatusFailed:
		// No StatusFailed constant in orders domain. Leave pending so user can retry.
		return nil
	default:
		return nil
	}

	// UpdateStatus returns (*Order, error) — discard the order value.
	_, err := s.orderRepo.UpdateStatus(ctx, p.OrderID, orderStatus)
	return err
}

// Get retrieves a single payment by ID.
func (s *Service) Get(ctx context.Context, id string) (*Payment, error) {
	return s.repo.FindByID(ctx, id)
}

// ListForOrder returns all payments for a given order.
func (s *Service) ListForOrder(ctx context.Context, orderID string) ([]*Payment, error) {
	return s.repo.FindByOrderID(ctx, orderID)
}
