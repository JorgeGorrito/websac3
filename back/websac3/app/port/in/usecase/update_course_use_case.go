package usecase

import (
	"websac3/app/domain/entity"
)

type UpdateCourseUseCase interface {
	Execute(course entity.Course, userID uint, lang string) error
}
