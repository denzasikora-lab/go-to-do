package bot

import (
	"context"
	"encoding/json"
	"errors"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	"github.com/denzasikora-lab/go-to-do/internal/fsm"
	todorepo "github.com/denzasikora-lab/go-to-do/internal/repository/todo"
)

func (s *Service) textCommitDescriptionEdit(ctx context.Context, u *duser.User, snap *fsm.Snapshot, chatID int64, text string) {
	id, err := fsm.PayloadInt64Required(snap.Payload, fsm.KeyEditTodoID)
	if err != nil {
		s.abortEdit(ctx, u.ID, chatID, snap)
		return
	}
	err = todorepo.UpdateDescription(ctx, s.Pool, u.ID, id, text)
	if errors.Is(err, dtodo.ErrNotFound) {
		_ = s.ReplyHTML(chatID, "<i>Work item not found.</i>", MainMenuMarkup())
	} else if err != nil {
		return
	}
	snap.State = fsm.StateIdle
	snap.Payload = json.RawMessage(`{}`)
	_ = s.SaveSession(ctx, u.ID, snap)
	_ = s.ReplyHTML(chatID, "<b>Narrative updated.</b>", MainMenuMarkup())
}
