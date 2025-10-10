package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type CourseTypeRepository struct {
	Repository
}

func NewCourseTypeRepository() *CourseTypeRepository {
	return &CourseTypeRepository{}
}

func (r *CourseTypeRepository) GetAll(
	page uint,
	perPage uint,
	name string,
	ctx _db.Context,
) ([]entity.CourseType, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var courseTypes []model.CourseType
	var total int64

	// Build query
	query := dbCtx.DB().Model(&model.CourseType{})

	// Apply name filter if provided
	if name != "" {
		query = query.Joins("JOIN course_type_names ON course_types.id = course_type_names.course_type_id").
			Where("course_type_names.name ILIKE ?", "%"+name+"%")
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
		Find(&courseTypes).Error; err != nil {
		return nil, 0, err
	}

	// Map to entities
	var courseTypeEntities []entity.CourseType
	for _, courseType := range courseTypes {
		courseTypeEntity, err := mapper.Map[model.CourseType, entity.CourseType](&courseType)
		if err != nil {
			return nil, 0, err
		}
		courseTypeEntities = append(courseTypeEntities, courseTypeEntity)
	}

	return courseTypeEntities, total, nil
}

func (r *CourseTypeRepository) GetAllWithLang(
	page uint,
	perPage uint,
	name string,
	lang string,
	ctx _db.Context,
) ([]entity.CourseType, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var courseTypes []model.CourseType
	var total int64

	// Build query
	query := dbCtx.DB().Model(&model.CourseType{})

	// Apply name filter if provided
	if name != "" {
		query = query.Joins("JOIN course_type_names ON course_types.id = course_type_names.course_type_id").
			Where("course_type_names.name ILIKE ?", "%"+name+"%")
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
		Find(&courseTypes).Error; err != nil {
		return nil, 0, err
	}

	// Map to entities using language-specific mapper
	var courseTypeEntities []entity.CourseType
	for _, courseType := range courseTypes {
		courseTypeEntity, err := mapper.GetCourseTypeMapperWithLanguage(lang)(&courseType)
		if err != nil {
			return nil, 0, err
		}
		courseTypeEntities = append(courseTypeEntities, courseTypeEntity)
	}

	return courseTypeEntities, total, nil
}

func (r *CourseTypeRepository) GetByID(
	id uint,
	ctx _db.Context,
) (entity.CourseType, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.CourseType{}, err
	}

	var courseType model.CourseType
	if err := dbCtx.DB().
		Preload("Names").
		First(&courseType, id).Error; err != nil {
		return entity.CourseType{}, err
	}

	return mapper.Map[model.CourseType, entity.CourseType](&courseType)
}
