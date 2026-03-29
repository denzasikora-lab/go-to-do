package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleHelp surfaces compliance-oriented operator guidance.
func (s *Service) HandleHelp(ctx context.Context, msg *tgbotapi.Message) error {
	_ = ctx
	body := `<b>Governance playbook</b>
• <code>/start</code> — reopen the command surface
• <code>/help</code> — show this notice
• Inline keyboards encode approvals; favor taps over free text except where prompted
• <code>Bot API 9.4+</code> clients render semantic button colors`
	return s.ReplyHTML(msg.Chat.ID, body, MainMenuMarkup())
}
