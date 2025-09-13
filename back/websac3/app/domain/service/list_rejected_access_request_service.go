package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type ListRejectedAccessRequestService struct {
	listRejectedAccessRequestPort persistence.ListRejectedAccessRequestPort
	persistenceManager            db.Manager
	messageProvider               message.Provider
}

func NewListRejectedAccessRequestService(
	listRejectedAccessRequestPort persistence.ListRejectedAccessRequestPort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListRejectedAccessRequestService {
	return &ListRejectedAccessRequestService{
		listRejectedAccessRequestPort: listRejectedAccessRequestPort,
		persistenceManager:            persistenceManager,
		messageProvider:               messageProvider,
	}
}

func (s *ListRejectedAccessRequestService) Execute(
	page uint,
	perPage uint,
	filters filter.Filters,
	userID uint,
	permissions []string,
	lang string,
) ([]entity.AccessRequest, int64, error) {
	var err error
	var accessRequests []entity.AccessRequest = make([]entity.AccessRequest, 0)
	var total int64

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			accessRequests, total, err = s.listRejectedAccessRequestPort.GetRejectedByFilters(
				page,
				perPage,
				filters,
				dbCtx,
			)
			if err != nil {
				accessRequests = nil
				return err
			}

			if len(accessRequests) == 0 {
				err = errs.NewNotFoundError(
					s.messageProvider.
						WithLang(lang).
						GetMessage("list_rejected_access_request", "not_found"),
				)
			}

			return err
		},
	)
	if err != nil {
		return nil, 0, err
	}

	return accessRequests, total, nil
}
