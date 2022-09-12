package repo

import (
	"github.com/kiririx/amasugi"
	"github.com/kiririx/krim/repo/model"
)

type _UserRepo struct {
	amasugi.AmiRepo[model.User]
}

func (receiver *_UserRepo) GetByUsername(username string) (*model.User, error) {
	user, err := receiver.Get("username = ?", username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (receiver *_UserRepo) GetByUsernameAndPassword(username, password string) (*model.User, error) {
	user, err := receiver.Get("username = ? and password", username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (receiver *_UserRepo) DeleteByUsername(username string) error {
	_, err := receiver.Delete("username = ?", username)
	return err
}
