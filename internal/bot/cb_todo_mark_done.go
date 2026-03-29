package bot

import (
	"context"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	todorepo "github.com/denzasikora-lab/go-to-do/internal/repository/todo"
)

func (s *Service) cbTodoMarkDone(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User, id int64) error {
	if err := s.AckCallback(q.ID, "Marked delivered", false); err != nil {
		return err
	}
	err := todorepo.SetStatus(ctx, s.Pool, u.ID, id, dtodo.StatusDone)
	if errors.Is(err, dtodo.ErrNotFound) {
		return s.ReplyHTML(q.Message.Chat.ID, "<i>Record not found.</i>", MainMenuMarkup())
	}
	if err != nil {
		return err
	}
	t, err := todorepo.GetByID(ctx, s.Pool, u.ID, id)
	if err != nil {
		return err
	}
	return s.ReplyHTML(q.Message.Chat.ID, FormatTaskDetailMessage(t), TaskDetailMarkup(t))
}
