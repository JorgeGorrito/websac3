package service

import (
	"errors"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type UpdateCourseService struct {
	updateCoursePort   persistence.UpdateCoursePort
	getCoursePort      persistence.GetCoursePort
	getUserPort        persistence.GetUserPort
	persistenceManager db.Manager
	msgProvider        message.Provider
}

func NewUpdateCourseService(
	updateCoursePort persistence.UpdateCoursePort,
	getCoursePort persistence.GetCoursePort,
	getUserPort persistence.GetUserPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *UpdateCourseService {
	return &UpdateCourseService{
		updateCoursePort:   updateCoursePort,
		getCoursePort:      getCoursePort,
		getUserPort:        getUserPort,
		persistenceManager: persistenceManager,
		msgProvider:        msgProvider,
	}
}

func (s *UpdateCourseService) Execute(
	course entity.Course,
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

			// Get the existing course to check ownership and permissions
			existingCourse, err := s.getCoursePort.GetByID(course.ID, ctx)
			if err != nil {
				if errors.Is(err, errs.NotFoundError) {
					return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("update_course", "course_not_found"))
				}
				return err
			}

			// Use business rules from entities to check permissions
			if !user.IsAdmin() && !existingCourse.CanBeUpdatedBy(userID) {
				return errs.NewValidationError(s.msgProvider.WithLang(lang).GetMessage("update_course", "insufficient_permissions"))
			}

			// Preserve the CreatedBy field from the existing course
			course.CreatedBy = existingCourse.CreatedBy

			// Proceed with update
			if err := s.updateCoursePort.Update(&course, ctx); err != nil {
				if errors.Is(err, errs.ConflictError) {
					return errs.NewConflictError(s.msgProvider.WithLang(lang).GetMessage("update_course", "course_code_already_exists"))
				}
				return err
			}
			return nil
		},
	)
}
