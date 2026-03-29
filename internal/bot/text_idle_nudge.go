package bot

// textIdleNudge reminds operators that structured navigation is mandatory.
func (s *Service) textIdleNudge(chatID int64) {
	_ = s.ReplyHTML(chatID, "<i>Please use the colored console actions. Free-form text is ignored in this lane.</i>", MainMenuMarkup())
}
