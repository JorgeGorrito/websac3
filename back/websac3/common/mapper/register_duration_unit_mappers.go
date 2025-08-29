package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerDurationUnitMappers() {
	RegisterMapFunc(
		func(durationUnitModel *model.DurationUnit) (entity.DurationUnit, error) {
			return entity.DurationUnit{
				ID:   durationUnitModel.ID,
				Name: durationUnitModel.Names[0].Name,
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
