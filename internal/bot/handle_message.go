package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) handleMessage(ctx context.Context, msg *tgbotapi.Message) {
	if msg.From == nil || msg.Chat == nil {
		return
	}
	if !msg.Chat.IsPrivate() {
		return
	}
	if msg.IsCommand() {
		switch msg.Command() {
		case "start":
			_ = s.HandleStart(ctx, msg)
		case "help":
			_ = s.HandleHelp(ctx, msg)
		default:
			_ = s.ReplyHTML(msg.Chat.ID, "<i>Unknown directive. Try /help</i>", MainMenuMarkup())
		}
		return
	}
	s.routeNonCommandText(ctx, msg)
}
