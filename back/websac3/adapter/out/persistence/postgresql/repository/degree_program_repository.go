package repository

import (
	psqlfilter "websac3/adapter/out/persistence/postgresql/filter"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
	"websac3/common/mapper"
)

type DegreeProgramRepository struct {
	Repository
}

func NewDegreeProgramRepository() *DegreeProgramRepository {
	return &DegreeProgramRepository{}
}

func (r *DegreeProgramRepository) Create(degreeProgramToSave *entity.DegreeProgram, ctx _db.Context) error {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var degreeProgram model.DegreeProgram
	if degreeProgram, err = mapper.Map[entity.DegreeProgram, model.DegreeProgram](degreeProgramToSave); err != nil {
		return err
	}

	// Crear el programa de grado sin las relaciones
	if err := dbCtx.DB().
		Select("snies", "name", "total_credits", "duration_value", "duration_unit_id", "formation_level_id", "program_focus", "entry_profile", "graduate_profile", "professional_profile", "created_by", "created_at", "updated_at").
		Create(&degreeProgram).
		Error; err != nil {
		return err
	}
	degreeProgramToSave.ID = degreeProgram.ID

	// Asignar roles profesionales si existen
	if len(degreeProgramToSave.ProfessionalRoles) > 0 {
		var professionalRoleModels []model.ProfessionalRole
		for _, pr := range degreeProgramToSave.ProfessionalRoles {
			professionalRoleModels = append(professionalRoleModels, model.ProfessionalRole{
				ID: pr.ID,
			})
		}

		if err := dbCtx.DB().
			Model(&degreeProgram).
			Association("ProfessionalRoles").
			Append(professionalRoleModels); err != nil {
			return err
		}
	}

	return nil
}

func (r *DegreeProgramRepository) Update(degreeProgramToUpdate *entity.DegreeProgram, ctx _db.Context) error {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var degreeProgram model.DegreeProgram
	if degreeProgram, err = mapper.Map[entity.DegreeProgram, model.DegreeProgram](degreeProgramToUpdate); err != nil {
		return err
	}

	// Actualizar el programa de grado sin las relaciones
	if err := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Where("id = ?", degreeProgram.ID).
		Updates(map[string]interface{}{
			"snies":                degreeProgram.Snies,
			"name":                 degreeProgram.Name,
			"total_credits":        degreeProgram.TotalCredits,
			"duration_value":       degreeProgram.DurationValue,
			"duration_unit_id":     degreeProgram.DurationUnitID,
			"formation_level_id":   degreeProgram.FormationLevelID,
			"program_focus":        degreeProgram.ProgramFocus,
			"entry_profile":        degreeProgram.EntryProfile,
			"graduate_profile":     degreeProgram.GraduateProfile,
			"professional_profile": degreeProgram.ProfessionalProfile,
		}).
		Error; err != nil {
		return err
	}

	// Actualizar roles profesionales si existen
	if len(degreeProgramToUpdate.ProfessionalRoles) > 0 {
		var professionalRoleModels []model.ProfessionalRole
		for _, pr := range degreeProgramToUpdate.ProfessionalRoles {
			professionalRoleModels = append(professionalRoleModels, model.ProfessionalRole{
				ID: pr.ID,
			})
		}

		// Primero eliminar todas las asociaciones existentes
		if err := dbCtx.DB().
			Model(&degreeProgram).
			Association("ProfessionalRoles").
			Clear(); err != nil {
			return err
		}

		// Luego agregar las nuevas asociaciones
		if err := dbCtx.DB().
			Model(&degreeProgram).
			Association("ProfessionalRoles").
			Append(professionalRoleModels); err != nil {
			return err
		}
	}

	return nil
}

func (r *DegreeProgramRepository) GetByID(id uint, ctx _db.Context) (entity.DegreeProgram, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.DegreeProgram{}, err
	}

	var dp model.DegreeProgram
	if err := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Preload("DurationUnit").
		Preload("DurationUnit.Names").
		Preload("FormationLevel").
		Preload("FormationLevel.Names").
		Preload("ProfessionalRoles").
		Preload("UserCreator").
		Preload("UserCreator.Person").
		Preload("UserCreator.Person.HigherEducationInstitution").
		Preload("UserCreator.Person.HigherEducationInstitution.Ownership").
		Preload("UserCreator.Person.HigherEducationInstitution.InstitutionalCategory").
		Preload("UserCreator.Person.HigherEducationInstitution.Municipality").
		Preload("UserCreator.Person.HigherEducationInstitution.Department").
		Preload("Courses").
		Preload("Courses.Nature").
		Preload("Courses.Type").
		Preload("Courses.CourseTopics").
		Preload("Courses.CourseTopics.Topic").
		Preload("Courses.CourseTopics.Topic.Names").
		Where("degree_programs.id = ?", id).
		First(&dp).Error; err != nil {
		return entity.DegreeProgram{}, err
	}

	return mapper.Map[model.DegreeProgram, entity.DegreeProgram](&dp)
}

