package api

import (
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/logic"
	"github.com/kiririx/krim/module/req"
)

// SendUserMessage 发送用户消息接口
func SendUserMessage(c *ctx.Ctx, param *req.SendUserMessageReq) (any, error) {
	err := logic.MessageLogic.SendUserMessage(c, c.UserId, param.TargetId, param.Message)
	if err != nil {
		return nil, err
	}
	return ``, nil
}

// SendGroupMessage todo
func SendGroupMessage(c *ctx.Ctx, param *req.SendGroupMessageReq) (any, error) {
	return nil, nil
}
