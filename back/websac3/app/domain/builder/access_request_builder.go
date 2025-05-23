package builder

import (
	"time"
	"websac3/app/domain/entity"
)

type AccessRequestBuilder interface {
	WithID(id uint) AccessRequestBuilder
	WithPerson(person *entity.Person) AccessRequestBuilder
	WithStatus(status *entity.Status) AccessRequestBuilder
	WithCreatedAt(createdAt time.Time) AccessRequestBuilder
	WithUpdatedAt(updatedAt time.Time) AccessRequestBuilder
	Build() *entity.AccessRequest
}

type accessRequestBuilder struct {
	accessRequest entity.AccessRequest
}

func NewAccessRequestBuilder() AccessRequestBuilder {
	return &accessRequestBuilder{}
}

func (b *accessRequestBuilder) WithID(id uint) AccessRequestBuilder {
	b.accessRequest.ID = id
	return b
}

func (b *accessRequestBuilder) WithPerson(person *entity.Person) AccessRequestBuilder {
	b.accessRequest.Person = person
	return b
}

func (b *accessRequestBuilder) WithStatus(status *entity.Status) AccessRequestBuilder {
	b.accessRequest.Status = status
	return b
}

func (b *accessRequestBuilder) WithCreatedAt(createdAt time.Time) AccessRequestBuilder {
	b.accessRequest.CreatedAt = createdAt
	return b
}

func (b *accessRequestBuilder) WithUpdatedAt(updatedAt time.Time) AccessRequestBuilder {
	b.accessRequest.UpdatedAt = updatedAt
	return b
}

func (b *accessRequestBuilder) Build() *entity.AccessRequest {
	return &b.accessRequest
}
