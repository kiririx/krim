package logic

import (
	"errors"
	"github.com/kiririx/krim/conf"
	"github.com/kiririx/krim/constx"
	"github.com/kiririx/krim/model"
	"github.com/kiririx/krim/service"
	"github.com/kiririx/krutils/algox"
)

var UserLogic = &userLogic{}

type userLogic struct {
}

func (u *userLogic) ReRegister(username, nickname, password string) (*model.User, error) {
	err := conf.Sqlx.Where("username = ?", username).Delete(&model.User{}).Error
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
	tx := conf.Sqlx.Save(userModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, errors.New("注册失败")
	}
	return userModel, nil
}
