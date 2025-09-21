package usecase

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/service"
)

type GetCourseByIDUseCaseImpl struct {
	getCourseByIDService *service.GetCourseByIDService
}

func NewGetCourseByIDUseCaseImpl(
	getCourseByIDService *service.GetCourseByIDService,
) *GetCourseByIDUseCaseImpl {
	return &GetCourseByIDUseCaseImpl{
		getCourseByIDService: getCourseByIDService,
	}
}

func (u *GetCourseByIDUseCaseImpl) Execute(
	courseID uint,
	lang string,
) (*entity.Course, error) {
	return u.getCourseByIDService.Execute(courseID, lang)
}
