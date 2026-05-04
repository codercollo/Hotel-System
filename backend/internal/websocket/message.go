package websocket

import "encoding/json"

// MessageType is a typed string identifying the kind of real-time message.
type MessageType string

const (
	TypeNotification  MessageType = "notification"
	TypeOrderStatus   MessageType = "order_status"
	TypePaymentStatus MessageType = "payment_status"
	TypePing          MessageType = "ping"
	TypePong          MessageType = "pong"
)

// Message is the typed envelope for every WebSocket message.
type Message struct {
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

// NewMessage creates a serialisable Message from any payload.
func NewMessage(t MessageType, payload any) (*Message, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return &Message{Type: t, Payload: b}, nil
}
