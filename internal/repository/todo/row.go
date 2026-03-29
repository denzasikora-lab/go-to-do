package todo

// pgxRow matches pgx row scanners used by repository helpers.
type pgxRow interface {
	Scan(dest ...any) error
}
