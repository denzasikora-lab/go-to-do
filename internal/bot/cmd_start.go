package bot

import (
	"context"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

// HandleStart resets the FSM and presents the executive dashboard.
func (s *Service) HandleStart(ctx context.Context, msg *tgbotapi.Message) error {
	u, err := s.ResolveTelegramUser(ctx, msg.From)
	if err != nil {
		return err
	}
	snap := &fsm.Snapshot{State: fsm.StateIdle, Payload: json.RawMessage(`{}`)}
	if err := s.SaveSession(ctx, u.ID, snap); err != nil {
		return err
	}
	welcome := "<b>CorpTodo Command Center</b>\nOperational excellence begins with disciplined backlog hygiene.\nSelect a workflow lane below."
	return s.ReplyHTML(msg.Chat.ID, welcome, MainMenuMarkup())
}
