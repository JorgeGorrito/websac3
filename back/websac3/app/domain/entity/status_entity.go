package entity

import "websac3/app/domain/constants"

type Status struct {
	ID   uint
	Name string
}

func (s *Status) IsRegistered() bool {
	return s.ID != 0
}

func (s *Status) IsPending() bool {
	return s.Name == constants.Pending
}

func (s *Status) IsApproved() bool {
	return s.Name == constants.Approved
}

func (s *Status) IsRejected() bool {
	return s.Name == constants.Rejected
}
