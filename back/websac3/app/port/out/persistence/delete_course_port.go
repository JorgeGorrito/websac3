package persistence

import (
	"websac3/app/port/out/persistence/db"
)

type DeleteCoursePort interface {
	DeleteByID(
		courseID uint,
		ctx db.Context,
	) error
}
