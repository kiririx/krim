package service

import (
	"github.com/kiririx/krim/repo"
	"github.com/kiririx/krim/repo/model"
	"github.com/kiririx/krutils/algox"
	"github.com/pkg/errors"
)

type userService struct {
}

func (u *userService) QueryByUsername(username string) (*model.User, error) {
	user, err := repo.UserRepo.GetByUsername(username)
	return user, err
}

func (u *userService) Save(user *model.User) error {
	_, err := repo.UserRepo.Insert(user)
	return err
}

func (u *userService) Login(username string, password string) (string, error) {
	user, err := repo.UserRepo.GetByUsernameAndPassword(username, algox.MD5(password))
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("用户不存在或密码错误")
	}
	return BuildToken(user.Id, user.Username)
}

func (u *userService) QueryById(userId uint64) (*model.User, error) {
	user, err := repo.UserRepo.GetById(userId)
	return user, err
}
func (u *userService) DeleteByUserName(username string) error {
	return repo.UserRepo.DeleteByUsername(username)
}
