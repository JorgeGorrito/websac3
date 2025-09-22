package usecase

import (
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
)

type DeleteCourseUseCaseImpl struct {
	deleteCourseService *service.DeleteCourseService
}

func NewDeleteCourseUseCaseImpl(
	deleteCourseService *service.DeleteCourseService,
) usecase.DeleteCourseUseCase {
	return &DeleteCourseUseCaseImpl{
		deleteCourseService: deleteCourseService,
	}
}

func (u *DeleteCourseUseCaseImpl) Execute(courseID uint, userID uint, lang string) error {
	return u.deleteCourseService.Execute(courseID, userID, lang)
}
