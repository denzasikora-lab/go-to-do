package fsm

// State names the conversational node persisted for each operator session.
type State string

const (
	StateIdle State = "idle"

	StateAddTitle       State = "add_wait_title"
	StateAddDescription State = "add_wait_description"
	StateAddPriority    State = "add_wait_priority"

	StateEditTitle       State = "edit_wait_title"
	StateEditDescription State = "edit_wait_description"
)
