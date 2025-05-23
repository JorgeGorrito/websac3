package builder

import "websac3/app/domain/entity"

type UserBuilder interface {
	WithID(id uint) UserBuilder
	WithEmail(email string) UserBuilder
	WithRole(role *entity.Role) UserBuilder

	Build() *entity.User
}

type userBuilder struct {
	user entity.User
}

func NewUserBuilder() UserBuilder {
	return &userBuilder{}
}

func (b *userBuilder) WithID(id uint) UserBuilder {
	b.user.ID = id
	return b
}

func (b *userBuilder) WithEmail(email string) UserBuilder {
	b.user.Email = email
	return b
}

func (b *userBuilder) WithRole(role *entity.Role) UserBuilder {
	b.user.Role = role
	return b
}

func (b *userBuilder) Build() *entity.User {
	return &b.user
}
