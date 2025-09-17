package entity

import "time"

type ReportFeedback struct {
	ID       uint
	ReportID uint
	Report   *Report

	// Auditor que proporciona la retroalimentación
	AuditorID uint
	Auditor   *User

	// Contenido de la retroalimentación
	GeneralComments string
	Recommendations string

	// Calificación general del auditor (1-5)
	AuditorRating float32

	// Retroalimentación específica por área de conocimiento
	KnowledgeAreaFeedbacks []KnowledgeAreaFeedback

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewReportFeedback(
	reportID uint,
	auditorID uint,
	generalComments string,
	recommendations string,
	auditorRating float32,
) *ReportFeedback {
	return &ReportFeedback{
		ReportID:        reportID,
		AuditorID:       auditorID,
		GeneralComments: generalComments,
		Recommendations: recommendations,
		AuditorRating:   auditorRating,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (rf *ReportFeedback) AddKnowledgeAreaFeedback(knowledgeAreaFeedback *KnowledgeAreaFeedback) {
	rf.KnowledgeAreaFeedbacks = append(rf.KnowledgeAreaFeedbacks, *knowledgeAreaFeedback)
}

type KnowledgeAreaFeedback struct {
	ID               uint
	ReportFeedbackID uint
	ReportFeedback   *ReportFeedback

	KnowledgeAreaReportID uint
	KnowledgeAreaReport   *KnowledgeAreaReport

	// Comentarios específicos del área
	Comments        string
	SecurityGaps    string
	Improvements    string
	ComplianceLevel string

	// Calificación del auditor para esta área (1-5)
	AuditorRating float32
}

func NewKnowledgeAreaFeedback(
	reportFeedbackID uint,
	knowledgeAreaReportID uint,
	comments string,
	securityGaps string,
	improvements string,
	complianceLevel string,
	auditorRating float32,
) *KnowledgeAreaFeedback {
	return &KnowledgeAreaFeedback{
		ReportFeedbackID:      reportFeedbackID,
		KnowledgeAreaReportID: knowledgeAreaReportID,
		Comments:              comments,
		SecurityGaps:          securityGaps,
		Improvements:          improvements,
		ComplianceLevel:       complianceLevel,
		AuditorRating:         auditorRating,
	}
}
