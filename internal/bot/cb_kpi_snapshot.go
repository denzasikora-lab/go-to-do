package bot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	todorepo "github.com/denzasikora-lab/go-to-do/internal/repository/todo"
)

func (s *Service) cbKPISnapshot(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User) error {
	if err := s.AckCallback(q.ID, "", false); err != nil {
		return err
	}
	openN, doneN, err := todorepo.CountOpenDone(ctx, s.Pool, u.ID)
	if err != nil {
		return err
	}
	html := fmt.Sprintf(
		"<b>KPI snapshot</b>\n• Open commitments: <code>%d</code>\n• Delivered (unarchived): <code>%d</code>",
		openN, doneN,
	)
	return s.ReplyHTML(q.Message.Chat.ID, html, MainMenuMarkup())
}
