package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreateExpertConsultationPort interface {
	Create(expertConsultation *entity.ExpertConsultation, db db.Context) error
}
