package bot

import (
	"context"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
	sessionrepo "github.com/denzasikora-lab/go-to-do/internal/repository/session"
)

// SaveSession commits FSM transitions to Postgres.
func (s *Service) SaveSession(ctx context.Context, appUserID int64, snap *fsm.Snapshot) error {
	return sessionrepo.Save(ctx, s.Pool, appUserID, snap)
}
