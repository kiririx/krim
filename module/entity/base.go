package entity

import "time"

type Page struct {
	Cur   int   `json:"cur"`   // 当前页
	Size  int   `json:"size"`  // 每页大小
	Total int64 `json:"total"` // 总数
	Data  any   `json:"data"`  // 数据
}

type BaseField struct {
	Id        uint      `gorm:"column:id;type:integer;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"createTime"`
	UpdatedAt time.Time `json:"updateTime"`
}

func (b *BaseField) IsEmpty() bool {
	if b.Id <= 0 {
		return true
	}
	return false
}
