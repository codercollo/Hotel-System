package items

import "time"

// Item is the generic entity model. Business-specific fields go in Metadata.
// Frontend maps: Item → Room, Product, Listing, etc.
type Item struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int64          `json:"price"` // smallest currency unit (cents)
	Currency    string         `json:"currency"`
	Status      string         `json:"status"` // active | inactive | archived
	Stock       int            `json:"stock"`
	Images      []string       `json:"images"`
	Metadata    map[string]any `json:"metadata"`
	SearchVec   string         `json:"-"`
	CreatedBy   *string        `json:"created_by,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// CreateItemInput is the validated DTO for item creation.
type CreateItemInput struct {
	Name        string
	Description string
	Price       int64
	Currency    string
	Stock       int
	Images      []string
	Metadata    map[string]any
	CreatedBy   *string
}

// UpdateItemInput is the validated DTO for item updates.
type UpdateItemInput struct {
	Name        *string
	Description *string
	Price       *int64
	Status      *string
	Stock       *int
	Images      []string
	Metadata    map[string]any
}

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusArchived = "archived"
)
