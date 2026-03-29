package bot

import (
	"context"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

func (s *Service) cbSessionReset(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User) error {
	if err := s.AckCallback(q.ID, "", false); err != nil {
		return err
	}
	snap := &fsm.Snapshot{State: fsm.StateIdle, Payload: json.RawMessage(`{}`)}
	if err := s.SaveSession(ctx, u.ID, snap); err != nil {
		return err
	}
	return s.ReplyHTML(q.Message.Chat.ID, "<b>Console reset.</b>\nAll conversational state cleared from <code>bot_sessions</code>.", MainMenuMarkup())
}
