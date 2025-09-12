package entity

import (
	"slices"
	"time"
)

type TopicReport struct {
	ID                 uint
	TopicID            uint
	Name               string
	LearnHoursExpected float32
	LearnHoursActual   float32
}

func NewTopicReport(
	topicID uint,
	name string,
	learnHoursExpected float32,
	learnHoursActual float32,
) *TopicReport {
	return &TopicReport{
		TopicID:            topicID,
		Name:               name,
		LearnHoursExpected: learnHoursExpected,
		LearnHoursActual:   learnHoursActual,
	}
}

type KnowledgeAreaReport struct {
	ID   uint
	Name string

	TotalLearnHoursExpected float32
	TotalLearnHoursActual   float32

	ScoreExpected float32
	ScoreGot      float32

	TopicReports []TopicReport
}

func NewKnowledgeAreaReport(
	knowledgeAreaName string,
	totalLearnHoursExpected float32,
) *KnowledgeAreaReport {
	return &KnowledgeAreaReport{
		Name:                    knowledgeAreaName,
		TotalLearnHoursExpected: totalLearnHoursExpected,
		TotalLearnHoursActual:   0.0,
		ScoreExpected:           0.0,
		ScoreGot:                0.0,
		TopicReports:            []TopicReport{},
	}
}

func (e *KnowledgeAreaReport) AddTopicReport(topicReport *TopicReport) {
	e.TopicReports = append(e.TopicReports, *topicReport)
}

type Report struct {
	ID                   uint
	DegreeProgram        DegreeProgram
	ProfessionalRole     ProfessionalRole
	KnowledgeAreaReports []KnowledgeAreaReport
	Score                float32
	// Información de la institución de educación superior
	HigherEducationInstitution *HigherEducationInstitution
	// Fecha de creación
	CreatedAt time.Time
}

func NewReport(
	degreeProgram DegreeProgram,
	professionalRole ProfessionalRole,
) *Report {
	var higherEducationInstitution *HigherEducationInstitution
	if degreeProgram.UserCreator != nil &&
		degreeProgram.UserCreator.Person != nil &&
		degreeProgram.UserCreator.Person.HigherEducationInstitution != nil {
		higherEducationInstitution = degreeProgram.UserCreator.Person.HigherEducationInstitution
	}

	return &Report{
		DegreeProgram:              degreeProgram,
		ProfessionalRole:           professionalRole,
		KnowledgeAreaReports:       []KnowledgeAreaReport{},
		Score:                      0.0,
		HigherEducationInstitution: higherEducationInstitution,
		CreatedAt:                  time.Now(),
	}
}

func (e *Report) AddKnowledgeAreaReport(knowledgeAreaReport *KnowledgeAreaReport) {
	e.KnowledgeAreaReports = append(e.KnowledgeAreaReports, *knowledgeAreaReport)
}

type ProfessionalRole struct {
	ID                    uint
	Name                  string
	KnowledgeAreaExpected []KnowledgeAreaExpected
}

func (e *ProfessionalRole) IsRegistered() bool {
	return e.ID != 0
}

func (e *ProfessionalRole) EvaluateDegreeProgram(degreeProgram *DegreeProgram) *Report {
	report := NewReport(
		*degreeProgram,
		*e,
	)

	courseTopics := degreeProgram.GetCourseTopics()
	knowledgeAreasExpected := e.KnowledgeAreaExpected

	var score float32 = 0.0
	var partialScore float32 = 0.0
	for _, knowledgeAreaExpected := range knowledgeAreasExpected {
		partialScore = 0.0
		totalLearnHoursExpected := knowledgeAreaExpected.GetTotalLearnHoursExpected()
		knowledgeAreaReport := NewKnowledgeAreaReport(
			knowledgeAreaExpected.GetKnowledgeAreaName(),
			totalLearnHoursExpected,
		)

		// Always include all expected topics for this knowledge area
		for _, topicExpected := range knowledgeAreaExpected.TopicExpected {
			iTopicFound := slices.IndexFunc(courseTopics, func(courseTopic CourseTopic) bool {
				return topicExpected.TopicID == courseTopic.TopicID
			})

			var actualHours float32 = 0.0
			if iTopicFound != -1 {
				// Topic found in degree program
				topicCourse := courseTopics[iTopicFound]
				actualHours = topicCourse.StudyHours

				if topicCourse.StudyHours > topicExpected.LearnHours {
					partialScore += float32(topicExpected.LearnHours)
				} else {
					partialScore += float32(topicCourse.StudyHours)
				}
				knowledgeAreaReport.TotalLearnHoursActual += topicCourse.StudyHours
			}
			// If not found, actualHours remains 0.0

			topicReport := NewTopicReport(
				topicExpected.TopicID,
				topicExpected.GetTopicName(),
				topicExpected.LearnHours,
				actualHours,
			)
			knowledgeAreaReport.AddTopicReport(topicReport)
		}

		knowledgeAreaReport.ScoreExpected = knowledgeAreaExpected.PriorityWeight

		partialScore = (partialScore / totalLearnHoursExpected) * float32(knowledgeAreaExpected.PriorityWeight)
		knowledgeAreaReport.ScoreGot = partialScore
		report.AddKnowledgeAreaReport(knowledgeAreaReport)

		score += partialScore
	}
	report.Score = score

	return report
}
