package persistence

import "websac3/app/domain/entity"

type UpdateAccessRequestPort interface {
	Update(request *entity.AccessRequest, ctx Context) error
}
