package model

type Event struct {
	BaseModel
	EventType uint // EventType 0: add contact 1:add group
	SourceId  uint64
	TargetId  uint64
	Progress  uint // Progress 0: not processed 1: processed
}

func (*Event) TableName() string {
	return "event"
}
