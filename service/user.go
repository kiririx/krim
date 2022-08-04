package service

import (
	"github.com/kiririx/krim/conf"
	"github.com/kiririx/krim/constx"
	"github.com/kiririx/krim/model"
	"github.com/kiririx/krutils/algox"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type userService struct {
}

func (u *userService) QueryByUsername(username string) (*model.User, error) {
	var user model.User
	tx := conf.Sqlx.Where("username = ?", username).Take(&user)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constx.DBRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userService) Save(user *model.User) bool {
	result := conf.Sqlx.Create(user)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return false
	}
	return true
}

// func (u *userService) Query(e *model.User) any {
// 	first := conf.Sqlx.First(e)
// 	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
// 		return nil
// 	}
// 	if first.Error == gorm.ErrRecordNotFound {
// 		return nil
// 	}
// 	return e
// }

func (u *userService) Register(username string, password string) (*model.User, error) {
	user, err := User.QueryByUsername(username)
	if err != nil && !errors.Is(err, constx.DBRecordNotFound) {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("用户已注册")
	}
	userModel := &model.User{
		Username: username,
		Password: algox.MD5(password),
		Nickname: "未命名",
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

func (u *userService) Login(username string, password string) (string, error) {
	var user model.User
	if err := conf.Sqlx.Where("username = ? and password = ?", username, algox.MD5(password)).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("用户不存在或密码错误")
		}
		return "", err
	}
	return BuildToken(user.Id, user.Username)
}

func (u *userService) QueryById(userId uint64) (*model.User, error) {
	var user model.User
	tx := conf.Sqlx.Where("id = ?", userId).Take(&user)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
