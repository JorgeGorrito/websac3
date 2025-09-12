package entity

type KnowledgeAreaExpected struct {
	ID uint

	KnowledgeAreaID uint
	KnowledgeArea   *KnowledgeArea

	TopicExpected  []TopicExpected
	PriorityWeight float32
}

func (e *KnowledgeAreaExpected) IsRegistered() bool {
	return e.ID != 0
}

func (e *KnowledgeAreaExpected) GetKnowledgeAreaName() string {
	if e.KnowledgeArea == nil {
		return ""
	}
	return e.KnowledgeArea.Name
}

func (e *KnowledgeAreaExpected) GetTotalLearnHoursExpected() float32 {
	var totalLearnHoursExpected float32 = 0.0
	for _, topicExpected := range e.TopicExpected {
		totalLearnHoursExpected += float32(topicExpected.LearnHours)
	}
	return totalLearnHoursExpected
}
