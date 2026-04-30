package events

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/codercollo/hotel-system/backend/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQBroker implements Broker using RabbitMQ via amqp091-go.
// Each topic maps to a fanout exchange so multiple consumers can receive the
// same event independently.
type RabbitMQBroker struct {
	conn     *amqp.Connection
	exchange string
	mu       sync.Mutex
	channels []*amqp.Channel
}

// NewRabbitMQBroker dials RabbitMQ and returns a ready broker.
func NewRabbitMQBroker(url, exchange string) (*RabbitMQBroker, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: dial failed: %w", err)
	}

	b := &RabbitMQBroker{conn: conn, exchange: exchange}

	// Declare the top-level topic exchange once.
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("rabbitmq: channel open failed: %w", err)
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil); err != nil {
		return nil, fmt.Errorf("rabbitmq: exchange declare failed: %w", err)
	}

	return b, nil
}

// Publish serialises event and routes it to the exchange using the topic as
// the routing key.
func (b *RabbitMQBroker) Publish(ctx context.Context, event Event) error {
	ch, err := b.conn.Channel()
	if err != nil {
		return fmt.Errorf("rabbitmq: channel: %w", err)
	}
	defer ch.Close()

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return ch.PublishWithContext(ctx, b.exchange, string(event.Topic), false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}

// Subscribe binds a queue to the exchange for topic and starts a consumer
// goroutine that calls handler for each delivery.
func (b *RabbitMQBroker) Subscribe(topic Topic, handler Handler) (cancel func()) {
	ch, err := b.conn.Channel()
	if err != nil {
		logger.Error(err, "rabbitmq: subscribe channel open failed")
		return func() {}
	}

	q, err := ch.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		logger.Error(err, "rabbitmq: queue declare failed")
		ch.Close()
		return func() {}
	}

	if err := ch.QueueBind(q.Name, string(topic), b.exchange, false, nil); err != nil {
		logger.Error(err, "rabbitmq: queue bind failed")
		ch.Close()
		return func() {}
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		logger.Error(err, "rabbitmq: consume failed")
		ch.Close()
		return func() {}
	}

	b.mu.Lock()
	b.channels = append(b.channels, ch)
	b.mu.Unlock()

	go func() {
		for msg := range msgs {
			var event Event
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				logger.Error(err, "rabbitmq: unmarshal event")
				continue
			}
			if err := handler(context.Background(), event); err != nil {
				logger.Error(err, "rabbitmq: handler error for topic "+string(topic))
			}
		}
	}()

	return func() { ch.Close() }
}

// Close shuts down the AMQP connection.
func (b *RabbitMQBroker) Close() error {
	return b.conn.Close()
}

// Compile-time interface check.
var _ Broker = (*RabbitMQBroker)(nil)
