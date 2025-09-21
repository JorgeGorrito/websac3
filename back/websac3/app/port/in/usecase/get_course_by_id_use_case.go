package usecase

import (
	"websac3/app/domain/entity"
)

type GetCourseByIDUseCase interface {
	Execute(
		courseID uint,
		lang string,
	) (*entity.Course, error)
}
