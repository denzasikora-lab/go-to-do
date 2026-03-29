package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleUpdate is the top-level demultiplexer for Telegram webhook/poll traffic.
func (s *Service) HandleUpdate(ctx context.Context, u tgbotapi.Update) {
	switch {
	case u.Message != nil:
		s.handleMessage(ctx, u.Message)
	case u.CallbackQuery != nil:
		s.handleCallback(ctx, u.CallbackQuery)
	}
}
