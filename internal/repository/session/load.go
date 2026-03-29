package session

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

// Load reads the FSM snapshot for an app user.
func Load(ctx context.Context, pool *pgxpool.Pool, appUserID int64) (*fsm.Snapshot, error) {
	const q = `SELECT state, payload FROM bot_sessions WHERE app_user_id = $1`
	var st string
	var raw []byte
	err := pool.QueryRow(ctx, q, appUserID).Scan(&st, &raw)
	if errors.Is(err, pgx.ErrNoRows) {
		return &fsm.Snapshot{State: fsm.StateIdle, Payload: json.RawMessage(`{}`)}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("session load: %w", err)
	}
	return &fsm.Snapshot{State: fsm.State(st), Payload: json.RawMessage(raw)}, nil
}
