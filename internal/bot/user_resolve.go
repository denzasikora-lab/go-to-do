package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
	userrepo "github.com/denzasikora-lab/go-to-do/internal/repository/user"
)

// ResolveTelegramUser maps the Telegram principal to an internal app_users row.
func (s *Service) ResolveTelegramUser(ctx context.Context, tgUser *tgbotapi.User) (*duser.User, error) {
	var username *string
	if tgUser.UserName != "" {
		u := tgUser.UserName
		username = &u
	}
	return userrepo.Ensure(ctx, s.Pool, tgUser.ID, username)
}
