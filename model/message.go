package model

type Message struct {
	BaseModel
	Msg      string
	SourceId uint64
	TargetId uint64
	GroupId  uint64
}
