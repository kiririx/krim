package logic

import (
	"fmt"
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/service"
	"github.com/kiririx/krim/wsbus"
)

var ContactLogic = &contactLogic{}

type contactLogic struct {
}

func (*contactLogic) AddContactEvent(ctx *ctx.Ctx, sourceId, targetId uint64, eventType uint) error {
	err := service.ContactService.AddContactEvent(ctx, sourceId, targetId, eventType)
	if err != nil {
		return err
	}
	user, err := service.UserService.QueryById(targetId)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("%v请求添加您为好友", user.Nickname)
	eventM, err := service.EventService.GetBySourceIdAndTargetIdAndType(ctx, sourceId, targetId, eventType)
	if err != nil {
		return err
	}
	// 将sourceId换成系统通知user
	err = wsbus.SendUserMessage(ctx, sourceId, targetId, msg, eventM.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
