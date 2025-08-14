package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type ListIdentificationTypeService struct {
	getIdentificationTypePort persistence.GetIdentificationTypePort
	persistenceManager        db.Manager
	messageProvider           message.Provider
}

func NewListIdentificationTypeService(
	getIdentificationTypePort persistence.GetIdentificationTypePort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListIdentificationTypeService {
	return &ListIdentificationTypeService{
		getIdentificationTypePort: getIdentificationTypePort,
		persistenceManager:        persistenceManager,
		messageProvider:           messageProvider,
	}
}

func (s *ListIdentificationTypeService) Execute(
	page uint,
	perPage uint,
	filters filter.Filters,
	lang string,
) ([]entity.IdentificationType, int64, error) {
	var err error
	var identificationTypes []entity.IdentificationType = make([]entity.IdentificationType, 0)
	var total int64

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			identificationTypes, total, err = s.getIdentificationTypePort.GetByFilters(
				page,
				perPage,
				filters,
				dbCtx,
			)
			if err != nil {
				identificationTypes = nil
				return err
			}

			if len(identificationTypes) == 0 {
				err = errs.NewNotFoundError(
					s.messageProvider.
						WithLang(lang).
						GetMessage("list_identification_type", "not_found"),
				)
			}

			return err
		},
	)
	if err != nil {
		return nil, 0, err
	}

	return identificationTypes, total, nil
}
