package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerIdentificationTypeMappers() {
	RegisterMapFunc(func(identificationType *model.IdentificationType) (entity.IdentificationType, error) {
		return entity.IdentificationType{
			ID:   identificationType.ID,
			Name: identificationType.Name,
		}, nil
	})

	RegisterMapFunc(func(identificationType *entity.IdentificationType) (response.ListIdentificationTypeResponse, error) {
		return response.ListIdentificationTypeResponse{
			ID:   identificationType.ID,
			Name: identificationType.Name,
		}, nil
	})
}
