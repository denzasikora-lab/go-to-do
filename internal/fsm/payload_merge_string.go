package fsm

import "encoding/json"

const (
	KeyDraftTitle       = "draft_title"
	KeyDraftDescription = "draft_description"
)

// MergePayloadString inserts or replaces a string key inside a JSON object.
func MergePayloadString(payload json.RawMessage, key, value string) (json.RawMessage, error) {
	m := payloadMap(payload)
	m[key] = value
	return json.Marshal(m)
}
