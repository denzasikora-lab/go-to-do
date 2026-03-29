package bot

import (
	"context"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

func (s *Service) textCaptureDescription(ctx context.Context, appUserID, chatID int64, snap *fsm.Snapshot, text string) {
	payload, err := fsm.MergePayloadString(snap.Payload, fsm.KeyDraftDescription, text)
	if err != nil {
		return
	}
	snap.State = fsm.StateAddPriority
	snap.Payload = payload
	if err := s.SaveSession(ctx, appUserID, snap); err != nil {
		return
	}
	_ = s.ReplyHTML(chatID, "<b>Priority classification</b>\nSelect a governance tier.", PriorityPickerMarkup())
}
