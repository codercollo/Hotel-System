package events

import (
	"encoding/json"
	"time"
)

// Topic is a typed string identifying an event channel.
type Topic string

const (
	// User domain
	TopicUserRegistered Topic = "user.registered"
	TopicUserUpdated    Topic = "user.updated"
	TopicUserDeleted    Topic = "user.deleted"

	// Item domain
	TopicItemCreated Topic = "item.created"
	TopicItemUpdated Topic = "item.updated"
	TopicItemDeleted Topic = "item.deleted"

	// Order domain
	TopicOrderPlaced    Topic = "order.placed"
	TopicOrderConfirmed Topic = "order.confirmed"
	TopicOrderCancelled Topic = "order.cancelled"
	TopicOrderCompleted Topic = "order.completed"

	// Payment domain
	TopicPaymentInitiated Topic = "payment.initiated"
	TopicPaymentCompleted Topic = "payment.completed"
	TopicPaymentFailed    Topic = "payment.failed"
	TopicPaymentRefunded  Topic = "payment.refunded"

	// Notification domain
	TopicNotificationSend Topic = "notification.send"
)

// Event is the canonical envelope for all async messages.
type Event struct {
	ID        string          `json:"id"`
	Topic     Topic           `json:"topic"`
	Payload   json.RawMessage `json:"payload"`
	OccuredAt time.Time       `json:"occurred_at"`
}

// Decode unmarshals the event payload into dst.
func (e Event) Decode(dst any) error {
	return json.Unmarshal(e.Payload, dst)
}

// NewEvent constructs an Event, marshalling payload to JSON.
func NewEvent(id string, topic Topic, payload any) (Event, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return Event{}, err
	}
	return Event{
		ID:        id,
		Topic:     topic,
		Payload:   b,
		OccuredAt: time.Now(),
	}, nil
}
