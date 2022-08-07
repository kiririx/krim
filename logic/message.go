package logic

import (
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/service"
	"github.com/kiririx/krim/wsbus"
)

var MessageLogic = &messageLogic{}

type messageLogic struct {
}

// SendUserMessage 发送消息
//
// 先保存到数据库，再发送到ws客户端
func (*messageLogic) SendUserMessage(ctx *ctx.Ctx, sourceId, targetId uint64, message string) error {
	messageM, err := service.MessageService.SaveUserMessage(ctx, sourceId, targetId, message)
	if err != nil {
		return err
	}
	err = wsbus.SendUserMessage(ctx, sourceId, targetId, message, messageM.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
