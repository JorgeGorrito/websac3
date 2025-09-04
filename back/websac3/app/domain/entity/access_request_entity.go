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

	ApprovedRoleID *uint
	ApprovedRole   *Role

	IsVerified bool
	CreatedAt  time.Time
}

func (a *AccessRequest) HasEmailVerified() bool {
	return a.IsVerified
}

func (a *AccessRequest) IsRegistered() bool {
	return a.ID != 0
}

func (a *AccessRequest) IsApproved() bool {
	return a.Status != nil && a.Status.IsApproved()
}

func (a *AccessRequest) IsRejected() bool {
	return a.Status != nil && a.Status.IsRejected()
}

func (a *AccessRequest) CanRegisterNewRequest(emailNewRequest string) bool {
	return !a.IsRegistered() || (a.IsRejected() && a.EmailValidation != nil && a.EmailValidation.To == emailNewRequest)
}

func (a *AccessRequest) IsApplicantRegistered() bool {
	return a.Applicant != nil && a.Applicant.IsRegistered()
}
