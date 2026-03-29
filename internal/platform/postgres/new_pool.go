package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPool configures a pgx pool suitable for the bot workload.
func NewPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse dsn: %w", err)
	}
	cfg.MaxConns = 12
	cfg.MinConns = 1
	cfg.HealthCheckPeriod = 45 * time.Second
	cfg.MaxConnLifetime = 55 * time.Minute
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgx pool: %w", err)
	}
	return pool, nil
}
