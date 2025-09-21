package handler

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/validator"
)

type GetCourseTopicsByCourseIDQueryHandler struct {
	handler.Authenticable
	getCourseTopicsByCourseIDUseCase usecase.GetCourseTopicsByCourseIDUseCase
	msgProvider                      message.Provider
	logger                           logging.Logger
	validator                        validator.Validator
}

func NewGetCourseTopicsByCourseIDQueryHandler(
	getCourseTopicsByCourseIDUseCase usecase.GetCourseTopicsByCourseIDUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *GetCourseTopicsByCourseIDQueryHandler {
	return &GetCourseTopicsByCourseIDQueryHandler{
		Authenticable:                    handler.Authenticable{PermissionsRequired: []string{"list"}},
		getCourseTopicsByCourseIDUseCase: getCourseTopicsByCourseIDUseCase,
		msgProvider:                      msgProvider,
		logger:                           logger,
		validator:                        validator,
	}
}

func (h *GetCourseTopicsByCourseIDQueryHandler) Handle(request query.GetCourseTopicsByCourseIDQuery, lang string) (response.ApiResponse[response.GetCourseTopicsByCourseIDResponse], error) {
	h.logger.Info("Inicio de consulta de tópicos de curso")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al obtener tópicos de curso. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.GetCourseTopicsByCourseIDResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario no tiene permisos para obtener tópicos de curso")
		return response.ApiResponse[response.GetCourseTopicsByCourseIDResponse]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	results, course, err := h.getCourseTopicsByCourseIDUseCase.Execute(request.CourseID, lang)

	if err != nil {
		h.logger.Error("Error al obtener tópicos de curso: %v", err)
		return response.ApiResponse[response.GetCourseTopicsByCourseIDResponse]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(err, h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_server_error")),
			},
		}, nil
	}

	// Map to response
	var topicsResponse []response.CourseTopicDetailResponse
	var totalHours float32 = 0.0
	var courseName, courseCode string

	for _, courseTopic := range results {
		var topicName, knowledgeAreaName string
		var knowledgeAreaID uint

		if courseTopic.Topic != nil {
			topicName = courseTopic.Topic.Name
			if courseTopic.Topic.KnowledgeArea != nil {
				knowledgeAreaID = courseTopic.Topic.KnowledgeAreaID
				knowledgeAreaName = courseTopic.Topic.KnowledgeArea.Name
			}
		}

		topicsResponse = append(topicsResponse, response.CourseTopicDetailResponse{
			TopicID:           courseTopic.TopicID,
			TopicName:         topicName,
			StudyHours:        courseTopic.StudyHours,
			KnowledgeAreaID:   knowledgeAreaID,
			KnowledgeAreaName: knowledgeAreaName,
		})

		totalHours += courseTopic.StudyHours
	}

	// Get course name and code from course entity
	if course != nil {
		courseName = course.Name
		courseCode = course.Code
	}

	resultMapped := response.GetCourseTopicsByCourseIDResponse{
		CourseID:    request.CourseID,
		CourseName:  courseName,
		CourseCode:  courseCode,
		Topics:      topicsResponse,
		TotalTopics: len(topicsResponse),
		TotalHours:  totalHours,
	}

	h.logger.Info("Consulta de tópicos de curso completada exitosamente")
	return response.ApiResponse[response.GetCourseTopicsByCourseIDResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         resultMapped,
	}, nil
}
