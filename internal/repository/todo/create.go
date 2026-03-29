package todo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
)

// Create inserts a new todo for the given app user primary key.
func Create(ctx context.Context, pool *pgxpool.Pool, userPK int64, title, description string, priority dtodo.Priority) (*dtodo.Todo, error) {
	const q = `
INSERT INTO todos (user_id, title, description, priority)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, title, description, status, priority, due_at, created_at, updated_at
`
	row := pool.QueryRow(ctx, q, userPK, title, description, string(priority))
	return scanTodo(row)
}
