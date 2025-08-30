package entity

import (
	"websac3/app/domain/constants"
)

type Role struct {
	ID          uint
	Name        string
	Permissions []Permission
}

func (r *Role) IsAdmin() bool {
	return r.Name == constants.Admin
}

func (r *Role) IsCybersecurityAuditor() bool {
	return r.Name == constants.CybersecurityAuditor
}

func (r *Role) IsProgramLead() bool {
	return r.Name == constants.ProgramLead
}

func (r *Role) IsGuess() bool {
	return r.Name == constants.Guess
}
