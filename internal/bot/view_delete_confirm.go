package bot

import (
	"fmt"

	"github.com/denzasikora-lab/go-to-do/internal/bot/callbacks"
	"github.com/denzasikora-lab/go-to-do/internal/platform/telegram/markup"
)

// DeleteConfirmMarkup asks for a secondtap before destructive retirement.
func DeleteConfirmMarkup(todoID int64) *markup.InlineKeyboard {
	return &markup.InlineKeyboard{
		InlineKeyboard: [][]markup.InlineButton{
			{
				{Text: "⚠ Confirm retire", CallbackData: markup.CallbackData(callbacks.DeleteConfirm(todoID)), Style: markup.StyleDanger},
				{Text: "🛡 Abort", CallbackData: markup.CallbackData(callbacks.View(todoID)), Style: markup.StyleSuccess},
			},
		},
	}
}

// FormatDeletePrompt copy for governance trail.
func FormatDeletePrompt(todoID int64) string {
	return fmt.Sprintf("<b>Governance check</b>\nRetire work item <code>#%d</code>? This cannot be reversed from chat.", todoID)
}
