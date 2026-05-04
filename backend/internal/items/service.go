package items

import (
	"context"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/codercollo/hotel-system/backend/pkg/pagination"
)

// Service implements item management logic.
type Service struct {
	repo   Repository
	broker events.Broker
}

// NewService creates an items Service.
func NewService(repo Repository, broker events.Broker) *Service {
	return &Service{repo: repo, broker: broker}
}

func (s *Service) Get(ctx context.Context, id string) (*Item, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) List(ctx context.Context, p pagination.Params, status string) ([]*Item, int, error) {
	return s.repo.List(ctx, ListFilters{Status: status, Limit: p.Limit, Offset: p.Offset})
}

func (s *Service) Create(ctx context.Context, input CreateItemInput) (*Item, error) {
	if input.Currency == "" {
		input.Currency = "USD"
	}
	item, err := s.repo.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	ev, _ := events.NewEvent(idgen.NewULID(), events.TopicItemCreated, item)
	_ = s.broker.Publish(ctx, ev)

	return item, nil
}

func (s *Service) Update(ctx context.Context, id string, input UpdateItemInput) (*Item, error) {
	if input.Status != nil {
		switch *input.Status {
		case StatusActive, StatusInactive, StatusArchived:
		default:
			return nil, apierror.BadRequest("invalid status value")
		}
	}

	item, err := s.repo.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}

	ev, _ := events.NewEvent(idgen.NewULID(), events.TopicItemUpdated, item)
	_ = s.broker.Publish(ctx, ev)

	return item, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	ev, _ := events.NewEvent(idgen.NewULID(), events.TopicItemDeleted, map[string]string{"id": id})
	_ = s.broker.Publish(ctx, ev)

	return nil
}

func (s *Service) Search(ctx context.Context, query string, p pagination.Params) ([]*Item, int, error) {
	if query == "" {
		return s.repo.List(ctx, ListFilters{Status: StatusActive, Limit: p.Limit, Offset: p.Offset})
	}
	return s.repo.Search(ctx, query, p.Limit, p.Offset)
}
