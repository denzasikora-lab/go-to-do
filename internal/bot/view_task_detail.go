package bot

import (
	"fmt"
	"html"
	"strings"

	"github.com/denzasikora-lab/go-to-do/internal/bot/callbacks"
	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
	"github.com/denzasikora-lab/go-to-do/internal/platform/telegram/markup"
)

// FormatTaskDetailMessage renders the corporate task record card.
func FormatTaskDetailMessage(t *dtodo.Todo) string {
	var desc string
	if strings.TrimSpace(t.Description) == "" {
		desc = "<i>No narrative captured.</i>"
	} else {
		desc = html.EscapeString(t.Description)
	}
	return fmt.Sprintf(
		"<b>Work item #%d</b>\n<b>Title:</b> %s\n<b>Priority:</b> %s\n<b>State:</b> %s\n<b>Detail:</b>\n%s",
		t.ID,
		html.EscapeString(t.Title),
		html.EscapeString(string(t.Priority)),
		html.EscapeString(string(t.Status)),
		desc,
	)
}

// TaskDetailMarkup shows lifecycle operations with semantic colors.
func TaskDetailMarkup(t *dtodo.Todo) *markup.InlineKeyboard {
	rows := [][]markup.InlineButton{
		{
			{Text: "✅ Close task", CallbackData: markup.CallbackData(callbacks.Done(t.ID)), Style: markup.StyleSuccess},
			{Text: "🔁 Re-open", CallbackData: markup.CallbackData(callbacks.Reopen(t.ID)), Style: markup.StylePrimary},
		},
		{
			{Text: "✏️ Rename", CallbackData: markup.CallbackData(callbacks.EditTitle(t.ID)), Style: markup.StylePrimary},
			{Text: "📝 Narrative", CallbackData: markup.CallbackData(callbacks.EditDescription(t.ID)), Style: markup.StylePrimary},
		},
		{
			{Text: "🗑 Retire", CallbackData: markup.CallbackData(callbacks.DeletePrompt(t.ID)), Style: markup.StyleDanger},
			{Text: "↩ Backlog", CallbackData: markup.CallbackData(callbacks.MenuList), Style: markup.StylePrimary},
		},
	}
	return &markup.InlineKeyboard{InlineKeyboard: rows}
}
