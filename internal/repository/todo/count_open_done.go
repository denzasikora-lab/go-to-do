package todo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// CountOpenDone returns how many todos are open vs done for dashboards.
func CountOpenDone(ctx context.Context, pool *pgxpool.Pool, userPK int64) (openN, doneN int64, err error) {
	const q = `
SELECT
  COALESCE(SUM(CASE WHEN status = 'open' THEN 1 ELSE 0 END), 0),
  COALESCE(SUM(CASE WHEN status = 'done' THEN 1 ELSE 0 END), 0)
FROM todos
WHERE user_id = $1 AND status <> 'archived'
`
	row := pool.QueryRow(ctx, q, userPK)
	if err = row.Scan(&openN, &doneN); err != nil {
		return 0, 0, fmt.Errorf("count todos: %w", err)
	}
	return openN, doneN, nil
}
