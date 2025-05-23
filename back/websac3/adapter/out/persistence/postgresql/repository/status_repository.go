package repository

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/persistence"
	"websac3/common/mapper"

	"gorm.io/gorm"
)

type StatusRepository struct{}

func NewStatusRepository() *StatusRepository {
	return &StatusRepository{}
}

func (a *StatusRepository) GetByName(name string, tx persistence.Transaction) (status entity.Status, err error) {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return status, fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var statusFound models.Status
	if err = pgTx.Tx().Where("name = ?", name).First(&statusFound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status, errs.NewNotFoundError("status not found")
		}
		return status, err
	}

	if err = mapper.Map(&statusFound, &status); err != nil {
		return status, err
	}

	return status, nil
}
