package postgres

import (
	"context"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/denzasikora-lab/go-to-do/internal/platform/postgres/migrations"
)

// ApplyMigrations runs embedded SQL files in lexical order against pool.
func ApplyMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	names, err := fs.Glob(migrations.FS, "*.sql")
	if err != nil {
		return fmt.Errorf("glob migrations: %w", err)
	}
	sort.Strings(names)
	for _, name := range names {
		body, rerr := migrations.FS.ReadFile(name)
		if rerr != nil {
			return fmt.Errorf("read %s: %w", name, rerr)
		}
		sqlText := strings.TrimSpace(string(body))
		if sqlText == "" {
			continue
		}
		for _, stmt := range splitSQLStatements(sqlText) {
			if _, err = pool.Exec(ctx, stmt); err != nil {
				return fmt.Errorf("exec %s: %w", name, err)
			}
		}
	}
	return nil
}
