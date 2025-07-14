package repository

import (
	"errors"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/common/mapper"

	"gorm.io/gorm"
)

type StatusRepository struct {
	Repository
}

func NewStatusRepository(messageProvider message.Provider) *StatusRepository {
	return &StatusRepository{}
}

func (a *StatusRepository) GetByName(name string, ctx persistence.Context) (entity.Status, error) {
	dbCtx, err := a.CastDbContext(ctx)
	if err != nil {
		return entity.Status{}, err
	}

	var status entity.Status
	var statusFound model.Status
	if err = dbCtx.DB().
		Where("name = ?", name).
		First(&statusFound).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Status{}, nil
		}
		return entity.Status{}, err
	}

	if status, err = mapper.Map[model.Status, entity.Status](&statusFound); err != nil {
		return entity.Status{}, err
	}

	return status, nil
}
