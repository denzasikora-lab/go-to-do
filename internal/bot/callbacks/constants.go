package callbacks

// Callback namespaces stay short to respect the 64-byte Telegram limit.
const (
	MenuList   = "m:l"
	MenuAdd    = "m:a"
	MenuStats  = "m:s"
	MenuCancel = "m:c"

	AddSkipDesc = "a:s"

	PriLow    = "p:l"
	PriNormal = "p:n"
	PriHigh   = "p:h"

	FilterAll  = "f:a"
	FilterOpen = "f:o"
	FilterDone = "f:d"
)
