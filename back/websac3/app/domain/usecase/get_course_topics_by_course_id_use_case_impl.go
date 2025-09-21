package usecase

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/service"
)

type GetCourseTopicsByCourseIDUseCaseImpl struct {
	getCourseTopicsByCourseIDService *service.GetCourseTopicsByCourseIDService
}

func NewGetCourseTopicsByCourseIDUseCaseImpl(
	getCourseTopicsByCourseIDService *service.GetCourseTopicsByCourseIDService,
) *GetCourseTopicsByCourseIDUseCaseImpl {
	return &GetCourseTopicsByCourseIDUseCaseImpl{
		getCourseTopicsByCourseIDService: getCourseTopicsByCourseIDService,
	}
}

func (u *GetCourseTopicsByCourseIDUseCaseImpl) Execute(
	courseID uint,
	lang string,
) ([]entity.CourseTopic, *entity.Course, error) {
	return u.getCourseTopicsByCourseIDService.Execute(courseID, lang)
}
