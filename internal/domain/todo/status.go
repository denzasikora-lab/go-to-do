package todo

// Status enumerates lifecycle columns stored in todos.status.
type Status string

const (
	StatusOpen     Status = "open"
	StatusDone     Status = "done"
	StatusArchived Status = "archived"
)
