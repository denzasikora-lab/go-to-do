package todo

import "time"

// Todo is a task row owned by an app user.
type Todo struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Status      Status
	Priority    Priority
	DueAt       *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
