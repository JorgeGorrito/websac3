package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ListCourseNaturesService struct {
	getCourseNaturePort persistence.GetCourseNaturePort
	persistenceManager  db.Manager
	msgProvider         message.Provider
}

func NewListCourseNaturesService(
	getCourseNaturePort persistence.GetCourseNaturePort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListCourseNaturesService {
	return &ListCourseNaturesService{
		getCourseNaturePort: getCourseNaturePort,
		persistenceManager:  persistenceManager,
		msgProvider:         msgProvider,
	}
}

func (s *ListCourseNaturesService) Execute(
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]entity.CourseNature, int64, error) {
	var (
		err           error
		courseNatures []entity.CourseNature = make([]entity.CourseNature, 0)
		total         int64
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			courseNatures, total, err = s.getCourseNaturePort.GetAllWithLang(page, perPage, name, lang, dbCtx)
			if err != nil {
				courseNatures = nil
				total = 0
				return err
			}

			if len(courseNatures) == 0 {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("list_course_natures", "not_found"),
				)
			}

			return nil
		},
	)

	return courseNatures, total, err
}
