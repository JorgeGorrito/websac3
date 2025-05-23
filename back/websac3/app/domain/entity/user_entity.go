package entity

import "time"

type User struct {
	ID            uint       `mapper:"userID"`
	PasswordHash  string     `mapper:"userPasswordHash"`
	Email         string     `mapper:"userEmail"`
	Role          *Role      `mapper:"role"`
	CreatedAt     time.Time  `mapper:"userCreatedAt"`
	UpdatedAt     *time.Time `mapper:"userUpdatedAt"`
	DeletedAt     *time.Time `mapper:"userDeleteAt"`
	DeactivatedAt *time.Time `mapper:"userDeletedAt"`
}

func (u *User) IsRegistered() bool {
	return u.ID != 0
}
