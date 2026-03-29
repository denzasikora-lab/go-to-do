package config

// Config aggregates process configuration loaded from the environment.
type Config struct {
	TelegramBotToken string
	PostgresDSN      string
}
