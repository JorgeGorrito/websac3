package service

import (
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type UpdateDegreeProgramService struct {
	updateDegreeProgramPort persistence.UpdateDegreeProgramPort
	getDegreeProgramPort    persistence.GetDegreeProgramPort
	getProfessionalRolePort persistence.GetProfessionalRolePort
	messageProvider         message.Provider
	persistenceManager      db.Manager
}

func NewUpdateDegreeProgramService(
	updateDegreeProgramPort persistence.UpdateDegreeProgramPort,
	getDegreeProgramPort persistence.GetDegreeProgramPort,
	getProfessionalRolePort persistence.GetProfessionalRolePort,
	messageProvider message.Provider,
	persistenceManager db.Manager,
) usecase.UpdateDegreeProgramUseCase {
	return &UpdateDegreeProgramService{
		updateDegreeProgramPort: updateDegreeProgramPort,
		getDegreeProgramPort:    getDegreeProgramPort,
		getProfessionalRolePort: getProfessionalRolePort,
		messageProvider:         messageProvider,
		persistenceManager:      persistenceManager,
	}
}

func (s *UpdateDegreeProgramService) Execute(
	command command.UpdateDegreeProgramCommand,
	lang string,
) error {
	return s.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			// Verificar que el programa de grado existe
			existingDegreeProgram, err := s.getDegreeProgramPort.GetByID(command.ID, ctx)
			if err != nil {
				return err
			}

			// Actualizar la entidad DegreeProgram
			degreeProgram := entity.DegreeProgram{
				ID:                  command.ID,
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
				CreatedBy:           existingDegreeProgram.CreatedBy, // Mantener el creador original
			}

			// Asignar roles profesionales antes de actualizar
			if len(command.ProfessionalRoleIDs) > 0 {
				var professionalRoles []entity.ProfessionalRole
				for _, roleID := range command.ProfessionalRoleIDs {
					role, err := s.getProfessionalRolePort.GetByID(roleID, lang, ctx)
					if err != nil {
						return err
					}
					professionalRoles = append(professionalRoles, role)
				}
				degreeProgram.ProfessionalRoles = professionalRoles
			}

			// Actualizar el programa de grado
			if err := s.updateDegreeProgramPort.Update(&degreeProgram, ctx); err != nil {
				return err
			}

			return nil
		},
	)
}
