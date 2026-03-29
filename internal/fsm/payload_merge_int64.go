package fsm

import "encoding/json"

// MergePayloadInt64 stores an int64 inside the JSON payload object.
func MergePayloadInt64(payload json.RawMessage, key string, value int64) (json.RawMessage, error) {
	m := payloadMap(payload)
	m[key] = value
	return json.Marshal(m)
}
