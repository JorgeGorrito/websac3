package service

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ListCourseModelsByDegreeProgramService struct {
	getCoursePort      persistence.GetCoursePort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewListCourseModelsByDegreeProgramService(
	getCoursePort persistence.GetCoursePort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *ListCourseModelsByDegreeProgramService {
	return &ListCourseModelsByDegreeProgramService{
		getCoursePort:      getCoursePort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *ListCourseModelsByDegreeProgramService) Execute(
	degreeProgramID uint,
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]model.Course, int64, error) {
	var (
		err     error
		courses []model.Course = make([]model.Course, 0)
		total   int64
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			courses, total, err = s.getCoursePort.GetModelsByDegreeProgramID(
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
