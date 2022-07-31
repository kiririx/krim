package api

import (
	"github.com/kiririx/krim/module/req"
	"github.com/kiririx/krim/service"
)

type UserApi struct {
}

// Register 用户注册
func (u *UserApi) Register(c *APICtx, param *req.Register) (any, error) {
	_, err := service.User.Register(param.Username, param.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Login 用户登陆
func (u *UserApi) Login(c *APICtx, param *req.Login) (any, error) {
	token, err := service.User.Login(param.Username, param.Password)
	if err != nil {
		return nil, err
	}
	return map[string]string{"token": token}, nil
}
