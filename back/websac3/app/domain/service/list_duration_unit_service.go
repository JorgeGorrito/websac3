package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ListDurationUnitService struct {
	getDurationUnitPort persistence.GetDurationUnitPort
	persistenceManager  db.Manager
	msgProvider         message.Provider
}

func NewListDurationUnitService(
	getDurationUnitPort persistence.GetDurationUnitPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListDurationUnitService {
	return &ListDurationUnitService{
		getDurationUnitPort: getDurationUnitPort,
		persistenceManager:  persistenceManager,
		msgProvider:         msgProvider,
	}
}

func (s *ListDurationUnitService) Execute(
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]entity.DurationUnit, int64, error) {
	var (
		err            error
		durationUnits  []entity.DurationUnit = make([]entity.DurationUnit, 0)
		total          int64
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			durationUnits, total, err = s.getDurationUnitPort.GetByNameAndLang(
				page,
				perPage,
				name,
				lang,
				dbCtx,
			)
			if err != nil {
				durationUnits = nil
				return err
			}

			if len(durationUnits) == 0 {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("list_duration_unit", "not_found"),
				)
			}

			return nil
		},
	)

	return durationUnits, total, err
}
