package persistence

import "websac3/app/domain/entity"

type UpdateEmailPort interface {
	Update(email *entity.EmailNotification, db Context) error
}
