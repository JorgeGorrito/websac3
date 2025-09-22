package usecase

import (
	"websac3/app/domain/entity"
)

type ListCourseTypesUseCase interface {
	Execute(
		page uint,
		perPage uint,
		name string,
		lang string,
	) ([]entity.CourseType, int64, error)
}
