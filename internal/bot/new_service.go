package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewService wires a bot runtime with its outbound API and transaction pool.
func NewService(api *tgbotapi.BotAPI, pool *pgxpool.Pool) *Service {
	return &Service{TG: api, Pool: pool}
}
