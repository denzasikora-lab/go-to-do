package bot

import (
	"github.com/denzasikora-lab/go-to-do/internal/bot/callbacks"
	"github.com/denzasikora-lab/go-to-do/internal/platform/telegram/markup"
)

// MainMenuMarkup builds the landing keyboard with enterprise color cues.
func MainMenuMarkup() *markup.InlineKeyboard {
	return &markup.InlineKeyboard{
		InlineKeyboard: [][]markup.InlineButton{
			{
				{Text: "📋 Open backlog", CallbackData: markup.CallbackData(callbacks.MenuList), Style: markup.StylePrimary},
				{Text: "➕ New task", CallbackData: markup.CallbackData(callbacks.MenuAdd), Style: markup.StyleSuccess},
				{Text: "📊 KPI snapshot", CallbackData: markup.CallbackData(callbacks.MenuStats), Style: markup.StylePrimary},
			},
			{
				{Text: "🔓 Active", CallbackData: markup.CallbackData(callbacks.FilterOpen), Style: markup.StyleSuccess},
				{Text: "✅ Closed", CallbackData: markup.CallbackData(callbacks.FilterDone), Style: markup.StylePrimary},
				{Text: "🗂️ All streams", CallbackData: markup.CallbackData(callbacks.FilterAll), Style: markup.StylePrimary},
			},
			{
				{Text: "⛔ Abort workflow", CallbackData: markup.CallbackData(callbacks.MenuCancel), Style: markup.StyleDanger},
			},
		},
	}
}
