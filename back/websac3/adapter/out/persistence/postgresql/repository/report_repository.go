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
		Preload("DegreeProgram.UserCreator.Person.HigherEducationInstitution").
		Preload("ProfessionalRole").
		Where("id = ?", reportID).
		First(&report).Error; err != nil {
		return entity.Report{}, err
	}

	// Mapeo manual simplificado para evitar errores
	entityReport := entity.Report{
		ID:        report.ID,
		Score:     report.Score,
		CreatedAt: report.CreatedAt,
		DegreeProgram: entity.DegreeProgram{
			ID:   report.DegreeProgram.ID,
			Name: report.DegreeProgram.Name,
			UserCreator: &entity.User{
				ID:    report.DegreeProgram.UserCreator.ID,
				Email: report.DegreeProgram.UserCreator.Email,
				Person: &entity.Person{
					ID:       report.DegreeProgram.UserCreator.Person.ID,
					Name:     report.DegreeProgram.UserCreator.Person.Name,
					Lastname: report.DegreeProgram.UserCreator.Person.Lastname,
					HigherEducationInstitution: &entity.HigherEducationInstitution{
						Snies: report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Snies,
						Name:  report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Name,
					},
				},
			},
		},
		ProfessionalRole: entity.ProfessionalRole{
			ID:   report.ProfessionalRole.ID,
			Name: report.ProfessionalRole.Name,
		},
		KnowledgeAreaReports: []entity.KnowledgeAreaReport{}, // Vacío por ahora
	}

	// Asignar HigherEducationInstitution desde UserCreator.Person
	if entityReport.DegreeProgram.UserCreator != nil &&
		entityReport.DegreeProgram.UserCreator.Person != nil &&
		entityReport.DegreeProgram.UserCreator.Person.HigherEducationInstitution != nil {
		entityReport.HigherEducationInstitution = entityReport.DegreeProgram.UserCreator.Person.HigherEducationInstitution
	}

	return entityReport, nil
}
