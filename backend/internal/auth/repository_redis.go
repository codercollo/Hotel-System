package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/cache"
)

const sessionPrefix = "session:"

type redisTokenRepo struct {
	c cache.Cache
}

// NewRedisTokenRepo returns a Redis-backed Repository.
// Called in main.go as: auth.NewRedisTokenRepo(redisCache)
func NewRedisTokenRepo(c cache.Cache) Repository {
	return &redisTokenRepo{c: c}
}

func (r *redisTokenRepo) CreateSession(ctx context.Context, s Session) error {
	b, err := json.Marshal(s)
	if err != nil {
		return apierror.ErrInternal
	}
	ttl := time.Until(s.ExpiresAt)
	if ttl <= 0 {
		ttl = 7 * 24 * time.Hour
	}
	return r.c.Set(ctx, sessionKey(s.RefreshToken), string(b), ttl)
}

func (r *redisTokenRepo) GetSession(ctx context.Context, refreshToken string) (*Session, error) {
	raw, err := r.c.Get(ctx, sessionKey(refreshToken))
	if err != nil {
		return nil, ErrSessionNotFound
	}
	var s Session
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		return nil, apierror.ErrInternal
	}
	return &s, nil
}

func (r *redisTokenRepo) BlockSession(ctx context.Context, refreshToken string) error {
	s, err := r.GetSession(ctx, refreshToken)
	if err != nil {
		return nil // already gone — treat as success
	}
	s.IsBlocked = true
	b, _ := json.Marshal(s)
	ttl := time.Until(s.ExpiresAt)
	if ttl <= 0 {
		return r.c.Delete(ctx, sessionKey(refreshToken))
	}
	return r.c.Set(ctx, sessionKey(refreshToken), string(b), ttl)
}

func (r *redisTokenRepo) DeleteExpiredSessions(_ context.Context) error {
	// Redis TTL handles expiry automatically; nothing to do here.
	return nil
}

func sessionKey(token string) string {
	return fmt.Sprintf("%s%s", sessionPrefix, token)
}
