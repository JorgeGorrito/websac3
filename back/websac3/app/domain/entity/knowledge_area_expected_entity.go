package entity

type KnowledgeAreaExpected struct {
	KnowledgeAreaID uint

	TopicExpected  []TopicExpected
	PriorityWeight float32
}
