package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type ReportRepository struct {
	Repository
}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{}
}

func (r *ReportRepository) Create(reportToSave *entity.Report, ctx _db.Context) (*entity.Report, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, err
	}

	var reportModel model.Report
	if reportModel, err = mapper.Map[entity.Report, model.Report](reportToSave); err != nil {
		return nil, err
	}

	if err := dbCtx.DB().
		Create(&reportModel).
		Error; err != nil {
		return nil, err
	}

	// Actualizar el ID en la entidad original
	reportToSave.ID = reportModel.ID

	// Guardar UnexpectedKnowledgeAreaReports por separado
	if len(reportToSave.UnexpectedKnowledgeAreaReports) > 0 {
		for _, ukar := range reportToSave.UnexpectedKnowledgeAreaReports {
			// Mapear UnexpectedKnowledgeAreaReport
			ukarModel, err := mapper.Map[entity.UnexpectedKnowledgeAreaReport, model.UnexpectedKnowledgeAreaReport](&ukar)
			if err != nil {
				return nil, err
			}
			ukarModel.ReportID = reportModel.ID

			// Guardar UnexpectedKnowledgeAreaReport
			if err := dbCtx.DB().Create(&ukarModel).Error; err != nil {
				return nil, err
			}

			// Guardar UnexpectedTopicReports
			for _, utr := range ukar.TopicReports {
				utrModel, err := mapper.Map[entity.UnexpectedTopicReport, model.UnexpectedTopicReport](&utr)
				if err != nil {
					return nil, err
				}
				utrModel.ReportID = reportModel.ID
				utrModel.UnexpectedKnowledgeAreaReportID = ukarModel.ID

				if err := dbCtx.DB().Create(&utrModel).Error; err != nil {
					return nil, err
				}
			}
		}
	}

	// Devolver la entidad guardada
	return reportToSave, nil
}

func (r *ReportRepository) GetByDegreeProgramID(degreeProgramID uint, page, perPage uint, ctx _db.Context) ([]entity.Report, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Contar total de registros
	var total int64
	if err := dbCtx.DB().
		Model(&model.Report{}).
		Where("degree_program_id = ?", degreeProgramID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calcular offset
	offset := (page - 1) * perPage

	// Obtener reportes paginados (solo campos básicos para list)
	var reports []model.Report
	if err := dbCtx.DB().
		Model(&model.Report{}).
		Select("id, degree_program_id, professional_role_id, score, created_at").
		Where("degree_program_id = ?", degreeProgramID).
		Order("created_at DESC").
		Limit(int(perPage)).
		Offset(int(offset)).
		Find(&reports).Error; err != nil {
		return nil, 0, err
	}

	var reportEntities []entity.Report
	for _, reportModel := range reports {
		// Para list solo necesitamos campos básicos
		reportEntity := entity.Report{
			ID:        reportModel.ID,
			Score:     reportModel.Score,
			CreatedAt: reportModel.CreatedAt,
		}
		reportEntities = append(reportEntities, reportEntity)
	}

	return reportEntities, total, nil
}

func (r *ReportRepository) GetByID(reportID uint, ctx _db.Context) (entity.Report, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.Report{}, err
	}

	var report model.Report
	if err := dbCtx.DB().
		// Preload DegreeProgram with all necessary fields
		Preload("DegreeProgram").
		Preload("DegreeProgram.DurationUnit").
		Preload("DegreeProgram.DurationUnit.Names").
		Preload("DegreeProgram.UserCreator").
		Preload("DegreeProgram.UserCreator.Person").
		Preload("DegreeProgram.UserCreator.Person.HigherEducationInstitution").
		Preload("DegreeProgram.UserCreator.Person.HigherEducationInstitution.Ownership").
		Preload("DegreeProgram.UserCreator.Person.HigherEducationInstitution.InstitutionalCategory").
		Preload("DegreeProgram.UserCreator.Person.HigherEducationInstitution.Municipality").
		Preload("DegreeProgram.UserCreator.Person.HigherEducationInstitution.Department").
		// Preload ProfessionalRole
		Preload("ProfessionalRole").
		// Preload KnowledgeAreaReports with TopicReports
		Preload("KnowledgeAreaReports").
		Preload("KnowledgeAreaReports.TopicReports").
		// Preload UnexpectedKnowledgeAreaReports with UnexpectedTopicReports
		Preload("UnexpectedKnowledgeAreaReports").
		Preload("UnexpectedKnowledgeAreaReports.TopicReports").
		// Preload Topic information for UnexpectedTopicReports to get localized names
		Preload("UnexpectedKnowledgeAreaReports.TopicReports.Topic").
		Preload("UnexpectedKnowledgeAreaReports.TopicReports.Topic.Names").
		Where("id = ?", reportID).
		First(&report).Error; err != nil {
		return entity.Report{}, err
	}

	// Use the mapper to get complete entity mapping
	entityReport, err := mapper.Map[model.Report, entity.Report](&report)
	if err != nil {
		return entity.Report{}, err
	}

	// Set HigherEducationInstitution from UserCreator.Person if available
	if entityReport.DegreeProgram.UserCreator != nil &&
		entityReport.DegreeProgram.UserCreator.Person != nil &&
		entityReport.DegreeProgram.UserCreator.Person.HigherEducationInstitution != nil {
		entityReport.HigherEducationInstitution = entityReport.DegreeProgram.UserCreator.Person.HigherEducationInstitution
	}

	return entityReport, nil
}
