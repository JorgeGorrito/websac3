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
	"websac3/app/port/out/persistence/filter"
	futil "websac3/common/filter"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/paginator"
	"websac3/common/validator"
)

type ListDegreeProgramQueryHandler struct {
	handler.Authenticable
	listDegreeProgramUseCase usecase.ListDegreeProgramUseCase
	msgProvider              message.Provider
	logger                   logging.Logger
	validator                validator.Validator

	validFilters []string
}

func NewListDegreeProgramQueryHandler(
	listDegreeProgramUseCase usecase.ListDegreeProgramUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *ListDegreeProgramQueryHandler {
	return &ListDegreeProgramQueryHandler{
		Authenticable:            handler.Authenticable{PermissionsRequired: []string{"list"}},
		listDegreeProgramUseCase: listDegreeProgramUseCase,
		msgProvider:              msgProvider,
		logger:                   logger,
		validator:                validator,

		validFilters: []string{"name", "snies"},
	}
}

func (h *ListDegreeProgramQueryHandler) Handle(request query.ListDegreeProgramQuery, lang string) (response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]], error) {
	h.logger.Info("Inicio de consulta de programas de grado")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar programas de grado. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para listar programas de grado", request.UserID)
		return response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	results, total, err := h.listDegreeProgramUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, filters, request.UserRole, request.UserID, lang)

	var resultsMapped []response.ListDegreeProgramResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListDegreeProgramResponse
		resultMapped, errMap = mapper.Map[entity.DegreeProgram, response.ListDegreeProgramResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear el programa de grado: %v", errMap)
			return response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]]{
				HttpStatusCode: http.StatusInternalServerError,
				Errors: []string{
					h.msgProvider.
						WithLang(lang).
						GetMessage("base_error", "internal_error"),
				},
			}, errMap
		}
		resultsMapped = append(resultsMapped, resultMapped)
	}

	resultPaginated, errPag := paginator.New[response.ListDegreeProgramResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar los resultados de la consulta de programas de grado. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, errPag
	}
	if err != nil {
		h.logger.Error("Error al obtener los resultados de la consulta de programas de grado. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]]{
			HttpStatusCode: httpStatusCode,
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.
						WithLang(lang).
						GetMessage("base_error", "internal_error"),
				),
			},
		}, err
	}

	h.logger.Info("Consulta de programas de grado finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
