package logic

import (
	"errors"
	"github.com/kiririx/krim/constx"
	"github.com/kiririx/krim/repo/model"
	"github.com/kiririx/krim/service"
	"github.com/kiririx/krutils/algox"
)

var UserLogic = &userLogic{}

type userLogic struct {
}

func (u *userLogic) ReRegister(username, nickname, password string) (*model.User, error) {
	err := service.UserService.DeleteByUserName(username)
	if err != nil {
		return nil, err
	}
	return u.Register(username, nickname, password)
}

func (u *userLogic) Register(username, nickname, password string) (*model.User, error) {
	user, err := service.UserService.QueryByUsername(username)
	if err != nil && !errors.Is(err, constx.DBRecordNotFound) {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("用户已注册")
	}
	userModel := &model.User{
		Username: username,
		Password: algox.MD5(password),
		Nickname: nickname,
		Sex:      0,
	}
	err = service.UserService.Save(userModel)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}
