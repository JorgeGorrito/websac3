package handler

import (
	"net/http"
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

type ListRoleQueryHandler struct {
	listRoleUseCase usecase.ListRoleUseCase
	msgProvider     message.Provider
	validator       validator.Validator
	logger          logging.Logger

	validFilters []string
}

func NewListRoleQueryHandler(
	listRoleUseCase usecase.ListRoleUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *ListRoleQueryHandler {
	return &ListRoleQueryHandler{
		listRoleUseCase: listRoleUseCase,
		msgProvider:     msgProvider,
		validator:       validator,
		logger:          logger,

		validFilters: []string{
			"name",
		},
	}
}

func (h *ListRoleQueryHandler) Handle(request query.ListRoleQuery, lang string) (response.ApiResponse[paginator.Page[response.ListRoleResponse]], error) {
	h.logger.Info("Inicio de consulta de roles del sistema")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar roles del sistema. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListRoleResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	results, total, err := h.listRoleUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, filters, lang)

	var resultsMapped []response.ListRoleResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListRoleResponse
		resultMapped, errMap = mapper.Map[entity.Role, response.ListRoleResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear resultado de rol del sistema. Error: %s", errMap.Error())
			return response.ApiResponse[paginator.Page[response.ListRoleResponse]]{
				HttpStatusCode: http.StatusInternalServerError,
				Errors: []string{
					h.msgProvider.
						WithLang(lang).
						GetMessage("base_error", "internal_error"),
				},
			}, nil
		}
		resultsMapped = append(resultsMapped, resultMapped)
	}

	resultPaginated, errPag := paginator.New[response.ListRoleResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar resultados de roles del sistema. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListRoleResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err != nil {
		h.logger.Error("Error al consultar roles del sistema. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListRoleResponse]]{
			Result:         *resultPaginated,
			HttpStatusCode: httpStatusCode,
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.
						WithLang(lang).
						GetMessage("base_error", "internal_error"),
				),
			},
		}, nil
	}

	h.logger.Info("Consulta de roles del sistema finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListRoleResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
