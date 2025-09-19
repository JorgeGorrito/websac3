package usecase

import (
	"websac3/adapter/out/persistence/postgresql/model"
)

type ListCourseModelsByDegreeProgramUseCase interface {
	Execute(
		degreeProgramID uint,
		page uint,
		perPage uint,
		name string,
		lang string,
	) ([]model.Course, int64, error)
}
