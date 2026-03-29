package markup

// InlineButton is a Telegram inline keyboard cell with optional color style (Bot API 9.4+).
type InlineButton struct {
	Text         string      `json:"text"`
	CallbackData *string     `json:"callback_data,omitempty"`
	Style        ButtonStyle `json:"style,omitempty"`
}

// InlineKeyboard is serialized as reply_markup for sendMessage.
type InlineKeyboard struct {
	InlineKeyboard [][]InlineButton `json:"inline_keyboard"`
}
