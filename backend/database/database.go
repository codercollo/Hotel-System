// Package database manages the PostgreSQL connection pool via pgx/v5.
// Call Init once at startup and pass the *pgxpool.Pool to repositories.
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/codercollo/hotel-system/backend/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Init creates and validates a pgx connection pool using the provided config.
func Init(ctx context.Context, cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("database: failed to parse DSN: %w", err)
	}

	poolCfg.MaxConns = cfg.MaxConns
	poolCfg.MinConns = cfg.MinConns
	poolCfg.MaxConnLifetime = cfg.MaxConnLifetime
	poolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("database: failed to create pool: %w", err)
	}

	if err := ping(ctx, pool); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

// ping verifies the database is reachable with a short timeout.
func ping(ctx context.Context, pool *pgxpool.Pool) error {
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		return fmt.Errorf("database: ping failed: %w", err)
	}
	return nil
}

// HealthCheck pings the database and returns an error if unreachable.
// Used by the /api/health endpoint.
func HealthCheck(ctx context.Context, pool *pgxpool.Pool) error {
	return ping(ctx, pool)
}
