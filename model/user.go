package model

type User struct {
	BaseModel
	UserName string
	Password string
	NickName string
	Sex      uint
}

type UserLogonLog struct {
	BaseModel
	UserId uint64
}
