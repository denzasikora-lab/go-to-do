// Command bot is the process entrypoint: load configuration, run embedded SQL migrations,
// connect to Postgres, authorize against Telegram, and poll for updates until a shutdown signal.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/denzasikora-lab/go-to-do/internal/bot"
	"github.com/denzasikora-lab/go-to-do/internal/config"
	"github.com/denzasikora-lab/go-to-do/internal/platform/postgres"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	pool, err := postgres.NewPool(ctx, cfg.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := postgres.ApplyMigrations(ctx, pool); err != nil {
		log.Fatal(err)
	}
	if err := postgres.Ping(ctx, pool); err != nil {
		log.Fatal(err)
	}

	api, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("authorized on @%s", api.Self.UserName)
	svc := bot.NewService(api, pool)

	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60 // long polling interval in seconds (Telegram getUpdates)
	updates := api.GetUpdatesChan(ucfg)

	for {
		select {
		case <-ctx.Done():
			log.Printf("shutdown: %v", ctx.Err())
			return
		case u, ok := <-updates:
			if !ok {
				return
			}
			svc.HandleUpdate(context.Background(), u)
		}
	}
}
