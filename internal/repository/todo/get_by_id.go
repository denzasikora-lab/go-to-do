package todo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
)

// GetByID returns a todo when it belongs to the app user.
func GetByID(ctx context.Context, pool *pgxpool.Pool, userPK, todoID int64) (*dtodo.Todo, error) {
	const q = `
SELECT id, user_id, title, description, status, priority, due_at, created_at, updated_at
FROM todos
WHERE id = $1 AND user_id = $2
`
	row := pool.QueryRow(ctx, q, todoID, userPK)
	t, err := scanTodo(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, dtodo.ErrNotFound
		}
		return nil, fmt.Errorf("get todo: %w", err)
	}
	return t, nil
}
