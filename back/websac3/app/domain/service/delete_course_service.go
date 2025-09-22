package service

import (
	"errors"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type DeleteCourseService struct {
	deleteCoursePort   persistence.DeleteCoursePort
	getCoursePort      persistence.GetCoursePort
	getUserPort        persistence.GetUserPort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewDeleteCourseService(
	deleteCoursePort persistence.DeleteCoursePort,
	getCoursePort persistence.GetCoursePort,
	getUserPort persistence.GetUserPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *DeleteCourseService {
	return &DeleteCourseService{
		deleteCoursePort:   deleteCoursePort,
		getCoursePort:      getCoursePort,
		getUserPort:        getUserPort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *DeleteCourseService) Execute(
	courseID uint,
	userID uint,
	lang string,
) error {
	return s.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			// Get the user to check if they are admin
			user, err := s.getUserPort.GetByID(userID, ctx)
			if err != nil {
				return err
			}

			// Get the course (including soft deleted ones) to check ownership and permissions
			course, err := s.getCoursePort.GetByIDWithDeleted(courseID, ctx)
			if err != nil {
				if errors.Is(err, errs.NotFoundError) {
					return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("delete_course", "course_not_found"))
				}
				return err
			}

			// Check if the course is already deleted
			if course.DeletedAt != nil {
				return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("delete_course", "course_not_found"))
			}

			// Use business rules from entities to check permissions
			if !user.IsAdmin() && !course.CanBeDeletedBy(userID) {
				return errs.NewValidationError(s.msgProvider.WithLang(lang).GetMessage("delete_course", "insufficient_permissions"))
			}

			// Proceed with deletion
			if err := s.deleteCoursePort.DeleteByID(courseID, ctx); err != nil {
				return err
			}
			return nil
		},
	)
}
