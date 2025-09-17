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

type UnexpectedTopicReport struct {
	ID               uint
	TopicID          uint
	Name             string
	LearnHoursActual float32
	KnowledgeAreaID  uint
	KnowledgeArea    *KnowledgeArea
}

func NewUnexpectedTopicReport(
	topicID uint,
	name string,
	learnHoursActual float32,
	knowledgeAreaID uint,
	knowledgeArea *KnowledgeArea,
) *UnexpectedTopicReport {
	return &UnexpectedTopicReport{
		TopicID:          topicID,
		Name:             name,
		LearnHoursActual: learnHoursActual,
		KnowledgeAreaID:  knowledgeAreaID,
		KnowledgeArea:    knowledgeArea,
	}
}

type UnexpectedKnowledgeAreaReport struct {
	ID              uint
	Name            string
	TotalLearnHours float32
	TopicReports    []UnexpectedTopicReport
}

func NewUnexpectedKnowledgeAreaReport(
	knowledgeAreaName string,
) *UnexpectedKnowledgeAreaReport {
	return &UnexpectedKnowledgeAreaReport{
		Name:            knowledgeAreaName,
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

func (e *Report) findAndAddUnexpectedTopics(courseTopics []CourseTopic, knowledgeAreasExpected []KnowledgeAreaExpected) {
	// Create a map of expected topic IDs for quick lookup
	expectedTopicIDs := make(map[uint]bool)
	for _, knowledgeArea := range knowledgeAreasExpected {
		for _, topicExpected := range knowledgeArea.TopicExpected {
			expectedTopicIDs[topicExpected.TopicID] = true
		}
	}

	// Group unexpected topics by knowledge area
	unexpectedTopicsByArea := make(map[uint][]CourseTopic)
	for _, courseTopic := range courseTopics {
		// Check if this topic is not expected
		if !expectedTopicIDs[courseTopic.TopicID] {
			// Skip if topic information is not available
			if courseTopic.Topic == nil {
				continue
			}

			knowledgeAreaID := courseTopic.Topic.KnowledgeAreaID
			unexpectedTopicsByArea[knowledgeAreaID] = append(unexpectedTopicsByArea[knowledgeAreaID], courseTopic)
		}
	}

	// Create unexpected knowledge area reports
	for knowledgeAreaID, topics := range unexpectedTopicsByArea {
		// Get knowledge area name from the first topic (all topics in this area should have the same knowledge area)
		var knowledgeAreaName string
		var knowledgeArea *KnowledgeArea
		if len(topics) > 0 && topics[0].Topic != nil && topics[0].Topic.KnowledgeArea != nil {
			knowledgeAreaName = topics[0].Topic.KnowledgeArea.Name
			knowledgeArea = topics[0].Topic.KnowledgeArea
			// If the name is empty, provide a default
			if knowledgeAreaName == "" {
				knowledgeAreaName = "Área de conocimiento adicional"
			}
		} else {
			knowledgeAreaName = "Área de conocimiento adicional"
			knowledgeArea = &KnowledgeArea{ID: knowledgeAreaID, Name: knowledgeAreaName}
		}

		unexpectedAreaReport := NewUnexpectedKnowledgeAreaReport(knowledgeAreaName)

		// Add topics to the unexpected knowledge area report
		for _, topic := range topics {
			var topicName string
			if topic.Topic != nil && topic.Topic.Name != "" {
				topicName = topic.Topic.Name
			} else {
				topicName = "Temática adicional"
			}

			topicReport := NewUnexpectedTopicReport(
				topic.TopicID,
				topicName,
				topic.StudyHours,
				knowledgeAreaID,
				knowledgeArea,
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

	// Find unexpected topics (topics covered in courses but not expected in the professional role)
	report.findAndAddUnexpectedTopics(courseTopics, knowledgeAreasExpected)

	return report
}
