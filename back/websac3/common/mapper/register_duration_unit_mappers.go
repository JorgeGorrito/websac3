package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerDurationUnitMappers() {
	RegisterMapFunc(
		func(durationUnitModel *model.DurationUnit) (entity.DurationUnit, error) {
			var name string
			if len(durationUnitModel.Names) > 0 {
				name = durationUnitModel.Names[0].Name
			}
			return entity.DurationUnit{
				ID:   durationUnitModel.ID,
				Name: name,
			}, nil
		},
	)

	RegisterMapFunc(
		func(durationUnitEntity *entity.DurationUnit) (response.ListDurationUnitResponse, error) {
			return response.ListDurationUnitResponse{
				ID:   durationUnitEntity.ID,
				Name: durationUnitEntity.Name,
			}, nil
		},
	)
}
