package auth

import "time"

// TokenPair is returned on successful login or token refresh.
type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// Session represents a refresh-token row stored in the DB / cache.
type Session struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	UserAgent    string    `db:"user_agent"`
	ClientIP     string    `db:"client_ip"`
	IsBlocked    bool      `db:"is_blocked"`
	ExpiresAt    time.Time `db:"expires_at"`
	CreatedAt    time.Time `db:"created_at"`
}

// AuthUser carries the minimal user fields the auth service needs.
// The users package satisfies the UserLookup interface using this type,
// keeping the dependency one-directional (no import cycle).
type AuthUser struct {
	ID           string
	Email        string
	Role         string
	PasswordHash string
	IsActive     bool
}

// LoginInput carries credentials submitted by the client.
type LoginInput struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// RegisterInput carries sign-up data.
type RegisterInput struct {
	Name     string `json:"name"     validate:"required,min=2,max=100"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

// RefreshInput carries the refresh token.
type RefreshInput struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
