package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
)

func (s *Service) cbTodoRetirePrompt(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User, id int64) error {
	if err := s.AckCallback(q.ID, "", false); err != nil {
		return err
	}
	_ = u
	return s.ReplyHTML(q.Message.Chat.ID, FormatDeletePrompt(id), DeleteConfirmMarkup(id))
}
