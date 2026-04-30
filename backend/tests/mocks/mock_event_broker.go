// Package mocks provides mock implementations for use in tests.
package mocks

import (
	"context"
	"sync"

	"github.com/codercollo/hotel-system/backend/events"
)

// MockBroker records published events and allows assertions in tests.
type MockBroker struct {
	mu        sync.Mutex
	Published []events.Event
}

func NewMockBroker() *MockBroker {
	return &MockBroker{}
}

func (m *MockBroker) Publish(_ context.Context, event events.Event) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Published = append(m.Published, event)
	return nil
}

func (m *MockBroker) Subscribe(_ events.Topic, _ events.Handler) (cancel func()) {
	return func() {}
}

func (m *MockBroker) Close() error { return nil }

// EventsForTopic returns all published events matching the given topic.
func (m *MockBroker) EventsForTopic(topic events.Topic) []events.Event {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []events.Event
	for _, e := range m.Published {
		if e.Topic == topic {
			out = append(out, e)
		}
	}
	return out
}

// Reset clears all recorded events.
func (m *MockBroker) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Published = nil
}

// Compile-time interface check.
var _ events.Broker = (*MockBroker)(nil)
