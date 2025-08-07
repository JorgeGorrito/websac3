package repository

import (
	"errors"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/out/message"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(
	messageProvider message.Provider,
) *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Create(userToSave *entity.User, ctx _db.Context) error {
	dbCtx, err := u.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var user model.User
	if user, err = mapper.Map[entity.User, model.User](userToSave); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Create(&user).
		Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UpdateByID(userToUpdate *entity.User, userID uint, ctx _db.Context) error {
	dbCtx, err := u.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var user model.User
	if user, err = mapper.Map[entity.User, model.User](userToUpdate); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Model(&user).
		Where("id = ?", userID).
		Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetByEmail(email string, ctx _db.Context) (entity.User, error) {
	dbCtx, err := u.CastDbContext(ctx)
	if err != nil {
		return entity.User{}, err
	}

	var userFound model.User
	var user entity.User
	if err := dbCtx.DB().
		Preload("Role.Permissions.Action").
		Preload("Role.Permissions.Module").
		Joins("Role").
		Joins("Person").
		Where("email = ?", email).
		First(&userFound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	if user, err = mapper.Map[model.User, entity.User](&userFound); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserRepository) GetByDNI(identificationType uint, identificationNumber string, ctx _db.Context) (entity.User, error) {
	dbCtx, err := u.CastDbContext(ctx)
	if err != nil {
		return entity.User{}, err
	}

	var userFound model.User
	var user entity.User
	if err := dbCtx.DB().
		Model(&userFound).
		Joins("Person").
		Where(`"Person".identification_type_id = ?`, identificationType).
		Where(`"Person".identification_number = ?`, identificationNumber).
		First(&userFound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	if user, err = mapper.Map[model.User, entity.User](&userFound); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserRepository) GetByID(ID uint, ctx _db.Context) (entity.User, error) {
	dbCtx, err := u.CastDbContext(ctx)
	if err != nil {
		return entity.User{}, err
	}

	var userFound model.User
	var user entity.User
	if err := dbCtx.DB().
		Model(&userFound).
		Preload("Role.Permissions.Action").
		Preload("Role.Permissions.Module").
		Joins("Role").
		Joins("Person").
		Where("users.id = ?", ID).
		First(&userFound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	if user, err = mapper.Map[model.User, entity.User](&userFound); err != nil {
		return entity.User{}, err
	}

	return user, nil
}
