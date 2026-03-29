package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// AckCallback closes the Telegram client wait state for an inline tap.
func (s *Service) AckCallback(callbackID string, text string, alert bool) error {
	cfg := tgbotapi.NewCallback(callbackID, text)
	cfg.ShowAlert = alert
	_, err := s.TG.Request(cfg)
	return err
}
