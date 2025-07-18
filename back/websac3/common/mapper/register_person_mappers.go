package mapper

import (
	"errors"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerPersonMappers() {
	RegisterMapFunc(func(person *entity.Person) (model.Person, error) {
		return model.Person{
			ID:                              person.ID,
			Name:                            person.Name,
			Lastname:                        person.Lastname,
			IdentificationTypeID:            person.IdentificationTypeID,
			IdentificationNumber:            person.IdentificationNumber,
			HigherEducationInstitutionSnies: person.HigherEducationInstitutionSnies,
			JobPosition:                     person.JobPosition,
		}, nil
	})

	RegisterMapFunc(func(person *model.Person) (entity.Person, error) {
		var errorList error

		identificationType, err := Map[model.IdentificationType, entity.IdentificationType](&person.IdentificationType)
		if err != nil {
			errorList = errors.Join(errorList, err)
		}

		higherEducationInstitution, err := Map[model.HigherEducationInstitution, entity.HigherEducationInstitution](&person.HigherEducationInstitution)
		if err != nil {
			errorList = errors.Join(errorList, err)
		}

		if errorList != nil {
			return entity.Person{}, errorList
		}

		return entity.Person{
			ID:                              person.ID,
			Name:                            person.Name,
			Lastname:                        person.Lastname,
			IdentificationTypeID:            person.IdentificationTypeID,
			IdentificationType:              &identificationType,
			IdentificationNumber:            person.IdentificationNumber,
			HigherEducationInstitutionSnies: person.HigherEducationInstitutionSnies,
			HigherEducationInstitution:      &higherEducationInstitution,
			JobPosition:                     person.JobPosition,
		}, nil
	})
}
