package auth

import "context"

// Repository defines storage operations for auth sessions.
// Implemented by postgresRepo (persistent) or redisTokenRepo (cache-backed).
type Repository interface {
	CreateSession(ctx context.Context, s Session) error
	GetSession(ctx context.Context, refreshToken string) (*Session, error)
	BlockSession(ctx context.Context, refreshToken string) error
	DeleteExpiredSessions(ctx context.Context) error
}
