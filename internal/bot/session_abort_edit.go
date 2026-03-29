package bot

import (
	"context"
	"encoding/json"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

func (s *Service) abortEdit(ctx context.Context, appUserID, chatID int64, snap *fsm.Snapshot) {
	snap.State = fsm.StateIdle
	snap.Payload = json.RawMessage(`{}`)
	_ = s.SaveSession(ctx, appUserID, snap)
	_ = s.ReplyHTML(chatID, "<i>Edit workflow aborted.</i>", MainMenuMarkup())
}
