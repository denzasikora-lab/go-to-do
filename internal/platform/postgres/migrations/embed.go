package migrations

import "embed"

// FS holds versioned SQL for deterministic startup migrations.
//
//go:embed *.sql
var FS embed.FS
