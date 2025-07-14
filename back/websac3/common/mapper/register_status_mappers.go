package mapper

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerStatusMappers() {
	RegisterMapFunc(func(status *model.Status) (entity.Status, error) {
		return entity.Status{
			ID:   status.ID,
			Name: status.Name,
		}, nil
	})

	RegisterMapFunc(func(status *entity.Status) (model.Status, error) {
		return model.Status{
			ID:   status.ID,
			Name: status.Name,
		}, nil
	})
}
