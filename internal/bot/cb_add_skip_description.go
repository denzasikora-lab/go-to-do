package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

func (s *Service) cbAddSkipDescription(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User) error {
	if err := s.AckCallback(q.ID, "", false); err != nil {
		return err
	}
	snap, err := s.LoadSession(ctx, u.ID)
	if err != nil {
		return err
	}
	if snap.State != fsm.StateAddDescription {
		return s.ReplyHTML(q.Message.Chat.ID, "<i>Skip is valid only during narrative capture.</i>", MainMenuMarkup())
	}
	payload, err := fsm.MergePayloadString(snap.Payload, fsm.KeyDraftDescription, "")
	if err != nil {
		return err
	}
	snap.State = fsm.StateAddPriority
	snap.Payload = payload
	if err := s.SaveSession(ctx, u.ID, snap); err != nil {
		return err
	}
	return s.ReplyHTML(q.Message.Chat.ID, "<b>Priority classification</b>\nSelect a governance tier.", PriorityPickerMarkup())
}
