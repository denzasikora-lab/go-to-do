package fsm

import "encoding/json"

// PayloadString extracts a string value from JSON object bytes.
func PayloadString(payload json.RawMessage, key string) (string, bool) {
	m := payloadMap(payload)
	v, ok := m[key]
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}

func payloadMap(payload json.RawMessage) map[string]interface{} {
	if len(payload) == 0 {
		return map[string]interface{}{}
	}
	var m map[string]interface{}
	if err := json.Unmarshal(payload, &m); err != nil || m == nil {
		return map[string]interface{}{}
	}
	return m
}
