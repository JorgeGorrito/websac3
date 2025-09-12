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

	// Devolver la entidad guardada
	return reportToSave, nil
}

func (r *ReportRepository) GetByDegreeProgramID(degreeProgramID uint, ctx _db.Context) ([]entity.Report, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, err
	}

	var reports []model.Report
	if err := dbCtx.DB().
		Model(&model.Report{}).
		Preload("DegreeProgram").
		Preload("DegreeProgram.DurationUnit").
		Preload("DegreeProgram.DurationUnit.Names").
		Preload("DegreeProgram.UserCreator").
		Preload("DegreeProgram.UserCreator.Person").
		Preload("DegreeProgram.UserCreator.Person.HigherEducationInstitution").
		Preload("ProfessionalRole").
		Preload("KnowledgeAreaReports").
		Preload("KnowledgeAreaReports.TopicReports").
		Preload("KnowledgeAreaReports.TopicReports.Topic").
		Preload("KnowledgeAreaReports.TopicReports.Topic.Names").
		Where("degree_program_id = ?", degreeProgramID).
		Order("created_at DESC").
		Find(&reports).Error; err != nil {
		return nil, err
	}

	var reportEntities []entity.Report
	for _, reportModel := range reports {
		reportEntity, err := mapper.Map[model.Report, entity.Report](&reportModel)
		if err != nil {
			return nil, err
		}
		reportEntities = append(reportEntities, reportEntity)
	}

	return reportEntities, nil
}

func (r *ReportRepository) GetByID(reportID uint, ctx _db.Context) (entity.Report, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.Report{}, err
	}

	var report model.Report
	if err := dbCtx.DB().
		Model(&model.Report{}).
		Preload("DegreeProgram").
		Preload("DegreeProgram.DurationUnit").
		Preload("DegreeProgram.DurationUnit.Names").
		Preload("DegreeProgram.UserCreator").
		Preload("DegreeProgram.UserCreator.Person").
		Preload("DegreeProgram.UserCreator.Person.HigherEducationInstitution").
		Preload("ProfessionalRole").
		Preload("KnowledgeAreaReports").
		Preload("KnowledgeAreaReports.TopicReports").
		Preload("KnowledgeAreaReports.TopicReports.Topic").
		Preload("KnowledgeAreaReports.TopicReports.Topic.Names").
		Where("reports.id = ?", reportID).
		First(&report).Error; err != nil {
		return entity.Report{}, err
	}

	return mapper.Map[model.Report, entity.Report](&report)
}
