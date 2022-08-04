package service

import (
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/model"
)

var MessageService = &messageService{}

type messageService struct{}

// SaveUserMessage 保存用户消息(unread)
func (*messageService) SaveUserMessage(ctx *ctx.Ctx, sourceId, targetId uint64, msg string) (*model.Message, error) {
	messageModel := model.Message{
		Msg:      msg,
		SourceId: sourceId,
		TargetId: targetId,
	}
	err := ctx.SqlCtl().Save(&messageModel).Error
	if err != nil {
		return nil, err
	}
	return &messageModel, nil
}
