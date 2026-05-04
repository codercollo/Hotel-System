package orders

import (
	"context"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/internal/items"
	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/codercollo/hotel-system/backend/pkg/pagination"
)

// Service implements order workflow logic.
type Service struct {
	repo     Repository
	itemRepo items.Repository
	broker   events.Broker
}

// NewService creates an orders Service.
func NewService(repo Repository, itemRepo items.Repository, broker events.Broker) *Service {
	return &Service{repo: repo, itemRepo: itemRepo, broker: broker}
}

func (s *Service) Get(ctx context.Context, id, userID, role string) (*Order, error) {
	order, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if role != "admin" && order.UserID != userID {
		return nil, apierror.ErrForbidden
	}
	return order, nil
}

func (s *Service) ListForUser(ctx context.Context, userID string, p pagination.Params) ([]*Order, int, error) {
	return s.repo.ListForUser(ctx, userID, p.Limit, p.Offset)
}

func (s *Service) ListAll(ctx context.Context, p pagination.Params) ([]*Order, int, error) {
	return s.repo.ListAll(ctx, p.Limit, p.Offset)
}

func (s *Service) Place(ctx context.Context, input CreateOrderInput) (*Order, error) {
	if len(input.Items) == 0 {
		return nil, apierror.BadRequest("order must contain at least one item")
	}

	var (
		orderItems []OrderItem
		total      int64
	)

	for _, oi := range input.Items {
		item, err := s.itemRepo.FindByID(ctx, oi.ItemID)
		if err != nil {
			return nil, err
		}
		if item.Stock < oi.Quantity {
			return nil, apierror.BadRequest("insufficient stock for item: " + item.Name)
		}

		subtotal := item.Price * int64(oi.Quantity)
		total += subtotal

		orderItems = append(orderItems, OrderItem{
			ItemID:   oi.ItemID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: oi.Quantity,
			Subtotal: subtotal,
		})
	}

	order := &Order{
		UserID:   input.UserID,
		Total:    total,
		Currency: "USD",
		Notes:    input.Notes,
		Metadata: input.Metadata,
	}
	if order.Metadata == nil {
		order.Metadata = map[string]any{}
	}

	created, err := s.repo.Create(ctx, order, orderItems)
	if err != nil {
		return nil, err
	}

	ev, _ := events.NewEvent(idgen.NewULID(), events.TopicOrderPlaced, created)
	_ = s.broker.Publish(ctx, ev)

	return created, nil
}

func (s *Service) UpdateStatus(ctx context.Context, id, status, role string) (*Order, error) {
	validTransitions := map[string][]string{
		StatusPending:    {StatusConfirmed, StatusCancelled},
		StatusConfirmed:  {StatusProcessing, StatusCancelled},
		StatusProcessing: {StatusCompleted, StatusCancelled},
	}

	order, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	allowed, ok := validTransitions[order.Status]
	if !ok {
		return nil, apierror.BadRequest("order cannot be updated from status: " + order.Status)
	}
	valid := false
	for _, s := range allowed {
		if s == status {
			valid = true
			break
		}
	}
	if !valid {
		return nil, apierror.BadRequest("invalid status transition: " + order.Status + " → " + status)
	}

	updated, err := s.repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return nil, err
	}

	topicMap := map[string]events.Topic{
		StatusConfirmed: events.TopicOrderConfirmed,
		StatusCompleted: events.TopicOrderCompleted,
		StatusCancelled: events.TopicOrderCancelled,
	}
	if topic, ok := topicMap[status]; ok {
		ev, _ := events.NewEvent(idgen.NewULID(), topic, updated)
		_ = s.broker.Publish(ctx, ev)
	}

	return updated, nil
}

func (s *Service) Cancel(ctx context.Context, id, userID, role string) (*Order, error) {
	order, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if role != "admin" && order.UserID != userID {
		return nil, apierror.ErrForbidden
	}
	if order.Status == StatusCompleted {
		return nil, apierror.BadRequest("completed orders cannot be cancelled")
	}
	return s.repo.UpdateStatus(ctx, id, StatusCancelled)
}
