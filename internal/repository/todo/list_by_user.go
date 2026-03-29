package todo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
)

// ListByUser returns recent todos for the app user, optionally filtered by status.
func ListByUser(ctx context.Context, pool *pgxpool.Pool, userPK int64, status *dtodo.Status, limit int) ([]dtodo.Todo, error) {
	if limit <= 0 || limit > 50 {
		limit = 15
	}
	var rows pgx.Rows
	var err error
	if status != nil {
		const q = `
SELECT id, user_id, title, description, status, priority, due_at, created_at, updated_at
FROM todos
WHERE user_id = $1 AND status = $2
ORDER BY updated_at DESC
LIMIT $3
`
		rows, err = pool.Query(ctx, q, userPK, string(*status), limit)
	} else {
		const q = `
SELECT id, user_id, title, description, status, priority, due_at, created_at, updated_at
FROM todos
WHERE user_id = $1 AND status <> 'archived'
ORDER BY updated_at DESC
LIMIT $2
`
		rows, err = pool.Query(ctx, q, userPK, limit)
	}
	if err != nil {
		return nil, fmt.Errorf("list todos: %w", err)
	}
	defer rows.Close()
	var out []dtodo.Todo
	for rows.Next() {
		t, sErr := scanTodo(rows)
		if sErr != nil {
			return nil, sErr
		}
		out = append(out, *t)
	}
	return out, rows.Err()
}
