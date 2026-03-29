package bot

import (
	"context"
	"encoding/json"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/denzasikora-lab/go-to-do/internal/bot/callbacks"
	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	"github.com/denzasikora-lab/go-to-do/internal/fsm"
	todorepo "github.com/denzasikora-lab/go-to-do/internal/repository/todo"
)

func (s *Service) cbAddPriorityCommit(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User, data string) error {
	if err := s.AckCallback(q.ID, "", false); err != nil {
		return err
	}
	snap, err := s.LoadSession(ctx, u.ID)
	if err != nil {
		return err
	}
	if snap.State != fsm.StateAddPriority {
		return s.ReplyHTML(q.Message.Chat.ID, "<i>Priority matrix is inactive.</i>", MainMenuMarkup())
	}
	title, ok := fsm.PayloadString(snap.Payload, fsm.KeyDraftTitle)
	if !ok || title == "" {
		snap.State = fsm.StateIdle
		snap.Payload = json.RawMessage(`{}`)
		_ = s.SaveSession(ctx, u.ID, snap)
		return s.ReplyHTML(q.Message.Chat.ID, "<i>Draft title missing. Restart intake.</i>", MainMenuMarkup())
	}
	desc, _ := fsm.PayloadString(snap.Payload, fsm.KeyDraftDescription)
	var pri dtodo.Priority
	switch data {
	case callbacks.PriLow:
		pri = dtodo.PriorityLow
	case callbacks.PriNormal:
		pri = dtodo.PriorityNormal
	case callbacks.PriHigh:
		pri = dtodo.PriorityHigh
	default:
		pri = dtodo.PriorityNormal
	}
	t, err := todorepo.Create(ctx, s.Pool, u.ID, title, desc, pri)
	if err != nil {
		return err
	}
	snap.State = fsm.StateIdle
	snap.Payload = json.RawMessage(`{}`)
	if err := s.SaveSession(ctx, u.ID, snap); err != nil {
		return err
	}
	msg := "<b>Committed.</b>\nWork item <code>#" + strconv.FormatInt(t.ID, 10) + "</code> is now auditable in Postgres."
	return s.ReplyHTML(q.Message.Chat.ID, msg, MainMenuMarkup())
}
