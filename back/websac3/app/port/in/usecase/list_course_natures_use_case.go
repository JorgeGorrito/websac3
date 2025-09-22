package usecase

import (
	"websac3/app/domain/entity"
)

type ListCourseNaturesUseCase interface {
	Execute(
		page uint,
		perPage uint,
		name string,
		lang string,
	) ([]entity.CourseNature, int64, error)
}
