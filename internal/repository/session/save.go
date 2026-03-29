package session

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

// Save upserts FSM state and payload for the app user.
func Save(ctx context.Context, pool *pgxpool.Pool, appUserID int64, snap *fsm.Snapshot) error {
	const q = `
INSERT INTO bot_sessions (app_user_id, state, payload)
VALUES ($1, $2, $3)
ON CONFLICT (app_user_id) DO UPDATE
  SET state = EXCLUDED.state,
      payload = EXCLUDED.payload,
      updated_at = now()
`
	_, err := pool.Exec(ctx, q, appUserID, string(snap.State), snap.Payload)
	if err != nil {
		return fmt.Errorf("session save: %w", err)
	}
	return nil
}
