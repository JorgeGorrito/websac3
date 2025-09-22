package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ListCourseTypesService struct {
	getCourseTypePort  persistence.GetCourseTypePort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewListCourseTypesService(
	getCourseTypePort persistence.GetCourseTypePort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListCourseTypesService {
	return &ListCourseTypesService{
		getCourseTypePort:  getCourseTypePort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *ListCourseTypesService) Execute(
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]entity.CourseType, int64, error) {
	var (
		err         error
		courseTypes []entity.CourseType = make([]entity.CourseType, 0)
		total       int64
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			courseTypes, total, err = s.getCourseTypePort.GetAllWithLang(page, perPage, name, lang, dbCtx)
			if err != nil {
				courseTypes = nil
				total = 0
				return err
			}

			if len(courseTypes) == 0 {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("list_course_types", "not_found"),
				)
			}

			return nil
		},
	)

	return courseTypes, total, err
}
