package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type ListProfessionalRoleService struct {
	listProfessionalRolePort persistence.ListProfessionalRolePort
	persistenceManager       db.Manager
	messageProvider          message.Provider
}

func NewListProfessionalRoleService(
	listProfessionalRolePort persistence.ListProfessionalRolePort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListProfessionalRoleService {
	return &ListProfessionalRoleService{
		listProfessionalRolePort: listProfessionalRolePort,
		persistenceManager:       persistenceManager,
		messageProvider:          messageProvider,
	}
}

func (s *ListProfessionalRoleService) Execute(
	page uint,
	perPage uint,
	filters filter.Filters,
	lang string,
) ([]entity.ProfessionalRole, int64, error) {
	var err error
	var professionalRoles []entity.ProfessionalRole = make([]entity.ProfessionalRole, 0)
	var total int64

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			professionalRoles, total, err = s.listProfessionalRolePort.GetByFilters(
				page,
				perPage,
				filters,
				dbCtx,
			)
			if err != nil {
				professionalRoles = nil
				return err
			}

			if len(professionalRoles) == 0 {
				err = errs.NewNotFoundError(
					s.messageProvider.
						WithLang(lang).
						GetMessage("list_professional_role", "not_found"),
				)
			}

			return err
		},
	)
	if err != nil {
		return nil, 0, err
	}

	return professionalRoles, total, nil
}
