package usecase

import (
	"websac3/app/domain/entity"
)

type CreateCourseUseCase interface {
	Execute(course entity.Course, lang string) error
}
