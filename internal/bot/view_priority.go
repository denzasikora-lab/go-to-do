package bot

import (
	"github.com/denzasikora-lab/go-to-do/internal/bot/callbacks"
	"github.com/denzasikora-lab/go-to-do/internal/platform/telegram/markup"
)

// PriorityPickerMarkup renders colored priority commitment buttons.
func PriorityPickerMarkup() *markup.InlineKeyboard {
	return &markup.InlineKeyboard{
		InlineKeyboard: [][]markup.InlineButton{
			{
				{Text: "Low touch", CallbackData: markup.CallbackData(callbacks.PriLow), Style: markup.StyleSuccess},
				{Text: "Standard", CallbackData: markup.CallbackData(callbacks.PriNormal), Style: markup.StylePrimary},
				{Text: "Critical", CallbackData: markup.CallbackData(callbacks.PriHigh), Style: markup.StyleDanger},
			},
		},
	}
}
