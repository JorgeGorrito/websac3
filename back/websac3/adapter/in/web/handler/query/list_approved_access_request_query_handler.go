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

type ListApprovedAccessRequestQueryHandler struct {
	handler.Authenticable
	listApprovedAccessRequestUseCase usecase.ListApprovedAccessRequestUseCase
	msgProvider                      message.Provider
	validator                        validator.Validator
	logger                           logging.Logger

	validFilters []string
}

func NewListApprovedAccessRequestQueryHandler(
	listApprovedAccessRequestUseCase usecase.ListApprovedAccessRequestUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *ListApprovedAccessRequestQueryHandler {
	return &ListApprovedAccessRequestQueryHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"list"},
		},
		listApprovedAccessRequestUseCase: listApprovedAccessRequestUseCase,
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
			"ApprovedRole.name",
		},
	}
}

func (h *ListApprovedAccessRequestQueryHandler) Handle(request query.ListApprovedAccessRequestQuery, lang string) (response.ApiResponse[paginator.Page[response.ListApprovedAccessRequestResponse]], error) {
	h.logger.Info("Inicio de consulta de solicitudes de acceso aprobadas")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar solicitudes de acceso aprobadas. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListApprovedAccessRequestResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	results, total, err := h.listApprovedAccessRequestUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, filters, request.UserID, request.Permissions, lang)

	var resultsMapped []response.ListApprovedAccessRequestResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListApprovedAccessRequestResponse
		resultMapped, errMap = mapper.Map[entity.AccessRequest, response.ListApprovedAccessRequestResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear resultado de solicitud de acceso aprobada. Error: %s", errMap.Error())
			return response.ApiResponse[paginator.Page[response.ListApprovedAccessRequestResponse]]{
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

	resultPaginated, errPag := paginator.New[response.ListApprovedAccessRequestResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar resultados de solicitudes de acceso aprobadas. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListApprovedAccessRequestResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err != nil {
		h.logger.Error("Error al consultar solicitudes de acceso aprobadas. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListApprovedAccessRequestResponse]]{
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

	h.logger.Info("Consulta de solicitudes de acceso aprobadas finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListApprovedAccessRequestResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
