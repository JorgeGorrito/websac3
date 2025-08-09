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

type ListAccessRequestQueryHandler struct {
	handler.Authenticable
	listAccessRequestUseCase usecase.ListAccessRequestUseCase
	msgProvider              message.Provider
	validator                validator.Validator
	logger                   logging.Logger

	validFilters []string
}

func NewListAccessRequestQueryHandler(
	listAccessRequestUseCase usecase.ListAccessRequestUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *ListAccessRequestQueryHandler {
	return &ListAccessRequestQueryHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"list"},
		},
		listAccessRequestUseCase: listAccessRequestUseCase,
		msgProvider:              msgProvider,
		validator:                validator,
		logger:                   logger,
		validFilters: []string{
			"Applicant.name",
			"Applicant.lastname",
			"Applicant.HigherEducationInstitution.name",
			"Applicant.HigherEducationInstitution.Municipality.name",
			"Applicant.HigherEducationInstitution.Department.name",
		},
	}
}

func (h *ListAccessRequestQueryHandler) Handle(request query.ListAccessRequestQuery, lang string) (response.ApiResponse[paginator.Page[response.ListAccessRequestResponse]], error) {
	h.logger.Info("Inicio la consulta de solicitudes de acceso")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al listar solicitudes de acceso. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListAccessRequestResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para listar las solicitudes de acceso", request.UserID)
		return response.ApiResponse[paginator.Page[response.ListAccessRequestResponse]]{
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
	results, total, err := h.listAccessRequestUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, filters, lang)

	var resultsMapped []response.ListAccessRequestResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListAccessRequestResponse
		resultMapped, errMap = mapper.Map[entity.AccessRequest, response.ListAccessRequestResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear resultado de solicitud de acceso. Error: %s", errMap.Error())
			return response.ApiResponse[paginator.Page[response.ListAccessRequestResponse]]{
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

	resultPaginated, errPag := paginator.New[response.ListAccessRequestResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar resultados de solicitudes de acceso. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListAccessRequestResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err != nil {
		h.logger.Error("Error al consultar solicitudes de acceso. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListAccessRequestResponse]]{
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

	h.logger.Info("Consulta de solicitudes de acceso finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListAccessRequestResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
