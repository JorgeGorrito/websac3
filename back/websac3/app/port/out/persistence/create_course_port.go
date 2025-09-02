package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreateCoursePort interface {
	Create(course *entity.Course, db db.Context) error
}
