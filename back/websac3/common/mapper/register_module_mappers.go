package mapper

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerModuleMappers() {
	RegisterMapFunc(func(module *model.Module) (entity.Module, error) {
		return entity.Module{
			ID:   module.ID,
			Name: module.Name,
		}, nil
	})

	RegisterMapFunc(func(module *entity.Module) (model.Module, error) {
		return model.Module{
			ID:   module.ID,
			Name: module.Name,
		}, nil
	})
}
