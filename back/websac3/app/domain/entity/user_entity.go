package entity

import "time"

type User struct {
	ID           uint
	Password     string
	PasswordHash string
	Email        string

	RoleID uint
	Role   *Role

	PersonID uint
	Person   *Person

	CreatedAt     time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
	DeactivatedAt *time.Time
}

func (u *User) IsRegistered() bool {
	return u.ID != 0
}

func (u *User) IsPasswordHashEqual(hash string) bool {
	return u.PasswordHash == hash
}

func (u *User) IsActive() bool {
	return u.DeactivatedAt == nil
}
