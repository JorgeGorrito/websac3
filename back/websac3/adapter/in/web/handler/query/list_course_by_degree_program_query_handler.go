package handler

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence/filter"
	futil "websac3/common/filter"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/paginator"
	"websac3/common/validator"
)

type ListCourseByDegreeProgramQueryHandler struct {
	handler.Authenticable
	listCourseModelsByDegreeProgramUseCase usecase.ListCourseModelsByDegreeProgramUseCase
	msgProvider                            message.Provider
	logger                                 logging.Logger
	validator                              validator.Validator

	validFilters []string
}

func NewListCourseByDegreeProgramQueryHandler(
	listCourseModelsByDegreeProgramUseCase usecase.ListCourseModelsByDegreeProgramUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *ListCourseByDegreeProgramQueryHandler {
	return &ListCourseByDegreeProgramQueryHandler{
		Authenticable:                          handler.Authenticable{PermissionsRequired: []string{"list"}},
		listCourseModelsByDegreeProgramUseCase: listCourseModelsByDegreeProgramUseCase,
		msgProvider:                            msgProvider,
		logger:                                 logger,
		validator:                              validator,

		validFilters: []string{"name"},
	}
}

func (h *ListCourseByDegreeProgramQueryHandler) Handle(request query.ListCourseByDegreeProgramQuery, lang string) (response.ApiResponse[paginator.Page[response.ListCourseByDegreeProgramResponse]], error) {
	h.logger.Info("Inicio de consulta de cursos por programa de grado")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar cursos por programa de grado. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListCourseByDegreeProgramResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario no tiene permisos para listar cursos por programa de grado")
		return response.ApiResponse[paginator.Page[response.ListCourseByDegreeProgramResponse]]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	pagination := &request.PaginationParams
	// Convert map[string]interface{} to futil.Params
	filterParams := make(futil.Params)
	for k, v := range request.Filters {
		if nestedMap, ok := v.(map[string]string); ok {
			filterParams[k] = nestedMap
		}
	}
	filters := futil.Transform(filterParams, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	var name string = ""
	if len(filters) > 0 {
		name, _ = filters[0].Value.(string)
	}

	results, total, err := h.listCourseModelsByDegreeProgramUseCase.Execute(
		request.DegreeProgramID,
		pagination.Currentpage,
		pagination.ItemsPerpage,
		name,
		lang,
	)

	var resultsMapped []response.ListCourseByDegreeProgramResponse
	var errMap error
	for _, result := range results {
		// Debug logging
		h.logger.Info("Course ID: %d, Nature: %+v, Type: %+v", result.ID, result.Nature, result.Type)

		var resultMapped response.ListCourseByDegreeProgramResponse
		if resultMapped, errMap = mapper.Map[model.Course, response.ListCourseByDegreeProgramResponse](&result); errMap != nil {
			h.logger.Error("Error al mapear el resultado de la consulta de cursos por programa de grado: %v", errMap)
			return response.ApiResponse[paginator.Page[response.ListCourseByDegreeProgramResponse]]{
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

	if err != nil {
		h.logger.Error("Error al consultar cursos por programa de grado: %v", err)
		return response.ApiResponse[paginator.Page[response.ListCourseByDegreeProgramResponse]]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(err, h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_server_error")),
			},
		}, nil
	}

	page := paginator.Page[response.ListCourseByDegreeProgramResponse]{
		Data:         resultsMapped,
		TotalCount:   total,
		Currentpage:  pagination.Currentpage,
		ItemsPerpage: pagination.ItemsPerpage,
	}

	h.logger.Info("Consulta de cursos por programa de grado completada exitosamente")
	return response.ApiResponse[paginator.Page[response.ListCourseByDegreeProgramResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         page,
	}, nil
}
