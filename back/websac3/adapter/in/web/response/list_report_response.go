package response

import "time"

type ListReportResponse struct {
	ID               uint                      `json:"id"`
	DegreeProgramID  uint                      `json:"degree_program_id"`
	DegreeProgram    ListDegreeProgramResponse `json:"degree_program"`
	ProfessionalRole ProfessionalRoleResponse  `json:"professional_role"`
	Score            float32                   `json:"score"`
	CreatedAt        time.Time                 `json:"created_at"`
	// Knowledge area reports summary
	KnowledgeAreaReports []KnowledgeAreaReportResponse `json:"knowledge_area_reports"`
}

type ProfessionalRoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type KnowledgeAreaReportResponse struct {
	ID                      uint                  `json:"id"`
	Name                    string                `json:"name"`
	TotalLearnHoursExpected float32               `json:"total_learn_hours_expected"`
	TotalLearnHoursActual   float32               `json:"total_learn_hours_actual"`
	ScoreExpected           float32               `json:"score_expected"`
	ScoreGot                float32               `json:"score_got"`
	TopicReports            []TopicReportResponse `json:"topic_reports"`
}

type TopicReportResponse struct {
	ID                 uint    `json:"id"`
	TopicID            uint    `json:"topic_id"`
	Name               string  `json:"name"`
	LearnHoursExpected float32 `json:"learn_hours_expected"`
	LearnHoursActual   float32 `json:"learn_hours_actual"`
}
