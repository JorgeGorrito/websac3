package service

import (
	"time"
	"websac3/adapter/in/web/response"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type GetDegreeProgramByIDService struct {
	getDegreeProgramPort persistence.GetDegreeProgramPort
	messageProvider      message.Provider
	persistenceManager   db.Manager
}

func NewGetDegreeProgramByIDService(
	getDegreeProgramPort persistence.GetDegreeProgramPort,
	messageProvider message.Provider,
	persistenceManager db.Manager,
) usecase.GetDegreeProgramByIDUseCase {
	return &GetDegreeProgramByIDService{
		getDegreeProgramPort: getDegreeProgramPort,
		messageProvider:      messageProvider,
		persistenceManager:   persistenceManager,
	}
}

func (s *GetDegreeProgramByIDService) Execute(degreeProgramID uint, lang string) (response.ApiResponse[response.GetDegreeProgramByIDResponse], error) {
	var degreeProgram entity.DegreeProgram
	var err error

	// Obtener el programa de grado por ID
	err = s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		degreeProgram, err = s.getDegreeProgramPort.GetByID(degreeProgramID, ctx)
		if err != nil {
			return errs.NewNotFoundError(s.messageProvider.WithLang(lang).GetMessage("degree_program", "not_found"))
		}
		return nil
	})

	if err != nil {
		return response.ApiResponse[response.GetDegreeProgramByIDResponse]{}, err
	}

	// Mapear a response DTO
	degreeProgramResponse, err := s.mapDegreeProgramToResponse(degreeProgram, lang)
	if err != nil {
		return response.ApiResponse[response.GetDegreeProgramByIDResponse]{}, err
	}

	successMessage := s.messageProvider.WithLang(lang).GetMessage("degree_program", "retrieved_successfully")
	return response.ApiResponse[response.GetDegreeProgramByIDResponse]{
		HttpStatusCode: 200,
		Result:         degreeProgramResponse,
		Errors:         []string{successMessage},
	}, nil
}

// mapDegreeProgramToResponse mapea la entidad DegreeProgram a la respuesta detallada
func (s *GetDegreeProgramByIDService) mapDegreeProgramToResponse(degreeProgram entity.DegreeProgram, lang string) (response.GetDegreeProgramByIDResponse, error) {
	// Mapear unidad de duración
	var durationUnitResponse response.ListDurationUnitResponse
	if degreeProgram.DurationUnit != nil {
		var err error
		durationUnitResponse, err = mapper.Map[entity.DurationUnit, response.ListDurationUnitResponse](degreeProgram.DurationUnit)
		if err != nil {
			return response.GetDegreeProgramByIDResponse{}, err
		}
	}

	// Mapear nivel de formación
	var formationLevelResponse response.ListFormationLevelsResponse
	if degreeProgram.FormationLevel != nil {
		var err error
		formationLevelResponse, err = mapper.Map[entity.FormationLevel, response.ListFormationLevelsResponse](degreeProgram.FormationLevel)
		if err != nil {
			return response.GetDegreeProgramByIDResponse{}, err
		}
	}

	// Mapear roles profesionales
	var professionalRolesResponse []response.ListProfessionalRoleResponse
	for _, role := range degreeProgram.ProfessionalRoles {
		roleResponse, err := mapper.Map[entity.ProfessionalRole, response.ListProfessionalRoleResponse](&role)
		if err != nil {
			return response.GetDegreeProgramByIDResponse{}, err
		}
		professionalRolesResponse = append(professionalRolesResponse, roleResponse)
	}

	// Mapear información de la institución
	var institutionInfo *response.HigherEducationInstitutionInfo
	if degreeProgram.UserCreator != nil && degreeProgram.UserCreator.Person != nil && degreeProgram.UserCreator.Person.HigherEducationInstitution != nil {
		institution := degreeProgram.UserCreator.Person.HigherEducationInstitution
		institutionInfo = &response.HigherEducationInstitutionInfo{
			Snies:                 institution.Snies,
			Name:                  institution.Name,
			Ownership:             institution.Ownership.Name,
			InstitutionalCategory: institution.InstitutionalCategory.Name,
			Municipality:          institution.Municipality.Name,
			Department:            institution.Department.Name,
		}
	}

	return response.GetDegreeProgramByIDResponse{
		ID:                         degreeProgram.ID,
		Snies:                      degreeProgram.Snies,
		Name:                       degreeProgram.Name,
		TotalCredits:               degreeProgram.TotalCredits,
		DurationValue:              degreeProgram.DurationValue,
		DurationUnit:               durationUnitResponse,
		FormationLevel:             formationLevelResponse,
		ProfessionalRoles:          professionalRolesResponse,
		ProgramFocus:               degreeProgram.ProgramFocus,
		EntryProfile:               degreeProgram.EntryProfile,
		GraduateProfile:            degreeProgram.GraduateProfile,
		ProfessionalProfile:        degreeProgram.ProfessionalProfile,
		CreatedBy:                  degreeProgram.CreatedBy,
		CreatedAt:                  time.Now().Format(time.RFC3339), // TODO: Usar el campo real cuando esté disponible
		UpdatedAt:                  time.Now().Format(time.RFC3339), // TODO: Usar el campo real cuando esté disponible
		HigherEducationInstitution: institutionInfo,
	}, nil
}
