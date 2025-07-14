package mapper

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerHigherEducationInstitutionMappers() {
	RegisterMapFunc(func(higherEducationInstitution *model.HigherEducationInstitution) (entity.HigherEducationInstitution, error) {
		return entity.HigherEducationInstitution{
			Snies: higherEducationInstitution.Snies,
			Name:  higherEducationInstitution.Name,
		}, nil
	})
}
