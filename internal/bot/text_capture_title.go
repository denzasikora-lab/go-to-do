package bot

import (
	"context"
	"strings"

	"github.com/denzasikora-lab/go-to-do/internal/fsm"
)

func (s *Service) textCaptureTitle(ctx context.Context, appUserID, chatID int64, snap *fsm.Snapshot, text string) {
	if len(strings.TrimSpace(text)) < 2 {
		_ = s.ReplyHTML(chatID, "<b>Validation</b>\nTitle must contain at least two visible characters.", nil)
		return
	}
	payload, err := fsm.MergePayloadString(snap.Payload, fsm.KeyDraftTitle, text)
	if err != nil {
		return
	}
	snap.State = fsm.StateAddDescription
	snap.Payload = payload
	if err := s.SaveSession(ctx, appUserID, snap); err != nil {
		return
	}
	_ = s.ReplyHTML(chatID, "<b>Narrative</b>\nProvide supporting detail, or tap <i>Skip narrative</i> below.", SkipDescriptionMarkup())
}
