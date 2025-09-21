package handler

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/validator"
)

type GetCourseByIDQueryHandler struct {
	handler.Authenticable
	getCourseByIDUseCase usecase.GetCourseByIDUseCase
	msgProvider          message.Provider
	logger               logging.Logger
	validator            validator.Validator
}

func NewGetCourseByIDQueryHandler(
	getCourseByIDUseCase usecase.GetCourseByIDUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *GetCourseByIDQueryHandler {
	return &GetCourseByIDQueryHandler{
		Authenticable:        handler.Authenticable{PermissionsRequired: []string{"read"}},
		getCourseByIDUseCase: getCourseByIDUseCase,
		msgProvider:          msgProvider,
		logger:               logger,
		validator:            validator,
	}
}

func (h *GetCourseByIDQueryHandler) Handle(request query.GetCourseByIDQuery, lang string) (response.ApiResponse[response.GetCourseByIDResponse], error) {
	h.logger.Info("Inicio de consulta de detalle de curso")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al obtener detalle de curso. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.GetCourseByIDResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario no tiene permisos para obtener detalle de curso")
		return response.ApiResponse[response.GetCourseByIDResponse]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	result, err := h.getCourseByIDUseCase.Execute(request.CourseID, lang)

	if err != nil {
		h.logger.Error("Error al obtener detalle de curso: %v", err)
		return response.ApiResponse[response.GetCourseByIDResponse]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(err, h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_server_error")),
			},
		}, nil
	}

	// Map to response manually to handle language
	resultMapped := h.mapCourseToResponse(result, lang)

	h.logger.Info("Consulta de detalle de curso completada exitosamente")
	return response.ApiResponse[response.GetCourseByIDResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         resultMapped,
	}, nil
}

// Helper method to map course entity to response with language support
func (h *GetCourseByIDQueryHandler) mapCourseToResponse(course *entity.Course, lang string) response.GetCourseByIDResponse {
	var natureName, typeName, degreeProgramName, creatorName string

	// Get nature name by language
	if course.Nature != nil {
		natureName = getNameByLanguage(course.Nature.Names, lang)
	}

	// Get type name by language
	if course.Type != nil {
		typeName = getTypeNameByLanguage(course.Type.Names, lang)
	}

	// Get degree program name
	if course.DegreeProgram != nil {
		degreeProgramName = course.DegreeProgram.Name
	}

	// Get creator name
	if course.UserCreator != nil && course.UserCreator.Person != nil {
		creatorName = course.UserCreator.Person.Name + " " + course.UserCreator.Person.Lastname
	}

	// Map course topics
	var courseTopicsResponse []response.CourseTopicResponse
	for _, courseTopic := range course.CourseTopics {
		var topicName string
		if courseTopic.Topic != nil {
			topicName = courseTopic.Topic.Name
		}
		courseTopicsResponse = append(courseTopicsResponse, response.CourseTopicResponse{
			TopicID:    courseTopic.TopicID,
			TopicName:  topicName,
			StudyHours: courseTopic.StudyHours,
		})
	}

	return response.GetCourseByIDResponse{
		ID:                          course.ID,
		Name:                        course.Name,
		Code:                        course.Code,
		Credits:                     course.Credits,
		PeriodNumber:                course.PeriodNumber,
		NatureID:                    course.NatureID,
		NatureName:                  natureName,
		TypeID:                      course.TypeID,
		TypeName:                    typeName,
		IsCybersecurity:             course.IsCybersecurity,
		ContainsCybersecurityTopics: course.ContainsCybersecurityTopics(),
		DegreeProgramID:             course.DegreeProgramID,
		DegreeProgramName:           degreeProgramName,
		CreatedBy:                   course.CreatedBy,
		CreatorName:                 creatorName,
		CourseTopics:                courseTopicsResponse,
	}
}

// Helper functions to get names by language
func getNameByLanguage(names []entity.CourseNatureName, lang string) string {
	for _, name := range names {
		if name.Lang == lang {
			return name.Name
		}
	}
	// If no name found for the language, return the first available name
	if len(names) > 0 {
		return names[0].Name
	}
	return ""
}

func getTypeNameByLanguage(names []entity.CourseTypeName, lang string) string {
	for _, name := range names {
		if name.Lang == lang {
			return name.Name
		}
	}
	// If no name found for the language, return the first available name
	if len(names) > 0 {
		return names[0].Name
	}
	return ""
}
