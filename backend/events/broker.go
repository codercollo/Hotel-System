// Package events provides the event-driven messaging layer for the platform.
// Modules publish events using the Broker interface; handlers subscribe and
// react asynchronously. This decouples domain modules from each other.
package events

import "context"

// Handler is a function that processes a received event.
type Handler func(ctx context.Context, event Event) error

// Broker is the messaging abstraction. Implementations include RabbitMQ
// (production) and InMemory (development/testing).
type Broker interface {
	// Publish sends an event to all subscribers of its topic.
	Publish(ctx context.Context, event Event) error

	// Subscribe registers a handler for a given topic.
	// The returned cancel func deregisters the handler.
	Subscribe(topic Topic, handler Handler) (cancel func())

	// Close gracefully shuts down the broker.
	Close() error
}
