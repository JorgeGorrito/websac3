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
	"websac3/common/mapper"
	"websac3/common/paginator"
	"websac3/common/validator"
)

type ListCourseTypesQueryHandler struct {
	handler.Authenticable
	listCourseTypesUseCase usecase.ListCourseTypesUseCase
	msgProvider            message.Provider
	logger                 logging.Logger
	validator              validator.Validator
	validFilters           []string
}

func NewListCourseTypesQueryHandler(
	listCourseTypesUseCase usecase.ListCourseTypesUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *ListCourseTypesQueryHandler {
	return &ListCourseTypesQueryHandler{
		Authenticable:          handler.Authenticable{PermissionsRequired: []string{"list"}},
		listCourseTypesUseCase: listCourseTypesUseCase,
		msgProvider:            msgProvider,
		logger:                 logger,
		validator:              validator,
		validFilters:           []string{"name"},
	}
}

func (h *ListCourseTypesQueryHandler) Handle(request query.ListCourseTypesQuery, lang string) (response.ApiResponse[paginator.Page[response.ListCourseTypesResponse]], error) {
	h.logger.Info("Inicio de consulta de tipos de curso")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar tipos de curso. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListCourseTypesResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario no tiene permisos para listar tipos de curso")
		return response.ApiResponse[paginator.Page[response.ListCourseTypesResponse]]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	// Extract filters
	var name string
	if nameFilter, exists := request.Filters["name"]; exists {
		if nameStr, ok := nameFilter.(string); ok {
			name = nameStr
		}
	}

	results, total, err := h.listCourseTypesUseCase.Execute(
		request.Currentpage,
		request.ItemsPerpage,
		name,
		lang,
	)

	if err != nil {
		h.logger.Error("Error al listar tipos de curso: %v", err)
		return response.ApiResponse[paginator.Page[response.ListCourseTypesResponse]]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(err, h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_server_error")),
			},
		}, nil
	}

	// Map to response
	var resultsMapped []response.ListCourseTypesResponse
	for _, result := range results {
		resultMapped, errMap := mapper.Map[entity.CourseType, response.ListCourseTypesResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear el resultado de la consulta de tipos de curso: %v", errMap)
			return response.ApiResponse[paginator.Page[response.ListCourseTypesResponse]]{
				HttpStatusCode: http.StatusInternalServerError,
				Errors: []string{
					h.msgProvider.
						WithLang(lang).
						GetMessage("base_error", "internal_server_error"),
				},
			}, nil
		}
		resultsMapped = append(resultsMapped, resultMapped)
	}

	page := paginator.Page[response.ListCourseTypesResponse]{
		Data:         resultsMapped,
		TotalCount:   total,
		Currentpage:  request.Currentpage,
		ItemsPerpage: request.ItemsPerpage,
	}

	h.logger.Info("Consulta de tipos de curso completada exitosamente")
	return response.ApiResponse[paginator.Page[response.ListCourseTypesResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         page,
	}, nil
}
