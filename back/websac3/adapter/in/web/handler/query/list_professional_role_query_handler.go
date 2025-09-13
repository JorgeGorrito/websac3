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

type ListProfessionalRoleQueryHandler struct {
	listProfessionalRoleUseCase usecase.ListProfessionalRoleUseCase
	msgProvider                 message.Provider
	validator                   validator.Validator
	logger                      logging.Logger

	validFilters []string
}

func NewListProfessionalRoleQueryHandler(
	listProfessionalRoleUseCase usecase.ListProfessionalRoleUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *ListProfessionalRoleQueryHandler {
	return &ListProfessionalRoleQueryHandler{
		listProfessionalRoleUseCase: listProfessionalRoleUseCase,
		msgProvider:                 msgProvider,
		validator:                   validator,
		logger:                      logger,

		validFilters: []string{
			"name",
		},
	}
}

func (h *ListProfessionalRoleQueryHandler) Handle(request query.ListProfessionalRoleQuery, lang string) (response.ApiResponse[paginator.Page[response.ListProfessionalRoleResponse]], error) {
	h.logger.Info("Inicio de consulta de roles profesionales")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar roles profesionales. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListProfessionalRoleResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	results, total, err := h.listProfessionalRoleUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, filters, lang)

	var resultsMapped []response.ListProfessionalRoleResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListProfessionalRoleResponse
		resultMapped, errMap = mapper.Map[entity.ProfessionalRole, response.ListProfessionalRoleResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear resultado de rol profesional. Error: %s", errMap.Error())
			return response.ApiResponse[paginator.Page[response.ListProfessionalRoleResponse]]{
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

	resultPaginated, errPag := paginator.New[response.ListProfessionalRoleResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar resultados de roles profesionales. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListProfessionalRoleResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err != nil {
		h.logger.Error("Error al consultar roles profesionales. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListProfessionalRoleResponse]]{
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

	h.logger.Info("Consulta de roles profesionales finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListProfessionalRoleResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
