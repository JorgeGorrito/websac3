package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type CourseRepository struct {
	Repository
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

func (r *CourseRepository) Create(courseToSave *entity.Course, ctx _db.Context) error {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var course model.Course
	if course, err = mapper.Map[entity.Course, model.Course](courseToSave); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Create(&course).
		Error; err != nil {
		// Check if it's a unique constraint violation
		if IsUniqueConstraintViolation(err) {
			return errs.NewConflictError(err.Error())
		}

		return err
	}
	courseToSave.ID = course.ID

	return nil
}
