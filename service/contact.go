package service

import (
	"errors"
	"github.com/kiririx/krim/constx"
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/model"
)

var ContactService = &contactService{}

type contactService struct{}

func (*contactService) AddContactEvent(ctx *ctx.Ctx, sourceId, targetId uint64, event uint) error {
	if sourceId == targetId {
		return errors.New("不能添加自己为联系人")
	}
	return ctx.SqlCtl().Save(&model.Event{
		SourceId:  sourceId,
		TargetId:  targetId,
		EventType: event,
		Progress:  constx.EventProgress,
	}).Error
}
