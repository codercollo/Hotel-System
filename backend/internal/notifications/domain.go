package notifications

import "time"

// Notification is a user-facing message delivered via one or more channels.
type Notification struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	Type      string         `json:"type"`    // order_confirmed|payment_completed|system
	Channel   string         `json:"channel"` // in_app|email|sms|push
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	IsRead    bool           `json:"is_read"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	ReadAt    *time.Time     `json:"read_at,omitempty"`
}

// SendInput is the data required to deliver a notification.
type SendInput struct {
	UserID   string
	Type     string
	Channel  string
	Title    string
	Body     string
	Metadata map[string]any
}
