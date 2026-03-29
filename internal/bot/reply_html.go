package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/denzasikora-lab/go-to-do/internal/platform/telegram/markup"
)

// ReplyHTML delivers a message using HTML formatting and optional inline keyboard.
func (s *Service) ReplyHTML(chatID int64, html string, kb *markup.InlineKeyboard) error {
	msg := tgbotapi.NewMessage(chatID, html)
	msg.ParseMode = tgbotapi.ModeHTML
	if kb != nil {
		msg.ReplyMarkup = kb
	}
	_, err := s.TG.Send(msg)
	return err
}
