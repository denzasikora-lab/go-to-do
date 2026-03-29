package todo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
)

// UpdateTitle changes the title when the row belongs to the user.
func UpdateTitle(ctx context.Context, pool *pgxpool.Pool, userPK, todoID int64, title string) error {
	const q = `
UPDATE todos SET title = $3, updated_at = now()
WHERE id = $1 AND user_id = $2
`
	tag, err := pool.Exec(ctx, q, todoID, userPK, title)
	if err != nil {
		return fmt.Errorf("update title: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return dtodo.ErrNotFound
	}
	return nil
}
