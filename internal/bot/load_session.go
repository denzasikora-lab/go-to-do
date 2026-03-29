package bot

import (
	"context"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
	sessionrepo "github.com/denzasikora-lab/go-to-do/internal/repository/session"
)

// LoadSession returns persisted FSM state for the internal user id.
func (s *Service) LoadSession(ctx context.Context, appUserID int64) (*fsm.Snapshot, error) {
	return sessionrepo.Load(ctx, s.Pool, appUserID)
}
