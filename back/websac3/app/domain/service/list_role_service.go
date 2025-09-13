package service

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type ListRoleService struct {
	listRolePort       persistence.ListRolePort
	persistenceManager db.Manager
	messageProvider    message.Provider
}

func NewListRoleService(
	listRolePort persistence.ListRolePort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListRoleService {
	return &ListRoleService{
		listRolePort:       listRolePort,
		persistenceManager: persistenceManager,
		messageProvider:    messageProvider,
	}
}

func (s *ListRoleService) Execute(
	page uint,
	perPage uint,
	filters filter.Filters,
	lang string,
) ([]entity.Role, int64, error) {
	var err error
	var roles []entity.Role = make([]entity.Role, 0)
	var total int64

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			roles, total, err = s.listRolePort.GetByFilters(
				page,
				perPage,
				filters,
				dbCtx,
			)
			if err != nil {
				roles = nil
				return err
			}

			// Si no hay roles, devolver lista vacía en lugar de error
			// if len(roles) == 0 {
			// 	err = errs.NewNotFoundError(
			// 		s.messageProvider.
			// 			WithLang(lang).
			// 			GetMessage("list_role", "not_found"),
			// 	)
			// }

			return err
		},
	)
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}
