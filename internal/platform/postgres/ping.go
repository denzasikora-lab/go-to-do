package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Ping validates database connectivity.
func Ping(ctx context.Context, pool *pgxpool.Pool) error {
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping: %w", err)
	}
	return nil
}
