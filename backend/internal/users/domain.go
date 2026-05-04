package users

import "time"

// User is the canonical user entity. Role is a plain string; RBAC details
// live in the roles module. The Password field is NEVER serialised to JSON.
type User struct {
	ID        string         `json:"id"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Role      string         `json:"role"`
	IsActive  bool           `json:"is_active"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`

	Password string `json:"-"` // bcrypt hash — never exposed
}

// CreateUserInput is the validated DTO for user creation.
type CreateUserInput struct {
	Name     string
	Email    string
	Password string // already hashed
	Role     string
}

// UpdateUserInput is the validated DTO for profile updates.
type UpdateUserInput struct {
	Name     *string
	Email    *string
	IsActive *bool
	Metadata map[string]any
}
