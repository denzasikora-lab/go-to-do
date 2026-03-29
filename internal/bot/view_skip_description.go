package bot

import (
	"github.com/denzasikora-lab/go-to-do/internal/bot/callbacks"
	"github.com/denzasikora-lab/go-to-do/internal/platform/telegram/markup"
)

// SkipDescriptionMarkup offers a bypass before the priority commitment gate.
func SkipDescriptionMarkup() *markup.InlineKeyboard {
	return &markup.InlineKeyboard{
		InlineKeyboard: [][]markup.InlineButton{
			{
				{Text: "Skip narrative", CallbackData: markup.CallbackData(callbacks.AddSkipDesc), Style: markup.StylePrimary},
			},
		},
	}
}
