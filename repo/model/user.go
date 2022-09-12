package model

import "github.com/kiririx/amasugi"

type User struct {
	BaseModel
	Username string `ami:"username"`
	Password string `ami:"password"`
	Nickname string `ami:"nickname"`
	Sex      uint   `ami:"sex"`
	amasugi.AmiRepo[User]
}

func (User) TableName() string {
	return "user"
}

type UserLogonLog struct {
	BaseModel
	UserId uint64
}
