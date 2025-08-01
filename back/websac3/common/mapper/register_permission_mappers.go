package mapper

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerPermissionMappers() {
	RegisterMapFunc(func(permission *model.Permission) (entity.Permission, error) {
		var err error

		var module entity.Module
		var action entity.Action
		module, err = Map[model.Module, entity.Module](&permission.Module)
		if err != nil {
			return entity.Permission{}, err
		}

		action, err = Map[model.Action, entity.Action](&permission.Action)
		if err != nil {
			return entity.Permission{}, err
		}

		return entity.Permission{
			ID:       permission.ID,
			ModuleID: permission.ModuleID,
			Module:   &module,
			ActionID: permission.ActionID,
			Action:   &action,
		}, nil
	})
}
