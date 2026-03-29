package bot

import (
	"context"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

func (s *Service) cbEditDescriptionBegin(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User, todoID int64) error {
	if err := s.AckCallback(q.ID, "", false); err != nil {
		return err
	}
	payload, err := fsm.MergePayloadInt64(json.RawMessage(`{}`), fsm.KeyEditTodoID, todoID)
	if err != nil {
		return err
	}
	snap := &fsm.Snapshot{State: fsm.StateEditDescription, Payload: payload}
	if err := s.SaveSession(ctx, u.ID, snap); err != nil {
		return err
	}
	return s.ReplyHTML(q.Message.Chat.ID, "<b>Narrative revision</b>\nSend the updated storyline.", nil)
}
