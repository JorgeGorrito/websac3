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
) *ReportFeedback {
	return &ReportFeedback{
		ReportID:        reportID,
		AuditorID:       auditorID,
		GeneralComments: generalComments,
		Recommendations: recommendations,
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
	Comments string
}

func NewKnowledgeAreaFeedback(
	reportFeedbackID uint,
	knowledgeAreaReportID uint,
	comments string,
) *KnowledgeAreaFeedback {
	return &KnowledgeAreaFeedback{
		ReportFeedbackID:      reportFeedbackID,
		KnowledgeAreaReportID: knowledgeAreaReportID,
		Comments:              comments,
	}
}
