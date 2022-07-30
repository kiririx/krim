package model

type User struct {
	BaseModel
	Username string
	Password string
	Nickname string
	Sex      uint
}

func (*User) TableName() string {
	return "user"
}

type UserLogonLog struct {
	BaseModel
	UserId uint64
}
