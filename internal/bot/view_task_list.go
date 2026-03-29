package bot

import (
	"fmt"
	"html"

	"github.com/denzasikora-lab/go-to-do/internal/bot/callbacks"
	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
	"github.com/denzasikora-lab/go-to-do/internal/platform/telegram/markup"
)

// FormatTaskListMessage builds an HTML summary block for a backlog slice.
func FormatTaskListMessage(items []dtodo.Todo, label string) string {
	if len(items) == 0 {
		return fmt.Sprintf("<b>%s</b>\n<i>No work items match this filter.</i>", html.EscapeString(label))
	}
	var b string
	b += fmt.Sprintf("<b>%s</b>\n", html.EscapeString(label))
	for _, t := range items {
		status := string(t.Status)
		if t.Status == dtodo.StatusDone {
			status = "✅ " + status
		} else {
			status = "🔓 " + status
		}
		b += fmt.Sprintf("• <code>#%d</code> %s — %s · %s\n",
			t.ID, html.EscapeString(t.Title), html.EscapeString(string(t.Priority)), status)
	}
	return b
}

// TaskListAnchors attaches one primary action row per visible task (open detail).
func TaskListAnchors(items []dtodo.Todo) *markup.InlineKeyboard {
	rows := make([][]markup.InlineButton, 0, len(items)+1)
	for _, t := range items {
		st := markup.StylePrimary
		if t.Status == dtodo.StatusDone {
			st = markup.StyleSuccess
		}
		label := fmt.Sprintf("Inspect #%d", t.ID)
		cb := callbacks.View(t.ID)
		rows = append(rows, []markup.InlineButton{
			{Text: label, CallbackData: markup.CallbackData(cb), Style: st},
		})
	}
	rows = append(rows, []markup.InlineButton{
		{Text: "↩ Main desk", CallbackData: markup.CallbackData(callbacks.MenuCancel), Style: markup.StylePrimary},
	})
	return &markup.InlineKeyboard{InlineKeyboard: rows}
}
