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

type ListRejectedAccessRequestQueryHandler struct {
	handler.Authenticable
	listRejectedAccessRequestUseCase usecase.ListRejectedAccessRequestUseCase
	msgProvider                      message.Provider
	validator                        validator.Validator
	logger                           logging.Logger

	validFilters []string
}

func NewListRejectedAccessRequestQueryHandler(
	listRejectedAccessRequestUseCase usecase.ListRejectedAccessRequestUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *ListRejectedAccessRequestQueryHandler {
	return &ListRejectedAccessRequestQueryHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"list"},
		},
		listRejectedAccessRequestUseCase: listRejectedAccessRequestUseCase,
		msgProvider:                      msgProvider,
		validator:                        validator,
		logger:                           logger,
		validFilters: []string{
			"Applicant.name",
			"Applicant.lastname",
			"Applicant.HigherEducationInstitution.name",
			"Applicant.HigherEducationInstitution.Municipality.name",
			"Applicant.HigherEducationInstitution.Department.name",
			"Status.name",
		},
	}
}

func (h *ListRejectedAccessRequestQueryHandler) Handle(request query.ListRejectedAccessRequestQuery, lang string) (response.ApiResponse[paginator.Page[response.ListRejectedAccessRequestResponse]], error) {
	h.logger.Info("Inicio de consulta de solicitudes de acceso rechazadas")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar solicitudes de acceso rechazadas. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListRejectedAccessRequestResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	results, total, err := h.listRejectedAccessRequestUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, filters, request.UserID, request.Permissions, lang)

	var resultsMapped []response.ListRejectedAccessRequestResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListRejectedAccessRequestResponse
		resultMapped, errMap = mapper.Map[entity.AccessRequest, response.ListRejectedAccessRequestResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear resultado de solicitud de acceso rechazada. Error: %s", errMap.Error())
			return response.ApiResponse[paginator.Page[response.ListRejectedAccessRequestResponse]]{
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

	resultPaginated, errPag := paginator.New[response.ListRejectedAccessRequestResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar resultados de solicitudes de acceso rechazadas. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListRejectedAccessRequestResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err != nil {
		h.logger.Error("Error al consultar solicitudes de acceso rechazadas. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListRejectedAccessRequestResponse]]{
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

	h.logger.Info("Consulta de solicitudes de acceso rechazadas finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListRejectedAccessRequestResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
