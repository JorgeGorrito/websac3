package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
	"websac3/common/paginator"
)

type ReportFeedbackRepository struct {
	Repository
}

func NewReportFeedbackRepository() *ReportFeedbackRepository {
	return &ReportFeedbackRepository{}
}

// Create implements CreateReportFeedbackPort
func (r *ReportFeedbackRepository) Create(feedback *entity.ReportFeedback, ctx _db.Context) error {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var feedbackModel model.ReportFeedback
	if feedbackModel, err = mapper.Map[entity.ReportFeedback, model.ReportFeedback](feedback); err != nil {
		return err
	}

	// Crear el feedback principal
	if err := dbCtx.DB().
		Create(&feedbackModel).
		Error; err != nil {
		return err
	}

	// Asignar el ID generado
	feedback.ID = feedbackModel.ID

	// Crear los knowledge area feedbacks
	for i := range feedback.KnowledgeAreaFeedbacks {
		feedback.KnowledgeAreaFeedbacks[i].ReportFeedbackID = feedback.ID

		var kaFeedbackModel model.KnowledgeAreaFeedback
		if kaFeedbackModel, err = mapper.Map[entity.KnowledgeAreaFeedback, model.KnowledgeAreaFeedback](&feedback.KnowledgeAreaFeedbacks[i]); err != nil {
			return err
		}

		if err := dbCtx.DB().
			Create(&kaFeedbackModel).
			Error; err != nil {
			return err
		}

		feedback.KnowledgeAreaFeedbacks[i].ID = kaFeedbackModel.ID
	}

	return nil
}

// GetByID implements GetReportFeedbackPort
func (r *ReportFeedbackRepository) GetByID(feedbackID uint, ctx _db.Context) (*entity.ReportFeedback, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, err
	}

	var feedbackModel model.ReportFeedback
	if err := dbCtx.DB().
		Preload("Report").
		Preload("Report.DegreeProgram").
		Preload("Auditor").
		Preload("Auditor.Person").
		Preload("KnowledgeAreaFeedbacks").
		Preload("KnowledgeAreaFeedbacks.KnowledgeAreaReport").
		First(&feedbackModel, feedbackID).
		Error; err != nil {
		return nil, err
	}

	var feedback entity.ReportFeedback
	if feedback, err = mapper.Map[model.ReportFeedback, entity.ReportFeedback](&feedbackModel); err != nil {
		return nil, err
	}

	return &feedback, nil
}

// GetByReportID implements GetReportFeedbackPort
func (r *ReportFeedbackRepository) GetByReportID(reportID uint, ctx _db.Context) (*entity.ReportFeedback, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, err
	}

	var feedbackModel model.ReportFeedback
	if err := dbCtx.DB().
		Preload("Report").
		Preload("Report.DegreeProgram").
		Preload("Auditor").
		Preload("Auditor.Person").
		Preload("KnowledgeAreaFeedbacks").
		Preload("KnowledgeAreaFeedbacks.KnowledgeAreaReport").
		Where("report_id = ?", reportID).
		First(&feedbackModel).
		Error; err != nil {
		return nil, err
	}

	var feedback entity.ReportFeedback
	if feedback, err = mapper.Map[model.ReportFeedback, entity.ReportFeedback](&feedbackModel); err != nil {
		return nil, err
	}

	return &feedback, nil
}

// Update implements UpdateReportFeedbackPort
func (r *ReportFeedbackRepository) Update(feedback *entity.ReportFeedback, ctx _db.Context) error {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var feedbackModel model.ReportFeedback
	if feedbackModel, err = mapper.Map[entity.ReportFeedback, model.ReportFeedback](feedback); err != nil {
		return err
	}

	// Actualizar el feedback principal
	if err := dbCtx.DB().
		Model(&feedbackModel).
		Where("id = ?", feedback.ID).
		Updates(feedbackModel).
		Error; err != nil {
		return err
	}

	// Eliminar knowledge area feedbacks existentes
	if err := dbCtx.DB().
		Where("report_feedback_id = ?", feedback.ID).
		Delete(&model.KnowledgeAreaFeedback{}).
		Error; err != nil {
		return err
	}

	// Crear los nuevos knowledge area feedbacks
	for i := range feedback.KnowledgeAreaFeedbacks {
		feedback.KnowledgeAreaFeedbacks[i].ReportFeedbackID = feedback.ID

		var kaFeedbackModel model.KnowledgeAreaFeedback
		if kaFeedbackModel, err = mapper.Map[entity.KnowledgeAreaFeedback, model.KnowledgeAreaFeedback](&feedback.KnowledgeAreaFeedbacks[i]); err != nil {
			return err
		}

		if err := dbCtx.DB().
			Create(&kaFeedbackModel).
			Error; err != nil {
			return err
		}

		feedback.KnowledgeAreaFeedbacks[i].ID = kaFeedbackModel.ID
	}

	return nil
}

// GetWithFilters implements ListReportFeedbacksPort
func (r *ReportFeedbackRepository) GetWithFilters(filters map[string]interface{}, paginationParams paginator.PaginationParams, ctx _db.Context) ([]entity.ReportFeedback, uint, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().Model(&model.ReportFeedback{}).
		Preload("Report").
		Preload("Report.DegreeProgram").
		Preload("Auditor").
		Preload("Auditor.Person").
		Preload("KnowledgeAreaFeedbacks").
		Preload("KnowledgeAreaFeedbacks.KnowledgeAreaReport")

	// Aplicar filtros
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	// Contar total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginación
	var feedbackModels []model.ReportFeedback
	if err := query.
		Offset(int((paginationParams.Currentpage - 1) * paginationParams.ItemsPerpage)).
		Limit(int(paginationParams.ItemsPerpage)).
		Find(&feedbackModels).
		Error; err != nil {
		return nil, 0, err
	}

	var feedbacks []entity.ReportFeedback
	for _, feedbackModel := range feedbackModels {
		var feedback entity.ReportFeedback
		if feedback, err = mapper.Map[model.ReportFeedback, entity.ReportFeedback](&feedbackModel); err != nil {
			return nil, 0, err
		}
		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, uint(total), nil
}

// GetReportsWithoutFeedback implements ListReportsPendingFeedbackPort
func (r *ReportFeedbackRepository) GetReportsWithoutFeedback(paginationParams paginator.PaginationParams, ctx _db.Context) ([]entity.Report, uint, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Subquery para obtener report IDs que ya tienen feedback
	subQuery := dbCtx.DB().
		Model(&model.ReportFeedback{}).
		Select("report_id")

	// Query principal para obtener reportes sin feedback
	query := dbCtx.DB().Model(&model.Report{}).
		Preload("DegreeProgram").
		Preload("ProfessionalRole").
		Where("id NOT IN (?)", subQuery)

	// Contar total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginación
	var reportModels []model.Report
	if err := query.
		Offset(int((paginationParams.Currentpage - 1) * paginationParams.ItemsPerpage)).
		Limit(int(paginationParams.ItemsPerpage)).
		Find(&reportModels).
		Error; err != nil {
		return nil, 0, err
	}

	var reports []entity.Report
	for _, reportModel := range reportModels {
		var report entity.Report
		if report, err = mapper.Map[model.Report, entity.Report](&reportModel); err != nil {
			return nil, 0, err
		}
		reports = append(reports, report)
	}

	return reports, uint(total), nil
}
