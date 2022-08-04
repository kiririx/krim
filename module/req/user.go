package req

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddContact struct {
	TargetId  uint64 `json:"target_id"`
	EventType uint   `json:"event_type"`
}

type GetContact struct {
	Username string `uri:"username" binding:"required"`
}

type AddContactEvent struct {
	TargetId uint64 `json:"target_id"`
	// EventType uint   `json:"event_type"`
}

type SendUserMessageReq struct {
	TargetId uint64
	Message  string
}

type SendGroupMessageReq struct {
	GroupId uint64
	Message string
}

type EventReq struct {
}
