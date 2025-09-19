package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ListCourseByDegreeProgramService struct {
	getCoursePort      persistence.GetCoursePort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewListCourseByDegreeProgramService(
	getCoursePort persistence.GetCoursePort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListCourseByDegreeProgramService {
	return &ListCourseByDegreeProgramService{
		getCoursePort:      getCoursePort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *ListCourseByDegreeProgramService) Execute(
	degreeProgramID uint,
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]entity.Course, int64, error) {
	var (
		err     error
		courses []entity.Course = make([]entity.Course, 0)
		total   int64
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			courses, total, err = s.getCoursePort.GetByDegreeProgramID(
				degreeProgramID,
				page,
				perPage,
				name,
				dbCtx,
			)
			if err != nil {
				courses = nil
				return err
			}

			if len(courses) == 0 {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("list_course_by_degree_program", "not_found"),
				)
			}

			return nil
		},
	)

	return courses, total, err
}
