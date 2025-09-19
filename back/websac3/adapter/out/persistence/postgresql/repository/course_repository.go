package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"

	"gorm.io/gorm"
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

func (r *CourseRepository) GetByDegreeProgramID(
	degreeProgramID uint,
	page uint,
	perPage uint,
	name string,
	ctx _db.Context,
) ([]entity.Course, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var courses []model.Course
	var total int64

	query := dbCtx.DB().Model(&model.Course{}).
		Where("degree_program_id = ?", degreeProgramID)

	// Apply name filter if provided
	if name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and get results
	offset := (page - 1) * perPage
	if err := query.
		Preload("Nature", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Names")
		}).
		Preload("Type", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Names")
		}).
		Preload("DegreeProgram").
		Preload("UserCreator.Person").
		Offset(int(offset)).
		Limit(int(perPage)).
		Order("period_number ASC, name ASC").
		Find(&courses).Error; err != nil {
		return nil, 0, err
	}

	// Map to entities
	var courseEntities []entity.Course
	for _, course := range courses {
		courseEntity, err := mapper.Map[model.Course, entity.Course](&course)
		if err != nil {
			return nil, 0, err
		}
		courseEntities = append(courseEntities, courseEntity)
	}

	return courseEntities, total, nil
}

func (r *CourseRepository) GetModelsByDegreeProgramID(
	degreeProgramID uint,
	page uint,
	perPage uint,
	name string,
	ctx _db.Context,
) ([]model.Course, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var courses []model.Course
	var total int64

	query := dbCtx.DB().Model(&model.Course{}).
		Where("degree_program_id = ?", degreeProgramID)

	// Apply name filter if provided
	if name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and get results
	offset := (page - 1) * perPage
	if err := query.
		Preload("Nature", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Names")
		}).
		Preload("Type", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Names")
		}).
		Preload("DegreeProgram").
		Preload("UserCreator.Person").
		Offset(int(offset)).
		Limit(int(perPage)).
		Order("period_number ASC, name ASC").
		Find(&courses).Error; err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}
