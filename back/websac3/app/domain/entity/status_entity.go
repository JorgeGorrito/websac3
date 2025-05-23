package entity

type Status struct {
	ID   uint   `mapper:"statusID"`
	Name string `mapper:"statusName"`
}

func (s *Status) IsRegistered() bool {
	return s.ID != 0
}

func (s *Status) IsPending() bool {
	return s.Name == "pending"
}

func (s *Status) IsApproved() bool {
	return s.Name == "approved"
}

func (s *Status) IsRejected() bool {
	return s.Name == "rejected"
}
