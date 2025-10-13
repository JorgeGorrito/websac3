package seeders

import (
	"fmt"
	"strconv"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_ROLE_SEED              = "adapter/out/persistence/postgresql/seeders/seeds/roles.json"
	DEFAULT_PATH_ROLES_PERMISSIONS_SEED = "adapter/out/persistence/postgresql/seeders/seeds/roles_permissions.json"
)

type roles struct{}

func Roles() Seeder {
	return &roles{}
}

func (r *roles) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Role = make([]model.Role, 0)
	var rolesPermission map[string][]uint = make(map[string][]uint)

	if err := decoder.Decode(DEFAULT_PATH_ROLE_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, role := range dataToSeed {
		if err := dbCtx.DB().Create(&role).Error; err != nil {
			return err
		}
	}

	if err := decoder.Decode(DEFAULT_PATH_ROLES_PERMISSIONS_SEED, &rolesPermission); err != nil {
		return err
	}
	for roleID, permissionIDs := range rolesPermission {
		roleIDUint, err := func() (uint, error) {
			roleIDUint64, err := strconv.ParseUint(roleID, 32, 0)
			return uint(roleIDUint64), err
		}()
		if err != nil {
			return fmt.Errorf("invalid role ID %s: %w", roleID, err)
		}

		var role model.Role
		if err := dbCtx.DB().First(&role, roleIDUint).Error; err != nil {
			return fmt.Errorf("failed to find role with ID %d: %w", roleIDUint, err)
		}

		for _, permissionID := range permissionIDs {
			err := dbCtx.DB().Model(&role).Association("Permissions").
				Append(&model.Permission{ID: permissionID})
			if err != nil {
				return fmt.Errorf("failed to associate permission %d with role %d: %w", permissionID, roleIDUint, err)
			}
		}
	}

	ResetAutoIncrement(dbCtx, "roles")

	return nil
}
