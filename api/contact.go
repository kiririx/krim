package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kiririx/krim/constx"
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/logic"
	"github.com/kiririx/krim/mapping"
	"github.com/kiririx/krim/module/req"
	"github.com/kiririx/krim/service"
)

var ContactAPI = &_ContactAPI{}

type _ContactAPI struct {
}

// AddContact 添加联系人
func (*_ContactAPI) AddContact(ctx *gin.Context, param *req.AddContact) (any, error) {
	return nil, nil
}

func (*_ContactAPI) GetContact(ctx *ctx.Ctx, param *req.GetContact) (any, error) {
	user, err := service.User.QueryByUsername(param.Username)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"user_id":  user.Id,
		"username": user.Username,
		"nickname": user.Nickname,
		"sex":      mapping.SexGet(user.Sex),
	}, nil
}

// AddContactEvent add contact event
func (*_ContactAPI) AddContactEvent(ctx *ctx.Ctx, param *req.AddContactEvent) (any, error) {
	err := logic.ContactLogic.AddContactEvent(ctx, ctx.UserId, param.TargetId, constx.EventAddContact)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
