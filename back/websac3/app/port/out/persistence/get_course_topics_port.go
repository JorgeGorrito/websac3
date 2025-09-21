package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetCourseTopicsPort interface {
	GetTopicsByCourseID(
		courseID uint,
		ctx db.Context,
	) ([]entity.CourseTopic, error)

	GetTopicsWithCourseByCourseID(
		courseID uint,
		ctx db.Context,
	) ([]entity.CourseTopic, *entity.Course, error)

	GetTopicsWithCourseByCourseIDAndLang(
		courseID uint,
		lang string,
		ctx db.Context,
	) ([]entity.CourseTopic, *entity.Course, error)
}
