package model

import "time"

type BaseModel struct {
	Id        uint64    `ami:"id"`
	CreatedAt time.Time `ami:"created_at"`
	UpdatedAt time.Time `ami:"updated_at"`
}
