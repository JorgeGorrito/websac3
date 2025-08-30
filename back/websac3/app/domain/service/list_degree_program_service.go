package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
	"websac3/app/port/out/persistence/filter"
)

type ListDegreeProgramService struct {
	getDegreeProgramPort persistence.GetDegreeProgramPort
	roleEnum             enum.RoleEnum
	persistenceManager   db.Manager
	msgProvider          message.Provider
}

func NewListDegreeProgramService(
	getDegreeProgramPort persistence.GetDegreeProgramPort,
	roleEnum enum.RoleEnum,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListDegreeProgramService {
	return &ListDegreeProgramService{
		getDegreeProgramPort: getDegreeProgramPort,
		roleEnum:             roleEnum,
		persistenceManager:   persistenceManager,
		msgProvider:          msgProvider,
	}
}

func (s *ListDegreeProgramService) Execute(
	page uint,
	perPage uint,
	filters filter.Filters,
	userRole string,
	userID uint,
	lang string,
) ([]entity.DegreeProgram, int64, error) {
	var (
		err           error
		durationUnits []entity.DegreeProgram = make([]entity.DegreeProgram, 0)
		total         int64
	)

	// Validar que el rol del usuario existe
	role, err := s.roleEnum.GetByName(userRole)
	if err != nil {
		return nil, 0, err
	}

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			// Si el usuario es admin, traer todos los degree programs
			if role.IsAdmin() || role.IsCybersecurityAuditor() {
				durationUnits, total, err = s.getDegreeProgramPort.GetByFilters(
					page,
					perPage,
					filters,
					dbCtx,
				)
			} else {
				// Si no es admin, traer solo los degree programs creados por el usuario
				durationUnits, total, err = s.getDegreeProgramPort.GetByIDAndFilters(
					page,
					perPage,
					userID,
					filters,
					dbCtx,
				)
			}

			if err != nil {
				durationUnits = nil
				return err
			}

			if len(durationUnits) == 0 {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("list_degree_program", "not_found"),
				)
			}

			return nil
		},
	)

	return durationUnits, total, err
}
