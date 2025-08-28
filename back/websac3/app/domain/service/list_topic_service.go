package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ListTopicService struct {
	getTopicPort       persistence.GetTopicPort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewListTopicService(
	getTopicPort persistence.GetTopicPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListTopicService {
	return &ListTopicService{
		getTopicPort:       getTopicPort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *ListTopicService) Execute(
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]entity.Topic, int64, error) {
	var (
		err    error
		topics []entity.Topic = make([]entity.Topic, 0)
		total  int64
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			topics, total, err = s.getTopicPort.GetByNameAndLang(
				page,
				perPage,
				name,
				lang,
				dbCtx,
			)
			if err != nil {
				topics = nil
				return err
			}

			if len(topics) == 0 {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("list_topic", "not_found"),
				)
			}

			return nil
		},
	)

	return topics, total, err
}
