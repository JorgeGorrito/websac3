package usecase

import (
	"websac3/app/domain/entity"
)

type GetCourseTopicsByCourseIDUseCase interface {
	Execute(
		courseID uint,
		lang string,
	) ([]entity.CourseTopic, *entity.Course, error)
}
