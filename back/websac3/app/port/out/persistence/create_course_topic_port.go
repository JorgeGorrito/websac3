package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreateCourseTopicPort interface {
	CreateCourseTopic(courseTopic *entity.CourseTopic, ctx db.Context) error
}



