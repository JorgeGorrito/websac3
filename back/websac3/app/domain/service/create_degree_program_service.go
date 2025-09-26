package service

import (
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type CreateDegreeProgramService struct {
	createDegreeProgramPort persistence.CreateDegreeProgramPort
	getProfessionalRolePort persistence.GetProfessionalRolePort
	persistenceManager      db.Manager
	msgProvider             message.Provider
}

func NewCreateDegreeProgramService(
	createDegreeProgramPort persistence.CreateDegreeProgramPort,
	getProfessionalRolePort persistence.GetProfessionalRolePort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *CreateDegreeProgramService {
	return &CreateDegreeProgramService{
		createDegreeProgramPort: createDegreeProgramPort,
		getProfessionalRolePort: getProfessionalRolePort,
		persistenceManager:      persistenceManager,
		msgProvider:             msgProvider,
	}
}

func (s *CreateDegreeProgramService) Execute(
	command command.CreateDegreeProgramCommand,
	lang string,
) error {
	return s.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			// Crear la entidad DegreeProgram básica
			degreeProgram := entity.DegreeProgram{
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
			}

			// Asignar roles profesionales antes de crear
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

			// Crear el programa de grado
			if err := s.createDegreeProgramPort.Create(&degreeProgram, ctx); err != nil {
				return err
			}

			return nil
		},
	)
}
