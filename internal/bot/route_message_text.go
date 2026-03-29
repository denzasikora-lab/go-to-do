package bot

import (
	"context"
	"encoding/json"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

func (s *Service) routeNonCommandText(ctx context.Context, msg *tgbotapi.Message) {
	u, err := s.ResolveTelegramUser(ctx, msg.From)
	if err != nil {
		return
	}
	snap, err := s.LoadSession(ctx, u.ID)
	if err != nil {
		return
	}
	text := strings.TrimSpace(msg.Text)
	switch snap.State {
	case fsm.StateIdle:
		s.textIdleNudge(msg.Chat.ID)
	case fsm.StateAddTitle:
		s.textCaptureTitle(ctx, u.ID, msg.Chat.ID, snap, text)
	case fsm.StateAddDescription:
		s.textCaptureDescription(ctx, u.ID, msg.Chat.ID, snap, text)
	case fsm.StateEditTitle:
		s.textCommitTitleEdit(ctx, u, snap, msg.Chat.ID, text)
	case fsm.StateEditDescription:
		s.textCommitDescriptionEdit(ctx, u, snap, msg.Chat.ID, text)
	case fsm.StateAddPriority:
		_ = s.ReplyHTML(msg.Chat.ID, "<i>Use the colored priority matrix.</i>", PriorityPickerMarkup())
	default:
		snap.State = fsm.StateIdle
		snap.Payload = json.RawMessage(`{}`)
		_ = s.SaveSession(ctx, u.ID, snap)
		_ = s.ReplyHTML(msg.Chat.ID, "<i>Workflow cleared. Pick a lane on the console.</i>", MainMenuMarkup())
	}
}
