package fsm

import "encoding/json"

// Snapshot is the serializable FSM view stored inside bot_sessions.payload.
type Snapshot struct {
	State   State           `json:"state"`
	Payload json.RawMessage `json:"-"`
}
