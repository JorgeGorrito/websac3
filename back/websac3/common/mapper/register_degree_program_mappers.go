package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerDegreeProgramMappers() {
	RegisterMapFunc(
		func(degreeProgramModel *model.DegreeProgram) (entity.DegreeProgram, error) {
			return entity.DegreeProgram{
				ID:                  degreeProgramModel.ID,
				Snies:               degreeProgramModel.Snies,
				Name:                degreeProgramModel.Name,
				TotalCredits:        degreeProgramModel.TotalCredits,
				DurationValue:       degreeProgramModel.DurationValue,
				DurationUnitID:      degreeProgramModel.DurationUnitID,
				ProgramFocus:        degreeProgramModel.ProgramFocus,
				EntryProfile:        degreeProgramModel.EntryProfile,
				GraduateProfile:     degreeProgramModel.GraduateProfile,
				ProfessionalProfile: degreeProgramModel.ProfessionalProfile,
				CreatedBy:           degreeProgramModel.CreatedBy,
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
				ProgramFocus:        command.ProgramFocus,
				EntryProfile:        command.EntryProfile,
				GraduateProfile:     command.GraduateProfile,
				ProfessionalProfile: command.ProfessionalProfile,
				CreatedBy:           command.CreatedBy,
			}, nil
		},
	)
}
