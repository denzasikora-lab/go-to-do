package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Load reads .env when present and builds Config from the environment.
func Load() (*Config, error) {
	_ = godotenv.Load()
	cfg := &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		PostgresDSN:      os.Getenv("POSTGRES_DSN"),
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
