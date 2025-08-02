package entity

import (
	"time"
)

type AccessRequest struct {
	ID uint

	ValidationEmailURL string
	CreateUserURL      string

	ValidationEmailCode string
	CreateUserCode      string

	ApplicantID uint
	Applicant   *Person

	EmailValidationID *uint
	EmailValidation   *EmailNotification

	EmailApprovedID *uint
	EmailApproved   *EmailNotification

	StatusID uint
	Status   *Status

	IsVerified bool
	CreatedAt  time.Time
}

func (a *AccessRequest) IsRegistered() bool {
	return a.ID != 0
}

func (a *AccessRequest) IsApproved() bool {
	return a.Status != nil && a.Status.IsApproved()
}

func (a *AccessRequest) CanRegister() bool {
	return !a.IsRegistered() || !a.IsApproved()
}

func (a *AccessRequest) IsApplicantRegistered() bool {
	return a.Applicant != nil && a.Applicant.IsRegistered()
}
