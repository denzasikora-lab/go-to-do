package callbacks

import "fmt"

// View formats todo detail open action.
func View(todoID int64) string {
	return fmt.Sprintf("t:v:%d", todoID)
}

// Done marks todo completed.
func Done(todoID int64) string {
	return fmt.Sprintf("t:d:%d", todoID)
}

// Reopen moves todo back to open.
func Reopen(todoID int64) string {
	return fmt.Sprintf("t:r:%d", todoID)
}

// DeletePrompt first step.
func DeletePrompt(todoID int64) string {
	return fmt.Sprintf("t:x:%d", todoID)
}

// DeleteConfirm irreversible.
func DeleteConfirm(todoID int64) string {
	return fmt.Sprintf("t:y:%d", todoID)
}

// EditTitle starts title edit FSM.
func EditTitle(todoID int64) string {
	return fmt.Sprintf("t:e:t:%d", todoID)
}

// EditDescription starts description edit FSM.
func EditDescription(todoID int64) string {
	return fmt.Sprintf("t:e:d:%d", todoID)
}
