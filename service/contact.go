package service

import (
	"github.com/kiririx/krim/conf"
	"github.com/kiririx/krim/constx"
	"github.com/kiririx/krim/model"
)

var ContactService = &contactService{}

type contactService struct{}

func (*contactService) AddContactEvent(sourceId, targetId uint64, event uint) error {
	return conf.Sqlx.Save(&model.Event{
		SourceId:  sourceId,
		TargetId:  targetId,
		EventType: event,
		Progress:  constx.EventProgress,
	}).Error
}
