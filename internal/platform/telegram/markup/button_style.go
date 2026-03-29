package markup

// ButtonStyle matches Telegram Bot API keyboard button style values.
type ButtonStyle string

const (
	StylePrimary ButtonStyle = "primary"
	StyleSuccess ButtonStyle = "success"
	StyleDanger  ButtonStyle = "danger"
)
