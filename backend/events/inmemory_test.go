package events_test

import (
	"context"
	"testing"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	logger.Init("error", "json")
}

func TestInMemoryBroker_PublishSubscribe(t *testing.T) {
	broker := events.NewInMemoryBroker()

	received := make([]events.Event, 0)
	cancel := broker.Subscribe(events.TopicOrderPlaced, func(_ context.Context, e events.Event) error {
		received = append(received, e)
		return nil
	})
	defer cancel()

	event, err := events.NewEvent(idgen.NewULID(), events.TopicOrderPlaced, map[string]string{"order_id": "ord-1"})
	require.NoError(t, err)

	err = broker.Publish(context.Background(), event)
	require.NoError(t, err)

	assert.Len(t, received, 1)
	assert.Equal(t, events.TopicOrderPlaced, received[0].Topic)
}

func TestInMemoryBroker_Cancel(t *testing.T) {
	broker := events.NewInMemoryBroker()

	count := 0
	cancel := broker.Subscribe(events.TopicItemCreated, func(_ context.Context, _ events.Event) error {
		count++
		return nil
	})

	event, _ := events.NewEvent(idgen.NewULID(), events.TopicItemCreated, nil)
	_ = broker.Publish(context.Background(), event)
	assert.Equal(t, 1, count)

	cancel()
	_ = broker.Publish(context.Background(), event)
	assert.Equal(t, 1, count, "handler should not be called after cancel")
}

func TestInMemoryBroker_MultipleSubscribers(t *testing.T) {
	broker := events.NewInMemoryBroker()

	countA, countB := 0, 0
	cancelA := broker.Subscribe(events.TopicUserRegistered, func(_ context.Context, _ events.Event) error {
		countA++
		return nil
	})
	cancelB := broker.Subscribe(events.TopicUserRegistered, func(_ context.Context, _ events.Event) error {
		countB++
		return nil
	})
	defer cancelA()
	defer cancelB()

	event, _ := events.NewEvent(idgen.NewULID(), events.TopicUserRegistered, nil)
	require.NoError(t, broker.Publish(context.Background(), event))

	assert.Equal(t, 1, countA)
	assert.Equal(t, 1, countB)
}

func TestInMemoryBroker_NoSubscribers(t *testing.T) {
	broker := events.NewInMemoryBroker()
	event, _ := events.NewEvent(idgen.NewULID(), events.TopicPaymentCompleted, nil)
	// Should not panic or error with no subscribers.
	assert.NoError(t, broker.Publish(context.Background(), event))
}
