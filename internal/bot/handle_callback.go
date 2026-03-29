package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/denzasikora-lab/go-to-do/internal/bot/callbacks"
	dtodo "github.com/denzasikora-lab/go-to-do/internal/domain/todo"
	duser "github.com/denzasikora-lab/go-to-do/internal/domain/user"
)

func (s *Service) handleCallback(ctx context.Context, q *tgbotapi.CallbackQuery) {
	if q.From == nil || q.Message == nil {
		return
	}
	if !q.Message.Chat.IsPrivate() {
		return
	}
	u, err := s.ResolveTelegramUser(ctx, q.From)
	if err != nil {
		log.Printf("resolve user: %v", err)
		return
	}
	data := q.Data

	switch data {
	case callbacks.MenuList, callbacks.FilterAll:
		err = s.cbListBacklog(ctx, q, u, nil)
	case callbacks.FilterOpen:
		st := dtodo.StatusOpen
		err = s.cbListBacklog(ctx, q, u, &st)
	case callbacks.FilterDone:
		st := dtodo.StatusDone
		err = s.cbListBacklog(ctx, q, u, &st)
	case callbacks.MenuAdd:
		err = s.cbAddWorkflowStart(ctx, q, u)
	case callbacks.MenuStats:
		err = s.cbKPISnapshot(ctx, q, u)
	case callbacks.MenuCancel:
		err = s.cbSessionReset(ctx, q, u)
	case callbacks.AddSkipDesc:
		err = s.cbAddSkipDescription(ctx, q, u)
	case callbacks.PriLow, callbacks.PriNormal, callbacks.PriHigh:
		err = s.cbAddPriorityCommit(ctx, q, u, data)
	default:
		err = s.dispatchPrefixedCallbacks(ctx, q, u, data)
	}
	if err != nil {
		log.Printf("callback %q: %v", data, err)
	}
}

func (s *Service) dispatchPrefixedCallbacks(ctx context.Context, q *tgbotapi.CallbackQuery, u *duser.User, data string) error {
	if id, ok := callbacks.ParseSuffixInt(data, callbacks.PrefView); ok {
		return s.cbTodoOpenDetail(ctx, q, u, id)
	}
	if id, ok := callbacks.ParseSuffixInt(data, callbacks.PrefDone); ok {
		return s.cbTodoMarkDone(ctx, q, u, id)
	}
	if id, ok := callbacks.ParseSuffixInt(data, callbacks.PrefReopen); ok {
		return s.cbTodoMarkOpen(ctx, q, u, id)
	}
	if id, ok := callbacks.ParseSuffixInt(data, callbacks.PrefDelAsk); ok {
		return s.cbTodoRetirePrompt(ctx, q, u, id)
	}
	if id, ok := callbacks.ParseSuffixInt(data, callbacks.PrefDelYes); ok {
		return s.cbTodoRetireConfirm(ctx, q, u, id)
	}
	if id, ok := callbacks.ParseSuffixInt(data, callbacks.PrefEditTitle); ok {
		return s.cbEditTitleBegin(ctx, q, u, id)
	}
	if id, ok := callbacks.ParseSuffixInt(data, callbacks.PrefEditDesc); ok {
		return s.cbEditDescriptionBegin(ctx, q, u, id)
	}
	_ = s.AckCallback(q.ID, "Unrecognized control", true)
	return nil
}
