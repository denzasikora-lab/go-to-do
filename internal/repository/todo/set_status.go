package todo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
)

// SetStatus updates lifecycle status when the row belongs to the user.
func SetStatus(ctx context.Context, pool *pgxpool.Pool, userPK, todoID int64, status dtodo.Status) error {
	const q = `
UPDATE todos
SET status = $3, updated_at = now()
WHERE id = $1 AND user_id = $2
`
	tag, err := pool.Exec(ctx, q, todoID, userPK, string(status))
	if err != nil {
		return fmt.Errorf("set status: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return dtodo.ErrNotFound
	}
	return nil
}
