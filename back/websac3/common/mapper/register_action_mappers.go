package mapper

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerActionMappers() {
	RegisterMapFunc(func(action *model.Action) (entity.Action, error) {
		return entity.Action{
			ID:   action.ID,
			Name: action.Name,
		}, nil
	})
	RegisterMapFunc(func(action *entity.Action) (model.Action, error) {
		return model.Action{
			ID:   action.ID,
			Name: action.Name,
		}, nil
	})
}
