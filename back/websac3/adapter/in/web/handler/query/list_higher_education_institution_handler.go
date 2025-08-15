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

type ListHigherEducationInstitutionQueryHandler struct {
	listHigherEducationInstitutionUseCase usecase.ListHigherEducationInstitutionUseCase
	msgProvider                           message.Provider
	validator                             validator.Validator
	logger                                logging.Logger

	validFilters []string
}

func NewListHigherEducationInstitutionQueryHandler(
	listHigherEducationInstitutionUseCase usecase.ListHigherEducationInstitutionUseCase,
	msgProvider message.Provider,
	validator validator.Validator,
	logger logging.Logger,
) *ListHigherEducationInstitutionQueryHandler {
	return &ListHigherEducationInstitutionQueryHandler{
		listHigherEducationInstitutionUseCase: listHigherEducationInstitutionUseCase,
		msgProvider:                           msgProvider,
		validator:                             validator,
		logger:                                logger,
		validFilters:                          []string{"name"},
	}
}

func (h *ListHigherEducationInstitutionQueryHandler) Handle(request query.ListHigherEducationInstitutionQuery, lang string) (response.ApiResponse[paginator.Page[response.ListHigherEducationInstitutionResponse]], error) {
	h.logger.Info("Inicio la consulta de instituciones de educacion superior")

	if errs := h.validator.ValidateFields(&request, lang); errs != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al listar instituciones de educacion superior. Errores: %v", errs)
		var validationErrors []string
		for _, err := range errs {
			validationErrors = append(validationErrors, err.Error())
		}
		return response.ApiResponse[paginator.Page[response.ListHigherEducationInstitutionResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	results, total, err := h.listHigherEducationInstitutionUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, filters, lang)
	var resultsMapped []response.ListHigherEducationInstitutionResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListHigherEducationInstitutionResponse
		resultMapped, errMap = mapper.Map[entity.HigherEducationInstitution, response.ListHigherEducationInstitutionResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear resultado de solicitud de acceso. Error: %s", errMap.Error())
			return response.ApiResponse[paginator.Page[response.ListHigherEducationInstitutionResponse]]{
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

	resultPaginated, errPag := paginator.New[response.ListHigherEducationInstitutionResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar resultados de instituciones de educación superior. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListHigherEducationInstitutionResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err != nil {
		h.logger.Error("Error al consultar instituciones de educación superior. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListHigherEducationInstitutionResponse]]{
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

	h.logger.Info("Consulta de instituciones de educación superior finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListHigherEducationInstitutionResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
