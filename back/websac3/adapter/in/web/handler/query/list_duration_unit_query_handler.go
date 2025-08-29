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

type ListDurationUnitQueryHandler struct {
	listDurationUnitUseCase usecase.ListDurationUnitUseCase
	msgProvider             message.Provider
	logger                  logging.Logger
	validator               validator.Validator

	validFilters []string
}

func NewListDurationUnitQueryHandler(
	listDurationUnitUseCase usecase.ListDurationUnitUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *ListDurationUnitQueryHandler {
	return &ListDurationUnitQueryHandler{
		listDurationUnitUseCase: listDurationUnitUseCase,
		msgProvider:             msgProvider,
		logger:                  logger,
		validator:               validator,

		validFilters: []string{"name"},
	}
}

func (h *ListDurationUnitQueryHandler) Handle(request query.ListDurationUnitQuery, lang string) (response.ApiResponse[paginator.Page[response.ListDurationUnitResponse]], error) {
	h.logger.Info("Inicio de consulta de unidades de duración")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar unidades de duración. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListDurationUnitResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	var name string = ""
	if len(filters) > 0 {
		name, _ = filters[0].Value.(string)
	}
	results, total, err := h.listDurationUnitUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, name, lang)

	var resultsMapped []response.ListDurationUnitResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListDurationUnitResponse
		resultMapped, errMap = mapper.Map[entity.DurationUnit, response.ListDurationUnitResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear la unidad de duración: %v", errMap)
			return response.ApiResponse[paginator.Page[response.ListDurationUnitResponse]]{
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

	resultPaginated, errPag := paginator.New[response.ListDurationUnitResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar los resultados de la consulta de unidades de duración. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListDurationUnitResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, errPag
	}
	if err != nil {
		h.logger.Error("Error al obtener los resultados de la consulta de unidades de duración. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListDurationUnitResponse]]{
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

	h.logger.Info("Consulta de unidades de duración finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListDurationUnitResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
