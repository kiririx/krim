package model

import "time"

type BaseModel struct {
	Id        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
