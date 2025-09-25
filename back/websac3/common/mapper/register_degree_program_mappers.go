package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerDegreeProgramMappers() {
	RegisterMapFunc(
		func(degreeProgramModel *model.DegreeProgram) (entity.DegreeProgram, error) {
			var durationUnitPtr *entity.DurationUnit
			if degreeProgramModel.DurationUnit.ID != 0 {
				if du, err := Map[model.DurationUnit, entity.DurationUnit](&degreeProgramModel.DurationUnit); err == nil {
					durationUnitPtr = &du
				}
			}

			var formationLevelPtr *entity.FormationLevel
			if degreeProgramModel.FormationLevel.ID != 0 {
				if fl, err := Map[model.FormationLevel, entity.FormationLevel](&degreeProgramModel.FormationLevel); err == nil {
					formationLevelPtr = &fl
				}
			}

			var userCreatorPtr *entity.User
			if degreeProgramModel.UserCreator.ID != 0 {
				if uc, err := Map[model.User, entity.User](&degreeProgramModel.UserCreator); err == nil {
					userCreatorPtr = &uc
				}
			}

			var courses []entity.Course
			for _, cm := range degreeProgramModel.Courses {
				ce, err := Map[model.Course, entity.Course](&cm)
				if err != nil {
					return entity.DegreeProgram{}, err
				}
				courses = append(courses, ce)
			}

			return entity.DegreeProgram{
				ID:                  degreeProgramModel.ID,
				Snies:               degreeProgramModel.Snies,
				Name:                degreeProgramModel.Name,
				TotalCredits:        degreeProgramModel.TotalCredits,
				DurationValue:       degreeProgramModel.DurationValue,
				DurationUnitID:      degreeProgramModel.DurationUnitID,
				DurationUnit:        durationUnitPtr,
				FormationLevelID:    degreeProgramModel.FormationLevelID,
				FormationLevel:      formationLevelPtr,
				ProgramFocus:        degreeProgramModel.ProgramFocus,
				EntryProfile:        degreeProgramModel.EntryProfile,
				GraduateProfile:     degreeProgramModel.GraduateProfile,
				ProfessionalProfile: degreeProgramModel.ProfessionalProfile,
				Courses:             courses,
				CreatedBy:           degreeProgramModel.CreatedBy,
				UserCreator:         userCreatorPtr,
			}, nil
		},
	)

	RegisterMapFunc(
		func(degreeProgramEntity *entity.DegreeProgram) (model.DegreeProgram, error) {
			return model.DegreeProgram{
				ID:                  degreeProgramEntity.ID,
				Snies:               degreeProgramEntity.Snies,
				Name:                degreeProgramEntity.Name,
				TotalCredits:        degreeProgramEntity.TotalCredits,
				DurationValue:       degreeProgramEntity.DurationValue,
				DurationUnitID:      degreeProgramEntity.DurationUnitID,
				FormationLevelID:    degreeProgramEntity.FormationLevelID,
				ProgramFocus:        degreeProgramEntity.ProgramFocus,
				EntryProfile:        degreeProgramEntity.EntryProfile,
				GraduateProfile:     degreeProgramEntity.GraduateProfile,
				ProfessionalProfile: degreeProgramEntity.ProfessionalProfile,
				CreatedBy:           degreeProgramEntity.CreatedBy,
			}, nil
		},
	)

	RegisterMapFunc(
		func(request *request.CreateDegreeProgramRequest) (command.CreateDegreeProgramCommand, error) {
			return command.CreateDegreeProgramCommand{
				Snies:               request.Snies,
				Name:                request.Name,
				TotalCredits:        request.TotalCredits,
				DurationValue:       request.DurationValue,
				DurationUnitID:      request.DurationUnitID,
				FormationLevelID:    request.FormationLevelID,
				ProgramFocus:        request.ProgramFocus,
				EntryProfile:        request.EntryProfile,
				GraduateProfile:     request.GraduateProfile,
				ProfessionalProfile: request.ProfessionalProfile,
			}, nil
		},
	)

	RegisterMapFunc(
		func(command *command.CreateDegreeProgramCommand) (entity.DegreeProgram, error) {
			return entity.DegreeProgram{
				Snies:               command.Snies,
				Name:                command.Name,
				TotalCredits:        command.TotalCredits,
				DurationValue:       command.DurationValue,
				DurationUnitID:      command.DurationUnitID,
				FormationLevelID:    command.FormationLevelID,
				ProgramFocus:        command.ProgramFocus,
				EntryProfile:        command.EntryProfile,
				GraduateProfile:     command.GraduateProfile,
				ProfessionalProfile: command.ProfessionalProfile,
				CreatedBy:           command.CreatedBy,
			}, nil
		},
	)

	RegisterMapFunc(
		func(degreeProgramEntity *entity.DegreeProgram) (response.ListDegreeProgramResponse, error) {
			var durResp response.ListDurationUnitResponse
			if degreeProgramEntity.DurationUnit != nil {
				durResp = response.ListDurationUnitResponse{
					ID:   degreeProgramEntity.DurationUnit.ID,
					Name: degreeProgramEntity.DurationUnit.Name,
				}
			}

			var flResp response.ListFormationLevelsResponse
			if degreeProgramEntity.FormationLevel != nil {
				flResp = response.ListFormationLevelsResponse{
					ID:   degreeProgramEntity.FormationLevel.ID,
					Name: degreeProgramEntity.FormationLevel.Name,
				}
			}

			var heiResp *response.HigherEducationInstitutionInfo
			if degreeProgramEntity.UserCreator != nil &&
				degreeProgramEntity.UserCreator.Person != nil &&
				degreeProgramEntity.UserCreator.Person.HigherEducationInstitution != nil {
				hei := degreeProgramEntity.UserCreator.Person.HigherEducationInstitution
				heiResp = &response.HigherEducationInstitutionInfo{
					Snies:                 hei.Snies,
					Name:                  hei.Name,
					Ownership:             "",
					InstitutionalCategory: "",
					Municipality:          "",
					Department:            "",
				}

				// Agregar información adicional si está disponible
				if hei.Ownership != nil {
					heiResp.Ownership = hei.Ownership.Name
				}
				if hei.InstitutionalCategory != nil {
					heiResp.InstitutionalCategory = hei.InstitutionalCategory.Name
				}
				if hei.Municipality != nil {
					heiResp.Municipality = hei.Municipality.Name
				}
				if hei.Department != nil {
					heiResp.Department = hei.Department.Name
				}
			}

			return response.ListDegreeProgramResponse{
				ID:                         degreeProgramEntity.ID,
				Snies:                      degreeProgramEntity.Snies,
				Name:                       degreeProgramEntity.Name,
				TotalCredits:               degreeProgramEntity.TotalCredits,
				DurationValue:              degreeProgramEntity.DurationValue,
				DurationUnit:               durResp,
				FormationLevel:             flResp,
				ProgramFocus:               degreeProgramEntity.ProgramFocus,
				EntryProfile:               degreeProgramEntity.EntryProfile,
				GraduateProfile:            degreeProgramEntity.GraduateProfile,
				ProfessionalProfile:        degreeProgramEntity.ProfessionalProfile,
				CreatedBy:                  degreeProgramEntity.CreatedBy,
				HigherEducationInstitution: heiResp,
			}, nil
		},
	)

	// Evaluate Degree Program mappers
	RegisterMapFunc(
		func(req *request.EvaluateDegreeProgramRequest) (command.EvaluateDegreeProgramCommand, error) {
			return command.EvaluateDegreeProgramCommand{
				DegreeProgramID:    req.DegreeProgramID,
				ProfessionalRoleID: req.ProfessionalRoleID,
			}, nil
		},
	)
}

