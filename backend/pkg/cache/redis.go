package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache implements Cache using go-redis.
type RedisCache struct {
	client *redis.Client
}

// NewRedis creates a RedisCache from a connection URL.
// URL format: redis://:password@host:port/db
func NewRedis(url string) (*RedisCache, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("cache: invalid redis URL: %w", err)
	}

	client := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("cache: redis ping failed: %w", err)
	}

	return &RedisCache{client: client}, nil
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (c *RedisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	var v string
	switch t := value.(type) {
	case string:
		v = t
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		v = string(b)
	}
	return c.client.Set(ctx, key, v, ttl).Err()
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	n, err := c.client.Exists(ctx, key).Result()
	return n > 0, err
}

// Close shuts down the Redis connection pool.
func (c *RedisCache) Close() error {
	return c.client.Close()
}
