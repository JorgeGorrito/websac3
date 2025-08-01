package mapper

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerRoleMappers() {
	RegisterMapFunc(func(role *model.Role) (entity.Role, error) {
		var err error

		var permissions []entity.Permission
		for _, permission := range role.Permissions {
			var permissionMapped entity.Permission
			permissionMapped, err = Map[model.Permission, entity.Permission](&permission)
			if err != nil {
				return entity.Role{}, err
			}
			permissions = append(permissions, permissionMapped)
		}

		return entity.Role{
			ID:          role.ID,
			Name:        role.Name,
			Permissions: permissions,
		}, nil
	})

	RegisterMapFunc(func(role *entity.Role) (model.Role, error) {
		return model.Role{
			ID:   role.ID,
			Name: role.Name,
		}, nil
	})
}
