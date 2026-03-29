package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Service hosts Telegram ingress dependencies for the corporate task surface.
type Service struct {
	TG   *tgbotapi.BotAPI
	Pool *pgxpool.Pool
}
