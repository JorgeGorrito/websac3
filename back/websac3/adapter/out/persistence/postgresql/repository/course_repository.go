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

	// Persist course topics if provided
	if len(courseToSave.CourseTopics) > 0 {
		var batch []model.CourseTopic
		for _, ct := range courseToSave.CourseTopics {
			batch = append(batch, model.CourseTopic{
				CourseID:   course.ID,
				TopicID:    ct.TopicID,
				StudyHours: uint(ct.StudyHours),
			})
		}
		if err := dbCtx.DB().Create(&batch).Error; err != nil {
			return err
		}
	}

	return nil
}