// GetDegreeProgramMapperWithLanguage returns a language-specific mapper for DegreeProgram
func GetDegreeProgramMapperWithLanguage(lang string) func(*model.DegreeProgram) (entity.DegreeProgram, error) {
	return func(degreeProgramModel *model.DegreeProgram) (entity.DegreeProgram, error) {
		// Map DurationUnit with language support
		var durationUnitPtr *entity.DurationUnit
		if degreeProgramModel.DurationUnit.ID != 0 {
			if du, err := Map[model.DurationUnit, entity.DurationUnit](&degreeProgramModel.DurationUnit); err == nil {
				durationUnitPtr = &du
			}
		}

		// Map UserCreator with Person and HigherEducationInstitution
		var userCreatorPtr *entity.User
		if degreeProgramModel.UserCreator.ID != 0 {
			var personPtr *entity.Person
			if degreeProgramModel.UserCreator.Person.ID != 0 {
				var higherEducationInstitutionPtr *entity.HigherEducationInstitution
				if degreeProgramModel.UserCreator.Person.HigherEducationInstitution.Snies != 0 {
					higherEducationInstitutionPtr = &entity.HigherEducationInstitution{
						Snies: degreeProgramModel.UserCreator.Person.HigherEducationInstitution.Snies,
						Name:  degreeProgramModel.UserCreator.Person.HigherEducationInstitution.Name,
					}
				}

				personPtr = &entity.Person{
					ID:                         degreeProgramModel.UserCreator.Person.ID,
					Name:                       degreeProgramModel.UserCreator.Person.Name,
					Lastname:                   degreeProgramModel.UserCreator.Person.Lastname,
					HigherEducationInstitution: higherEducationInstitutionPtr,
				}
			}

			userCreatorPtr = &entity.User{
				ID:     degreeProgramModel.UserCreator.ID,
				Email:  degreeProgramModel.UserCreator.Email,
				Person: personPtr,
			}
		}

		// Map Courses with language support
		var courses []entity.Course
		for _, course := range degreeProgramModel.Courses {
			// Use language-specific mapper for Course if available
			var courseEntity entity.Course
			if courseLangMapper := GetCourseMapperWithLanguage(lang); courseLangMapper != nil {
				var err error
				courseEntity, err = courseLangMapper(&course)
				if err != nil {
					return entity.DegreeProgram{}, err
				}
			} else {
				// Fallback to default mapper
				if mappedCourse, err := Map[model.Course, entity.Course](&course); err == nil {
					courseEntity = mappedCourse
				} else {
					return entity.DegreeProgram{}, err
				}
			}
			courses = append(courses, courseEntity)
		}

		return entity.DegreeProgram{
			ID:                  degreeProgramModel.ID,
			Snies:               degreeProgramModel.Snies,
			Name:                degreeProgramModel.Name,
			TotalCredits:        degreeProgramModel.TotalCredits,
			DurationValue:       degreeProgramModel.DurationValue,
			DurationUnitID:      degreeProgramModel.DurationUnitID,
			DurationUnit:        durationUnitPtr,
			ProgramFocus:        degreeProgramModel.ProgramFocus,
			EntryProfile:        degreeProgramModel.EntryProfile,
			GraduateProfile:     degreeProgramModel.GraduateProfile,
			ProfessionalProfile: degreeProgramModel.ProfessionalProfile,
			CreatedBy:           degreeProgramModel.CreatedBy,
			UserCreator:         userCreatorPtr,
			Courses:             courses,
		}, nil
	}
}
