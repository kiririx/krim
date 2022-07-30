package model

type Group struct {
	BaseModel
	Name string
	Desc string
}

type GroupUser struct {
	BaseModel
	GroupId uint64
	UserId  uint64
	// Role 0: master 1:manager 2:normal member
	Role uint
}
