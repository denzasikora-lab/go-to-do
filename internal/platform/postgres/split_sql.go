package postgres

import "strings"

// splitSQLStatements splits a migration script on semicolons. It is suitable for DDL that does not
// embed semicolons inside string literals or function bodies.
func splitSQLStatements(script string) []string {
	chunks := strings.Split(script, ";")
	out := make([]string, 0, len(chunks))
	for _, c := range chunks {
		s := strings.TrimSpace(c)
		if s == "" {
			continue
		}
		out = append(out, s+";")
	}
	return out
}
