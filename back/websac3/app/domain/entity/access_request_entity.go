package entity

import (
	"time"
	"websac3/app/domain/errs"
)

type AccessRequest struct {
	ID        uint      `mapper:"accessRequestID"`
	Person    *Person   `mapper:"person"`
	Status    *Status   `mapper:"status"`
	CreatedAt time.Time `mapper:"accessRequestcreatedAt"`
	UpdatedAt time.Time `mapper:"accessRequestupdatedAt"`
	DeleteAt  time.Time `mapper:"accessRequestdeleteAt"`
}

func (a *AccessRequest) IsRegistered() bool {
	return a.ID != 0
}

func (a *AccessRequest) CanRegisterForFirstTime() bool {
	return !a.IsRegistered()
}

func (a *AccessRequest) CanRegisterAnother() (bool, error) {
	if isFirstTime := a.CanRegisterForFirstTime(); !isFirstTime {
		switch {
		case a.Status.IsRejected():
			return true, nil
		case a.Status.IsApproved():
			return false, errs.NewConflictError("access request already approved")
		case a.Status.IsPending():
			return false, errs.NewConflictError("access request already pending")
		}
	}
	return false, nil
}
