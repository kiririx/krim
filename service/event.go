package service

import (
	"github.com/kiririx/krim/ctx"
	"github.com/kiririx/krim/model"
)

var EventService = &_Event{}

type _Event struct {
}

// GetBySourceIdAndTargetIdAndType 通过 sourceId、targetId、eventType、获取event
func (*_Event) GetBySourceIdAndTargetIdAndType(ctx *ctx.Ctx, sourceId, targetId uint64, eventType uint) (*model.Event, error) {
	eventM := model.Event{}
	err := ctx.SqlCtl().Where("source_id = ? and target_id = ? and event_type = ?", sourceId, targetId, eventType).Take(&eventM).Error
	if err != nil {
		return nil, err
	}
	return &eventM, err
}
