package repository

import (
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence"
	"websac3/common/mapper"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Create(user *entity.User, tx persistence.Transaction) error {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var userToSave models.User
	if err := mapper.Map(user, &userToSave); err != nil {
		return err
	}

	if err := pgTx.Tx().Create(&userToSave).Error; err != nil {
		return err
	}
	user.ID = userToSave.ID

	return nil
}

func (u *UserRepository) UpdateById(user *entity.User, userID uint, tx persistence.Transaction) error {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var userToUpdate models.User
	if err := mapper.Map(user, &userToUpdate); err != nil {
		return err
	}

	if err := pgTx.Tx().Model(&userToUpdate).Where("id = ?", userID).Updates(userToUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetByEmail(email string, tx persistence.Transaction) (entity.User, error) {
	pgTx, ok := tx.(*db.Transaction)
	if !ok {
		return entity.User{}, fmt.Errorf("expected *postgres.Transaction, got %T", tx)
	}

	var user models.User
	if err := pgTx.Tx().Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	var userEntity entity.User
	if err := mapper.Map(user, &userEntity); err != nil {
		return entity.User{}, err
	}

	return userEntity, nil
}
