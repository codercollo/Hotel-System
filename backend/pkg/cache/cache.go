// Package cache defines the Cache interface and provides a Redis implementation.
// An in-memory stub is also available for testing without Redis.
package cache

import (
	"context"
	"time"
)

// Cache is the generic key-value cache abstraction used across modules.
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}
