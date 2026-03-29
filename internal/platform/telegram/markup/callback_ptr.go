package markup

// CallbackData builds a stable pointer used by the Telegram client.
func CallbackData(v string) *string {
	return &v
}
