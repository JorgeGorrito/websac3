package usecase

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
)

type UpdateCourseUseCaseImpl struct {
	updateCourseService *service.UpdateCourseService
}

func NewUpdateCourseUseCaseImpl(
	updateCourseService *service.UpdateCourseService,
) usecase.UpdateCourseUseCase {
	return &UpdateCourseUseCaseImpl{
		updateCourseService: updateCourseService,
	}
}

func (u *UpdateCourseUseCaseImpl) Execute(course entity.Course, userID uint, lang string) error {
	return u.updateCourseService.Execute(course, userID, lang)
}
