package callbacks

import "strconv"

// ParseSuffixInt parses an int64 after a fixed callback prefix.
func ParseSuffixInt(data, prefix string) (int64, bool) {
	if len(data) <= len(prefix) || data[:len(prefix)] != prefix {
		return 0, false
	}
	id, err := strconv.ParseInt(data[len(prefix):], 10, 64)
	if err != nil {
		return 0, false
	}
	return id, true
}
