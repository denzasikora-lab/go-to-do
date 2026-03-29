package fsm

import (
	"encoding/json"
	"fmt"
)

const KeyEditTodoID = "edit_todo_id"

// PayloadInt64 reads a numeric id previously stored in JSON payload.
func PayloadInt64(payload json.RawMessage, key string) (int64, bool) {
	m := payloadMap(payload)
	v, ok := m[key]
	if !ok {
		return 0, false
	}
	switch x := v.(type) {
	case float64:
		return int64(x), true
	case json.Number:
		i, err := x.Int64()
		return i, err == nil
	default:
		return 0, false
	}
}

// PayloadInt64Required returns an error when key missing or not numeric.
func PayloadInt64Required(payload json.RawMessage, key string) (int64, error) {
	v, ok := PayloadInt64(payload, key)
	if !ok {
		return 0, fmt.Errorf("missing numeric payload key %q", key)
	}
	return v, nil
}
