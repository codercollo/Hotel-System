package orders

import "time"

// Order is the generic order entity — maps to reservations, bookings, purchases.
type Order struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	Status    string         `json:"status"` // pending|confirmed|processing|completed|cancelled
	Total     int64          `json:"total"`
	Currency  string         `json:"currency"`
	Notes     string         `json:"notes"`
	Metadata  map[string]any `json:"metadata"`
	Items     []OrderItem    `json:"items"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// OrderItem is a snapshot of an item at the time of order creation.
type OrderItem struct {
	ID       string `json:"id"`
	OrderID  string `json:"order_id"`
	ItemID   string `json:"item_id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
	Subtotal int64  `json:"subtotal"`
}

// CreateOrderInput holds validated data for order placement.
type CreateOrderInput struct {
	UserID   string
	Items    []OrderItemInput
	Notes    string
	Metadata map[string]any
}

// OrderItemInput references an item and desired quantity.
type OrderItemInput struct {
	ItemID   string `json:"item_id"  validate:"required"`
	Quantity int    `json:"quantity" validate:"gte=1"`
}

const (
	StatusPending    = "pending"
	StatusConfirmed  = "confirmed"
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
	StatusCancelled  = "cancelled"
)
