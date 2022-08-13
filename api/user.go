package api

import (
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/logic"
	"github.com/kiririx/krim/module/req"
	"github.com/kiririx/krim/service"
)

type UserApi struct {
}

// Register 用户注册
func (u *UserApi) Register(c *ctx.Ctx, param *req.Register) (any, error) {
	_, err := logic.UserLogic.Register(param.Username, "未命名", param.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Login 用户登陆
func (u *UserApi) Login(c *ctx.Ctx, param *req.Login) (any, error) {
	token, err := service.UserService.Login(param.Username, param.Password)
	if err != nil {
		return nil, err
	}
	return map[string]string{"token": token}, nil
}
