package service

import (
	"errors"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type CreateCourseService struct {
	createCoursePort   persistence.CreateCoursePort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewCreateCourseService(
	createCoursePort persistence.CreateCoursePort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *CreateCourseService {
	return &CreateCourseService{
		createCoursePort:   createCoursePort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *CreateCourseService) Execute(
	course entity.Course,
	lang string,
) error {
	return s.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			if err := s.createCoursePort.Create(&course, ctx); err != nil {
				if errors.Is(err, errs.ConflictError) {
					return errs.NewConflictError(s.msgProvider.WithLang(lang).GetMessage("create_course", "course_already_exists"))
				}
				return err
			}
			return nil
		},
	)
}
