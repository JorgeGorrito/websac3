package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type CourseNatureRepository struct {
	Repository
}

func NewCourseNatureRepository() *CourseNatureRepository {
	return &CourseNatureRepository{}
}

func (r *CourseNatureRepository) GetAll(
	page uint,
	perPage uint,
	name string,
	ctx _db.Context,
) ([]entity.CourseNature, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var courseNatures []model.CourseNature
	var total int64

	// Build query
	query := dbCtx.DB().Model(&model.CourseNature{})

	// Apply name filter if provided
	if name != "" {
		query = query.Joins("JOIN course_nature_names ON course_natures.id = course_nature_names.course_nature_id").
			Where("course_nature_names.name ILIKE ?", "%"+name+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * perPage

	// Execute query with pagination
	if err := query.
		Preload("Names").
		Offset(int(offset)).
		Limit(int(perPage)).
		Order("id ASC").
		Find(&courseNatures).Error; err != nil {
		return nil, 0, err
	}

	// Map to entities
	var courseNatureEntities []entity.CourseNature
	for _, courseNature := range courseNatures {
		courseNatureEntity, err := mapper.Map[model.CourseNature, entity.CourseNature](&courseNature)
		if err != nil {
			return nil, 0, err
		}
		courseNatureEntities = append(courseNatureEntities, courseNatureEntity)
	}

	return courseNatureEntities, total, nil
}

func (r *CourseNatureRepository) GetAllWithLang(
	page uint,
	perPage uint,
	name string,
	lang string,
	ctx _db.Context,
) ([]entity.CourseNature, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var courseNatures []model.CourseNature
	var total int64

	// Build query
	query := dbCtx.DB().Model(&model.CourseNature{})

	// Apply name filter if provided
	if name != "" {
		query = query.Joins("JOIN course_nature_names ON course_natures.id = course_nature_names.course_nature_id").
			Where("course_nature_names.name ILIKE ?", "%"+name+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * perPage

	// Execute query with pagination
	if err := query.
		Preload("Names").
		Offset(int(offset)).
		Limit(int(perPage)).
		Order("id ASC").
		Find(&courseNatures).Error; err != nil {
		return nil, 0, err
	}

	// Map to entities using language-specific mapper
	var courseNatureEntities []entity.CourseNature
	for _, courseNature := range courseNatures {
		courseNatureEntity, err := mapper.GetCourseNatureMapperWithLanguage(lang)(&courseNature)
		if err != nil {
			return nil, 0, err
		}
		courseNatureEntities = append(courseNatureEntities, courseNatureEntity)
	}

	return courseNatureEntities, total, nil
}

func (r *CourseNatureRepository) GetByID(
	id uint,
	ctx _db.Context,
) (entity.CourseNature, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.CourseNature{}, err
	}

	var courseNature model.CourseNature
	if err := dbCtx.DB().
		Preload("Names").
		First(&courseNature, id).Error; err != nil {
		return entity.CourseNature{}, err
	}

	return mapper.Map[model.CourseNature, entity.CourseNature](&courseNature)
}
