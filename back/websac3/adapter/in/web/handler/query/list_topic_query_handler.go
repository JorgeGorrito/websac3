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

type ListTopicQueryHandler struct {
	handler.Authenticable
	listTopicUseCase usecase.ListTopicUseCase
	msgProvider      message.Provider
	logger           logging.Logger
	validator        validator.Validator

	validFilters []string
}

func NewListTopicQueryHandler(
	listTopicUseCase usecase.ListTopicUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *ListTopicQueryHandler {
	return &ListTopicQueryHandler{
		Authenticable:    handler.Authenticable{PermissionsRequired: []string{"list"}},
		listTopicUseCase: listTopicUseCase,
		msgProvider:      msgProvider,
		logger:           logger,
		validator:        validator,

		validFilters: []string{"name", "id"},
	}
}

func (h *ListTopicQueryHandler) Handle(request query.ListTopicQuery, lang string) (response.ApiResponse[paginator.Page[response.ListTopicResponse]], error) {
	h.logger.Info("Inicio de consulta de tematicas de ciberseguridad")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar tematicas de ciberseguridad. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListTopicResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario no tiene permisos para listar temáticas de ciberseguridad")
		return response.ApiResponse[paginator.Page[response.ListTopicResponse]]{
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
	var name string = ""
	var id string = ""
	for _, f := range filters {
		if f.Field == "name" {
			name, _ = f.Value.(string)
		} else if f.Field == "id" {
			id, _ = f.Value.(string)
		}
	}
	results, total, err := h.listTopicUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, name, id, lang)

	var resultsMapped []response.ListTopicResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListTopicResponse
		resultMapped, errMap = mapper.Map[entity.Topic, response.ListTopicResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear la temática de ciberseguridad: %v", errMap)
			return response.ApiResponse[paginator.Page[response.ListTopicResponse]]{
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

	resultPaginated, errPag := paginator.New[response.ListTopicResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar los resultados de la consulta de tematicas de ciberseguridad. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListTopicResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, errPag
	}
	if err != nil {
		h.logger.Error("Error al obtener los resultados de la consulta de tematicas de ciberseguridad. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListTopicResponse]]{
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

	h.logger.Info("Consulta de temáticas de ciberseguridad finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListTopicResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
