package entity

import (
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

type UnexpectedTopicReport struct {
	ID               uint
	TopicID          uint
	Topic            *Topic
	LearnHoursActual float32
	KnowledgeAreaID  uint
	KnowledgeArea    *KnowledgeArea
}

func NewUnexpectedTopicReport(
	topicID uint,
	learnHoursActual float32,
	knowledgeAreaID uint,
	knowledgeArea *KnowledgeArea,
	topic *Topic,
) *UnexpectedTopicReport {
	return &UnexpectedTopicReport{
		TopicID:          topicID,
		Topic:            topic,
		LearnHoursActual: learnHoursActual,
		KnowledgeAreaID:  knowledgeAreaID,
		KnowledgeArea:    knowledgeArea,
	}
}

type UnexpectedKnowledgeAreaReport struct {
	ID              uint
	Name            string
	Lang            string
	TotalLearnHours float32
	TopicReports    []UnexpectedTopicReport
}

func NewUnexpectedKnowledgeAreaReport(
	knowledgeAreaName string,
	lang string,
) *UnexpectedKnowledgeAreaReport {
	return &UnexpectedKnowledgeAreaReport{
		Name:            knowledgeAreaName,
		Lang:            lang,
		TotalLearnHours: 0.0,
		TopicReports:    []UnexpectedTopicReport{},
	}
}

func (e *UnexpectedKnowledgeAreaReport) AddTopicReport(topicReport *UnexpectedTopicReport) {
	e.TopicReports = append(e.TopicReports, *topicReport)
	e.TotalLearnHours += topicReport.LearnHoursActual
}

type Report struct {
	ID                             uint
	DegreeProgram                  DegreeProgram
	ProfessionalRole               ProfessionalRole
	KnowledgeAreaReports           []KnowledgeAreaReport
	UnexpectedKnowledgeAreaReports []UnexpectedKnowledgeAreaReport
	Score                          float32
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
		DegreeProgram:                  degreeProgram,
		ProfessionalRole:               professionalRole,
		KnowledgeAreaReports:           []KnowledgeAreaReport{},
		UnexpectedKnowledgeAreaReports: []UnexpectedKnowledgeAreaReport{},
		Score:                          0.0,
		HigherEducationInstitution:     higherEducationInstitution,
		CreatedAt:                      time.Now(),
	}
}

func (e *Report) AddKnowledgeAreaReport(knowledgeAreaReport *KnowledgeAreaReport) {
	e.KnowledgeAreaReports = append(e.KnowledgeAreaReports, *knowledgeAreaReport)
}

func (e *Report) AddUnexpectedKnowledgeAreaReport(knowledgeAreaReport *UnexpectedKnowledgeAreaReport) {
	e.UnexpectedKnowledgeAreaReports = append(e.UnexpectedKnowledgeAreaReports, *knowledgeAreaReport)
}

func (e *Report) findAndAddUnexpectedTopics(courseTopics []CourseTopic, knowledgeAreasExpected []KnowledgeAreaExpected, lang string) {
	// Create a map of expected topic IDs for quick lookup
	expectedTopicIDs := make(map[uint]bool)
	for _, knowledgeArea := range knowledgeAreasExpected {
		for _, topicExpected := range knowledgeArea.TopicExpected {
			expectedTopicIDs[topicExpected.TopicID] = true
		}
	}

	// Collect all unexpected topics in a single list
	var unexpectedTopics []CourseTopic
	for _, courseTopic := range courseTopics {
		// Check if this topic is not expected
		if !expectedTopicIDs[courseTopic.TopicID] {
			// Skip if topic information is not available
			if courseTopic.Topic == nil {
				continue
			}
			unexpectedTopics = append(unexpectedTopics, courseTopic)
		}
	}

	// Create a single unexpected knowledge area report with all topics
	if len(unexpectedTopics) > 0 {
		unexpectedAreaReport := NewUnexpectedKnowledgeAreaReport("Área de conocimiento adicional", lang)

		// Add all topics to the single unexpected knowledge area report
		for _, topic := range unexpectedTopics {
			knowledgeAreaID := topic.Topic.KnowledgeAreaID
			knowledgeArea := topic.Topic.KnowledgeArea
			if knowledgeArea == nil {
				knowledgeArea = &KnowledgeArea{ID: knowledgeAreaID, Name: "Área de conocimiento adicional"}
			}

			topicReport := NewUnexpectedTopicReport(
				topic.TopicID,
				topic.StudyHours,
				knowledgeAreaID,
				knowledgeArea,
				topic.Topic,
			)
			unexpectedAreaReport.AddTopicReport(topicReport)
		}

		e.AddUnexpectedKnowledgeAreaReport(unexpectedAreaReport)
	}
}

type ProfessionalRole struct {
	ID                    uint
	Name                  string
	KnowledgeAreaExpected []KnowledgeAreaExpected
}

func (e *ProfessionalRole) IsRegistered() bool {
	return e.ID != 0
}

func (e *ProfessionalRole) EvaluateDegreeProgram(degreeProgram *DegreeProgram, lang string) *Report {
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
			// Find all course topics that match this topic ID and sum their hours
			var actualHours float32 = 0.0
			var totalTopicHours float32 = 0.0

			for _, courseTopic := range courseTopics {
				if topicExpected.TopicID == courseTopic.TopicID {
					totalTopicHours += courseTopic.StudyHours
				}
			}

			if totalTopicHours > 0 {
				// Topic found in degree program (possibly in multiple courses)
				actualHours = totalTopicHours

				if totalTopicHours > topicExpected.LearnHours {
					partialScore += float32(topicExpected.LearnHours)
				} else {
					partialScore += float32(totalTopicHours)
				}
				knowledgeAreaReport.TotalLearnHoursActual += totalTopicHours
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

	// Find unexpected topics (topics covered in courses but not expected in the professional role)
	report.findAndAddUnexpectedTopics(courseTopics, knowledgeAreasExpected, lang)

	return report
}
