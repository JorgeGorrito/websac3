package repository

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/out/message"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/filter"
	"websac3/common/mapper"
	"websac3/common/paginator"

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

	fmt.Printf("Updating user in DB: ID=%d, DeactivatedAt=%+v\n", user.ID, user.DeactivatedAt)

	// Usar Select para especificar qué campos actualizar, incluyendo deactivated_at
	if err := dbCtx.DB().
		Model(&model.User{}).
		Where("id = ?", userID).
		Select("email", "password_hash", "role_id", "person_id", "deactivated_at", "created_at").
		Updates(&user).Error; err != nil {
		fmt.Printf("Error updating user: %v\n", err)
		return err
	}

	fmt.Printf("User updated successfully in DB\n")
	return nil
}

func (u *UserRepository) GetByEmail(email string, ctx _db.Context) (entity.User, error) {
	dbCtx, err := u.CastDbContext(ctx)
	if err != nil {
		return entity.User{}, err
	}

	var userFound model.User
	if err := dbCtx.DB().
		Preload("Role.Permissions.Action").
		Preload("Role.Permissions.Module").
		Joins("Role").
		Joins("Person").
		Where("email = ?", email).
		First(&userFound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Retornar un usuario vacío sin error cuando no se encuentra
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	// Solo mapear si se encontró un usuario
	if userFound.ID == 0 {
		return entity.User{}, nil
	}

	user, err := mapper.Map[model.User, entity.User](&userFound)
	if err != nil {
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
	if err := dbCtx.DB().
		Model(&userFound).
		Joins("Person").
		Where(`"Person".identification_type_id = ?`, identificationType).
		Where(`"Person".identification_number = ?`, identificationNumber).
		First(&userFound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Retornar un usuario vacío sin error cuando no se encuentra
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	// Solo mapear si se encontró un usuario
	if userFound.ID == 0 {
		return entity.User{}, nil
	}

	user, err := mapper.Map[model.User, entity.User](&userFound)
	if err != nil {
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
	if err := dbCtx.DB().
		Model(&userFound).
		Preload("Role.Permissions.Action").
		Preload("Role.Permissions.Module").
		Joins("Role").
		Joins("Person").
		Where("users.id = ?", ID).
		First(&userFound).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Retornar un usuario vacío sin error cuando no se encuentra
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	// Solo mapear si se encontró un usuario
	if userFound.ID == 0 {
		return entity.User{}, nil
	}

	user, err := mapper.Map[model.User, entity.User](&userFound)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserRepository) ListAllWithPagination(ctx _db.Context, paginationParams paginator.PaginationParams, filters filter.Params) ([]entity.User, int64, error) {
	dbCtx, err := u.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().Model(&model.User{}).
		Preload("Role").
		Preload("Person").
		Preload("Person.HigherEducationInstitution").
		Preload("Person.IdentificationType")

	// Aplicar filtros
	for field, filterData := range filters {
		for operator, value := range filterData {
			if value == "" {
				continue
			}

			switch field {
			case "email":
				if operator == "cont" {
					query = query.Where("users.email ILIKE ?", "%"+value+"%")
				} else if operator == "eq" {
					query = query.Where("users.email = ?", value)
				}
			case "Person.name":
				if operator == "cont" {
					query = query.Joins("Person").Where("\"Person\".name ILIKE ?", "%"+value+"%")
				} else if operator == "eq" {
					query = query.Joins("Person").Where("\"Person\".name = ?", value)
				}
			case "Person.lastname":
				if operator == "cont" {
					query = query.Joins("Person").Where("\"Person\".lastname ILIKE ?", "%"+value+"%")
				} else if operator == "eq" {
					query = query.Joins("Person").Where("\"Person\".lastname = ?", value)
				}
			case "Role.name":
				if operator == "cont" {
					query = query.Joins("Role").Where("\"Role\".name ILIKE ?", "%"+value+"%")
				} else if operator == "eq" {
					query = query.Joins("Role").Where("\"Role\".name = ?", value)
				}
			case "is_active":
				if operator == "eq" {
					if value == "true" {
						query = query.Where("users.deactivated_at IS NULL")
					} else if value == "false" {
						query = query.Where("users.deactivated_at IS NOT NULL")
					}
				}
			}
		}
	}

	// Contar total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginación
	var users []model.User
	if err := query.
		Offset(int((paginationParams.Currentpage - 1) * paginationParams.ItemsPerpage)).
		Limit(int(paginationParams.ItemsPerpage)).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	var entityUsers []entity.User
	for _, user := range users {
		entityUser, err := mapper.Map[model.User, entity.User](&user)
		if err != nil {
			return nil, 0, err
		}
		entityUsers = append(entityUsers, entityUser)
	}

	return entityUsers, total, nil
}
