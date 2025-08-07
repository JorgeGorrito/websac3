package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type ListAccessRequestService struct {
	getAccessRequestPort persistence.GetAccessRequestPort
	persistenceManager   db.Manager
	messageProvider      message.Provider
}

func NewListAccessRequestService(getAccessRequestPort persistence.GetAccessRequestPort, persistenceManager db.Manager, messageProvider message.Provider) *ListAccessRequestService {
	return &ListAccessRequestService{
		getAccessRequestPort: getAccessRequestPort,
		persistenceManager:   persistenceManager,
		messageProvider:      messageProvider,
	}
}

func (s *ListAccessRequestService) Execute(page, perPage uint, filters filter.Filters, lang string) ([]entity.AccessRequest, int64, error) {
	var err error
	var accessRequests []entity.AccessRequest = make([]entity.AccessRequest, 0)
	var total int64

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {

			accessRequests, total, err = s.getAccessRequestPort.GetAuthenticatedEmailByFilters(
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
						GetMessage("list_access_request", "not_found"),
				)
			}

			return err
		},
	)

	return accessRequests, total, err
}
