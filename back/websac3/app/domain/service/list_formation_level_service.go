package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ListFormationLevelService struct {
	getFormationLevelPort persistence.GetFormationLevelPort
	persistenceManager    db.Manager
	msgProvider           message.Provider
}

func NewListFormationLevelService(
	getFormationLevelPort persistence.GetFormationLevelPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListFormationLevelService {
	return &ListFormationLevelService{
		getFormationLevelPort: getFormationLevelPort,
		persistenceManager:    persistenceManager,
		msgProvider:           msgProvider,
	}
}

func (s *ListFormationLevelService) Execute(
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]entity.FormationLevel, int64, error) {
	var (
		err             error
		formationLevels []entity.FormationLevel = make([]entity.FormationLevel, 0)
		total           int64
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			formationLevels, total, err = s.getFormationLevelPort.GetAllWithLang(
				page,
				perPage,
				name,
				lang,
				dbCtx,
			)
			if err != nil {
				formationLevels = nil
				return err
			}

			if len(formationLevels) == 0 {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("list_formation_level", "not_found"),
				)
			}

			return nil
		},
	)

	return formationLevels, total, err
}
