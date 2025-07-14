package entity

import "time"

type User struct {
	ID            uint
	PasswordHash  string
	Email         string
	Role          *Role
	CreatedAt     time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
	DeactivatedAt *time.Time
}

func (u *User) IsRegistered() bool {
	return u.ID != 0
}
