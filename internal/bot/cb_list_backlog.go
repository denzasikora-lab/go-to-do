package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	todorepo "github.com/denzasikora-lab/go-to-do/internal/repository/todo"
)

func (s *Service) cbListBacklog(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User, status *dtodo.Status) error {
	if err := s.AckCallback(q.ID, "", false); err != nil {
		return err
	}
	items, err := todorepo.ListByUser(ctx, s.Pool, u.ID, status, 24)
	if err != nil {
		return err
	}
	label := "Enterprise backlog — consolidated"

	if status != nil {
		switch *status {
		case dtodo.StatusOpen:
			label = "Execution lane — open commitments"
		case dtodo.StatusDone:
			label = "Closure ledger — delivered results"
		}
	}
	return s.ReplyHTML(q.Message.Chat.ID, FormatTaskListMessage(items, label), TaskListAnchors(items))
}
