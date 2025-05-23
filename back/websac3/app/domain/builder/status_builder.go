package builder

import "websac3/app/domain/entity"

type StatusBuilder interface {
	WithName(name string) StatusBuilder
	WithID(id uint) StatusBuilder
	Build() *entity.Status
}

type statusBuilder struct {
	status entity.Status
}

func NewStatusBuilder() StatusBuilder {
	return &statusBuilder{
		status: entity.Status{},
	}
}

func (b *statusBuilder) WithName(name string) StatusBuilder {
	b.status.Name = name
	return b
}

func (b *statusBuilder) WithID(id uint) StatusBuilder {
	b.status.ID = id
	return b
}

func (b *statusBuilder) Build() *entity.Status {
	return &b.status
}
