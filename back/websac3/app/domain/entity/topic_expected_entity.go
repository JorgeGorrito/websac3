package entity

type TopicExpected struct {
	ID uint

	TopicID uint
	Topic   *Topic

	LearnHours float32
}

func (e *TopicExpected) GetTopicName() string {
	if e.Topic == nil {
		return ""
	}
	return e.Topic.Name
}
