package config

import "fmt"

// Validate enforces required production fields.
func (c *Config) Validate() error {
	if c.TelegramBotToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN is required")
	}
	if c.PostgresDSN == "" {
		return fmt.Errorf("POSTGRES_DSN is required")
	}
	return nil
}
