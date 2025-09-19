package usecase

import (
	"websac3/app/domain/entity"
)

type ListCourseByDegreeProgramUseCase interface {
	Execute(
		degreeProgramID uint,
		page uint,
		perPage uint,
		name string,
		lang string,
	) ([]entity.Course, int64, error)
}
