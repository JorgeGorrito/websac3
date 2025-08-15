package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
)

type ListHigherEducationInstitutionService struct {
	getHigherEducationInstitutionPort persistence.GetHigherEducationInstitutionPort
	persistenceManager                db.Manager
	msgProvider                       message.Provider
}

func NewListHigherEducationInstitutionService(
	getHigherEducationInstitutionPort persistence.GetHigherEducationInstitutionPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListHigherEducationInstitutionService {
	return &ListHigherEducationInstitutionService{
		getHigherEducationInstitutionPort: getHigherEducationInstitutionPort,
		persistenceManager:                persistenceManager,
		msgProvider:                       msgProvider,
	}
}

func (s *ListHigherEducationInstitutionService) Execute(
	page uint,
	perPage uint,
	filters filter.Filters,
	lang string,
) ([]entity.HigherEducationInstitution, int64, error) {
	var err error
	var higherEducationInstitutions []entity.HigherEducationInstitution
	var total int64

	err = s.persistenceManager.ExecuteNonTransactional(func(ctx db.Context) error {
		higherEducationInstitutions, total, err = s.getHigherEducationInstitutionPort.GetByFilters(
			page,
			perPage,
			filters,
			ctx,
		)
		if err != nil {
			higherEducationInstitutions = nil
			return err
		}

		if len(higherEducationInstitutions) == 0 {
			err = errs.NewNotFoundError(
				s.msgProvider.
					WithLang(lang).
					GetMessage("list_higher_education_institution", "not_found"),
			)
			return err
		}

		return nil
	})

	return higherEducationInstitutions, total, err
}
