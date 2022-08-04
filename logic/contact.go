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

func (*contactLogic) AddContactEvent(ctx *ctx.Ctx, sourceId, targetId uint64, event uint) error {
	err := service.ContactService.AddContactEvent(ctx, sourceId, targetId, event)
	if err != nil {
		return err
	}
	user, err := service.User.QueryById(targetId)
	if err != nil {
		return err
	}
	err = wsbus.SendUserMessage(ctx, sourceId, targetId, fmt.Sprintf("%v请求添加您为好友", user.Nickname))
	if err != nil {
		return err
	}
	return nil
}
