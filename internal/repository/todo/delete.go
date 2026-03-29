package todo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
)

// Delete removes a todo owned by the app user.
func Delete(ctx context.Context, pool *pgxpool.Pool, userPK, todoID int64) error {
	const q = `DELETE FROM todos WHERE id = $1 AND user_id = $2`
	tag, err := pool.Exec(ctx, q, todoID, userPK)
	if err != nil {
		return fmt.Errorf("delete todo: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return dtodo.ErrNotFound
	}
	return nil
}
