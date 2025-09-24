package entity

type ExpertConsultationStatus struct {
	ID   uint
	Name string
}

func (ecs *ExpertConsultationStatus) IsRegistered() bool {
	return ecs.ID != 0
}
