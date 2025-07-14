package persistence

import "websac3/app/domain/entity"

type CreateEmailPort interface {
	Create(emailToCreate *entity.EmailNotification, db Context) error
}
