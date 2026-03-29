package user

import "time"

// User is the persisted operator identity mapped from Telegram.
type User struct {
	ID         int64
	TelegramID int64
	Username   *string
	CreatedAt  time.Time
}
