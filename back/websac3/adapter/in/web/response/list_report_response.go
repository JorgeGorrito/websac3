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
	KnowledgeAreaReports           []KnowledgeAreaReportResponse           `json:"knowledge_area_reports"`
	UnexpectedKnowledgeAreaReports []UnexpectedKnowledgeAreaReportResponse `json:"unexpected_knowledge_area_reports"`
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

type UnexpectedKnowledgeAreaReportResponse struct {
	ID              uint                            `json:"id"`
	Name            string                          `json:"name"`
	Lang            string                          `json:"lang"`
	TotalLearnHours float32                         `json:"total_learn_hours"`
	TopicReports    []UnexpectedTopicReportResponse `json:"topic_reports"`
}

type UnexpectedTopicReportResponse struct {
	ID               uint    `json:"id"`
	TopicID          uint    `json:"topic_id"`
	Name             string  `json:"name"`
	LearnHoursActual float32 `json:"learn_hours_actual"`
	KnowledgeAreaID  uint    `json:"knowledge_area_id"`
}
