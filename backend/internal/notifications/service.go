package notifications

import (
	"context"

	"github.com/codercollo/hotel-system/backend/internal/notifications/channels"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
	"github.com/codercollo/hotel-system/backend/pkg/pagination"
)

// Service handles notification creation and delivery.
type Service struct {
	repo     Repository
	channels map[string]channels.Channel
}

// NewService creates a notifications Service with registered channel adapters.
func NewService(repo Repository, channelList []channels.Channel) *Service {
	cm := make(map[string]channels.Channel, len(channelList))
	for _, c := range channelList {
		cm[c.Name()] = c
	}
	return &Service{repo: repo, channels: cm}
}

// Send stores an in-app notification and delivers it via the requested channel.
func (s *Service) Send(ctx context.Context, input SendInput) (*Notification, error) {
	n, err := s.repo.Create(ctx, &Notification{
		UserID:   input.UserID,
		Type:     input.Type,
		Channel:  input.Channel,
		Title:    input.Title,
		Body:     input.Body,
		Metadata: input.Metadata,
	})
	if err != nil {
		return nil, err
	}

	// Attempt delivery via the requested channel adapter (non-blocking on failure)
	if ch, ok := s.channels[input.Channel]; ok {
		go func() {
			if err := ch.Send(ctx, channels.Message{
				Subject: input.Title,
				Body:    input.Body,
			}); err != nil {
				logger.Error(err, "notification channel delivery failed: "+input.Channel)
			}
		}()
	}

	return n, nil
}

// List returns paginated in-app notifications for a user.
func (s *Service) List(ctx context.Context, userID string, p pagination.Params) ([]*Notification, int, error) {
	return s.repo.ListForUser(ctx, userID, p.Limit, p.Offset)
}

// MarkRead marks a single notification as read.
func (s *Service) MarkRead(ctx context.Context, id, userID string) error {
	return s.repo.MarkRead(ctx, id, userID)
}

// MarkAllRead marks all notifications as read for a user.
func (s *Service) MarkAllRead(ctx context.Context, userID string) error {
	return s.repo.MarkAllRead(ctx, userID)
}
