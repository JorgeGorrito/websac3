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

type ListIdentificationTypeQueryHandler struct {
	listIdentificationTypeUseCase usecase.ListIdentificationTypeUseCase
	msgProvider                   message.Provider
	validator                     validator.Validator
	logger                        logging.Logger

	validFilters []string
}

func NewListIdentificationTypeQueryHandler(
	listIdentificationTypeUseCase usecase.ListIdentificationTypeUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *ListIdentificationTypeQueryHandler {
	return &ListIdentificationTypeQueryHandler{
		listIdentificationTypeUseCase: listIdentificationTypeUseCase,
		msgProvider:                   msgProvider,
		validator:                     validator,
		logger:                        logger,

		validFilters: []string{
			"name",
		},
	}
}

func (h *ListIdentificationTypeQueryHandler) Handle(request query.ListIdentificationTypeQuery, lang string) (response.ApiResponse[paginator.Page[response.ListIdentificationTypeResponse]], error) {
	h.logger.Info("Inicio de consulta de tipos de identificación")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar tipos de identificación. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListIdentificationTypeResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	results, total, err := h.listIdentificationTypeUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, filters, lang)

	var resultsMapped []response.ListIdentificationTypeResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListIdentificationTypeResponse
		resultMapped, errMap = mapper.Map[entity.IdentificationType, response.ListIdentificationTypeResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear resultado de tipo de identificación. Error: %s", errMap.Error())
			return response.ApiResponse[paginator.Page[response.ListIdentificationTypeResponse]]{
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

	resultPaginated, errPag := paginator.New[response.ListIdentificationTypeResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar resultados de tipos de identificación. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListIdentificationTypeResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err != nil {
		h.logger.Error("Error al consultar tipos de identificación. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListIdentificationTypeResponse]]{
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

	h.logger.Info("Consulta de tipos de identificación finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListIdentificationTypeResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
