package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type GetCourseByIDService struct {
	getCoursePort      persistence.GetCoursePort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewGetCourseByIDService(
	getCoursePort persistence.GetCoursePort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *GetCourseByIDService {
	return &GetCourseByIDService{
		getCoursePort:      getCoursePort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *GetCourseByIDService) Execute(
	courseID uint,
	lang string,
) (*entity.Course, error) {
	var (
		err    error
		course *entity.Course
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			course, err = s.getCoursePort.GetByIDWithLang(courseID, lang, dbCtx)
			if err != nil {
				return err
			}

			if course == nil {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("get_course_by_id", "not_found"),
				)
			}

			return nil
		},
	)

	return course, err
}
