package todo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
)

// UpdateDescription updates body text for a todo row owned by the user.
func UpdateDescription(ctx context.Context, pool *pgxpool.Pool, userPK, todoID int64, description string) error {
	const q = `
UPDATE todos SET description = $3, updated_at = now()
WHERE id = $1 AND user_id = $2
`
	tag, err := pool.Exec(ctx, q, todoID, userPK, description)
	if err != nil {
		return fmt.Errorf("update description: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return dtodo.ErrNotFound
	}
	return nil
}
