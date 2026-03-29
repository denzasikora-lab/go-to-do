package todo

import (
	"fmt"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
)

func scanTodo(row pgxRow) (*dtodo.Todo, error) {
	var t dtodo.Todo
	var status, priority string
	if err := row.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &status, &priority, &t.DueAt, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, fmt.Errorf("todo scan: %w", err)
	}
	t.Status = dtodo.Status(status)
	t.Priority = dtodo.Priority(priority)
	return &t, nil
}
