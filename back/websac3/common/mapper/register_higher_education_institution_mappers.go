package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerHigherEducationInstitutionMappers() {
	RegisterMapFunc(func(category *model.InstitutionalCategory) (entity.InstitutionalCategory, error) {
		return entity.InstitutionalCategory{
			ID:   category.ID,
			Name: category.Name,
		}, nil
	})

	RegisterMapFunc(func(ownership *model.Ownership) (entity.Ownership, error) {
		return entity.Ownership{
			ID:   ownership.ID,
			Name: ownership.Name,
		}, nil
	})

	RegisterMapFunc(func(municipality *model.Municipality) (entity.Municipality, error) {
		return entity.Municipality{
			ID:   municipality.ID,
			Name: municipality.Name,
		}, nil
	})

	RegisterMapFunc(func(department *model.Department) (entity.Department, error) {
		return entity.Department{
			ID:   department.ID,
			Name: department.Name,
		}, nil
	})

	RegisterMapFunc(func(higherEducationInstitution *model.HigherEducationInstitution) (entity.HigherEducationInstitution, error) {
		ownership, err := Map[model.Ownership, entity.Ownership](&higherEducationInstitution.Ownership)
		if err != nil {
			return entity.HigherEducationInstitution{}, err
		}

		category, err := Map[model.InstitutionalCategory, entity.InstitutionalCategory](&higherEducationInstitution.InstitutionalCategory)
		if err != nil {
			return entity.HigherEducationInstitution{}, err
		}

		municipality, err := Map[model.Municipality, entity.Municipality](&higherEducationInstitution.Municipality)
		if err != nil {
			return entity.HigherEducationInstitution{}, err
		}

		department, err := Map[model.Department, entity.Department](&higherEducationInstitution.Department)
		if err != nil {
			return entity.HigherEducationInstitution{}, err
		}

		return entity.HigherEducationInstitution{
			Snies:                   higherEducationInstitution.Snies,
			Name:                    higherEducationInstitution.Name,
			OwnershipID:             higherEducationInstitution.OwnershipID,
			Ownership:               &ownership,
			InstitutionalCategoryID: higherEducationInstitution.InstitutionalCategoryID,
			InstitutionalCategory:   &category,
			MunicipalityID:          higherEducationInstitution.MunicipalityID,
			Municipality:            &municipality,
			DepartmentID:            higherEducationInstitution.DepartmentID,
			Department:              &department,
		}, nil
	})

	RegisterMapFunc(func(higherEducationInstitution *entity.HigherEducationInstitution) (response.ListHigherEducationInstitutionResponse, error) {
		var departmentName string
		var municipalityName string

		if higherEducationInstitution.Department != nil {
			departmentName = higherEducationInstitution.Department.Name
		}

		if higherEducationInstitution.Municipality != nil {
			municipalityName = higherEducationInstitution.Municipality.Name
		}

		return response.ListHigherEducationInstitutionResponse{
			Snies:        higherEducationInstitution.Snies,
			Name:         higherEducationInstitution.Name,
			Department:   departmentName,
			Municipality: municipalityName,
		}, nil
	})
}
