package todo

import "errors"

var (
	// ErrNotFound signals a missing row for the requested identifier.
	ErrNotFound = errors.New("todo not found")
)
