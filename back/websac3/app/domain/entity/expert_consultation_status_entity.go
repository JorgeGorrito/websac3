package entity

type ExpertConsultationStatus struct {
	ID    uint
	Name  string
	Names []ExpertConsultationStatusName
}

type ExpertConsultationStatusName struct {
	ID   uint
	Lang string
	Name string
}

func (ecs *ExpertConsultationStatus) IsRegistered() bool {
	return ecs.ID != 0
}