func (r *DegreeProgramRepository) GetByIDWithLang(id uint, lang string, ctx _db.Context) (entity.DegreeProgram, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.DegreeProgram{}, err
	}

	var dp model.DegreeProgram
	if err := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Preload("DurationUnit").
		Preload("DurationUnit.Names").
		Preload("FormationLevel").
		Preload("FormationLevel.Names").
		Preload("ProfessionalRoles").
		Preload("UserCreator").
		Preload("UserCreator.Person").
		Preload("UserCreator.Person.HigherEducationInstitution").
		Preload("UserCreator.Person.HigherEducationInstitution.Ownership").
		Preload("UserCreator.Person.HigherEducationInstitution.InstitutionalCategory").
		Preload("UserCreator.Person.HigherEducationInstitution.Municipality").
		Preload("UserCreator.Person.HigherEducationInstitution.Department").
		Preload("Courses").
		Preload("Courses.Nature").
		Preload("Courses.Type").
		Preload("Courses.CourseTopics").
		Preload("Courses.CourseTopics.Topic").
		Preload("Courses.CourseTopics.Topic.Names").
		Where("degree_programs.id = ?", id).
		First(&dp).Error; err != nil {
		return entity.DegreeProgram{}, err
	}

	// Use language-specific mapper if available, otherwise fallback to default
	var degreeProgram entity.DegreeProgram
	if langMapper := mapper.GetDegreeProgramMapperWithLanguage(lang); langMapper != nil {
		degreeProgram, err = langMapper(&dp)
		if err != nil {
			return entity.DegreeProgram{}, err
		}
	} else {
		// Fallback to default mapper
		if degreeProgram, err = mapper.Map[model.DegreeProgram, entity.DegreeProgram](&dp); err != nil {
			return entity.DegreeProgram{}, err
		}
	}

	return degreeProgram, nil
}

func (r *DegreeProgramRepository) GetByFilters(page, perPage uint, filters filter.Filters, ctx _db.Context) ([]entity.DegreeProgram, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Preload("DurationUnit").
		Preload("DurationUnit.Names").
		Preload("FormationLevel").
		Preload("FormationLevel.Names").
		Preload("ProfessionalRoles").
		Preload("UserCreator").
		Preload("UserCreator.Person").
		Preload("UserCreator.Person.HigherEducationInstitution")

	dbCtx.DBSet(baseQuery)

	dbCtxFiltered, err := filters.Apply(dbCtx, psqlfilter.FiltersRegistry)
	if err != nil {
		return nil, 0, err
	}

	dbCtx, err = r.CastDbContext(dbCtxFiltered)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	if err := dbCtx.DB().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var degreeProgramsFound []model.DegreeProgram
	if err := dbCtx.DB().Find(&degreeProgramsFound).Error; err != nil {
		return nil, 0, err
	}

	var degreePrograms []entity.DegreeProgram
	for _, degreeProgram := range degreeProgramsFound {
		var degreeProgramEntity entity.DegreeProgram
		if degreeProgramEntity, err = mapper.Map[model.DegreeProgram, entity.DegreeProgram](&degreeProgram); err != nil {
			return nil, 0, err
		}
		degreePrograms = append(degreePrograms, degreeProgramEntity)
	}

	return degreePrograms, total, nil
}

func (r *DegreeProgramRepository) GetByUserIDAndFilters(page, perPage uint, userID uint, filters filter.Filters, ctx _db.Context) ([]entity.DegreeProgram, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Preload("DurationUnit").
		Preload("DurationUnit.Names").
		Preload("FormationLevel").
		Preload("FormationLevel.Names").
		Preload("ProfessionalRoles").
		Preload("UserCreator").
		Preload("UserCreator.Person").
		Preload("UserCreator.Person.HigherEducationInstitution").
		Where("created_by = ?", userID)

	dbCtx.DBSet(baseQuery)

	dbCtxFiltered, err := filters.Apply(dbCtx, psqlfilter.FiltersRegistry)
	if err != nil {
		return nil, 0, err
	}

	dbCtx, err = r.CastDbContext(dbCtxFiltered)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	if err := dbCtx.DB().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var degreeProgramsFound []model.DegreeProgram
	if err := dbCtx.DB().Find(&degreeProgramsFound).Error; err != nil {
		return nil, 0, err
	}

	var degreePrograms []entity.DegreeProgram
	for _, degreeProgram := range degreeProgramsFound {
		var degreeProgramEntity entity.DegreeProgram
		if degreeProgramEntity, err = mapper.Map[model.DegreeProgram, entity.DegreeProgram](&degreeProgram); err != nil {
			return nil, 0, err
		}
		degreePrograms = append(degreePrograms, degreeProgramEntity)
	}

	return degreePrograms, total, nil
}
