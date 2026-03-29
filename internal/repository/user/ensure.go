package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
)

// Ensure returns the user row, inserting when the Telegram account is new.
func Ensure(ctx context.Context, pool *pgxpool.Pool, telegramID int64, username *string) (*duser.User, error) {
	const q = `
INSERT INTO app_users (telegram_id, username)
VALUES ($1, $2)
ON CONFLICT (telegram_id) DO UPDATE
  SET username = COALESCE(EXCLUDED.username, app_users.username)
RETURNING id, telegram_id, username, created_at
`
	row := pool.QueryRow(ctx, q, telegramID, username)
	var u duser.User
	if err := row.Scan(&u.ID, &u.TelegramID, &u.Username, &u.CreatedAt); err != nil {
		return nil, fmt.Errorf("ensure user: %w", err)
	}
	return &u, nil
}
