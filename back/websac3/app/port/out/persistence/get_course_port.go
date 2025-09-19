package persistence

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetCoursePort interface {
	GetByDegreeProgramID(
		degreeProgramID uint,
		page uint,
		perPage uint,
		name string,
		ctx db.Context,
	) ([]entity.Course, int64, error)

	GetModelsByDegreeProgramID(
		degreeProgramID uint,
		page uint,
		perPage uint,
		name string,
		ctx db.Context,
	) ([]model.Course, int64, error)
}
