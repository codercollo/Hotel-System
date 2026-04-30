package events

import (
	"context"
	"sync"

	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

// InMemoryBroker is a synchronous, in-process event bus for development and
// unit testing. Handlers are called sequentially in the goroutine that
// calls Publish. Not suitable for production.
type InMemoryBroker struct {
	mu       sync.RWMutex
	handlers map[Topic]map[uint64]Handler
	nextID   uint64
}

// NewInMemoryBroker creates a ready-to-use in-memory broker.
func NewInMemoryBroker() *InMemoryBroker {
	return &InMemoryBroker{
		handlers: make(map[Topic]map[uint64]Handler),
	}
}

// Publish calls every handler registered for event.Topic synchronously.
// Handler errors are logged but do not abort delivery to other handlers.
func (b *InMemoryBroker) Publish(ctx context.Context, event Event) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	handlers, ok := b.handlers[event.Topic]
	if !ok {
		return nil
	}

	for _, h := range handlers {
		if err := h(ctx, event); err != nil {
			logger.Error(err, "inmemory broker: handler error for topic "+string(event.Topic))
		}
	}
	return nil
}

// Subscribe registers handler for topic and returns a cancel function.
func (b *InMemoryBroker) Subscribe(topic Topic, handler Handler) (cancel func()) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.handlers[topic]; !ok {
		b.handlers[topic] = make(map[uint64]Handler)
	}

	b.nextID++
	id := b.nextID
	b.handlers[topic][id] = handler

	return func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		delete(b.handlers[topic], id)
	}
}

// Close is a no-op for the in-memory broker.
func (b *InMemoryBroker) Close() error { return nil }

// Compile-time interface check.
var _ Broker = (*InMemoryBroker)(nil)
