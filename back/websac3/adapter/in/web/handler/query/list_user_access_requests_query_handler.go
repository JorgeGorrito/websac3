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
	psqlfilter "websac3/app/port/out/persistence/filter"
	commonfilter "websac3/common/filter"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/paginator"
	"websac3/common/validator"
)

type ListUserAccessRequestsQueryHandler struct {
	handler.Authenticable
	validator                     validator.Validator
	logger                        logging.Logger
	listUserAccessRequestsUseCase usecase.ListUserAccessRequestsUseCase
	msgProvider                   message.Provider
	validFilters                  []string
}

func NewListUserAccessRequestsQueryHandler(
	validator validator.Validator,
	logger logging.Logger,
	listUserAccessRequestsUseCase usecase.ListUserAccessRequestsUseCase,
	msgProvider message.Provider,
) *ListUserAccessRequestsQueryHandler {
	return &ListUserAccessRequestsQueryHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"list"},
		},
		validator:                     validator,
		logger:                        logger,
		listUserAccessRequestsUseCase: listUserAccessRequestsUseCase,
		msgProvider:                   msgProvider,
		validFilters: []string{
			"Status.name",
			"Applicant.HigherEducationInstitution.name",
		},
	}
}

func (h *ListUserAccessRequestsQueryHandler) Handle(req query.ListUserAccessRequestsQuery, lang string) (response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]], error) {
	h.logger.Info("Listando solicitudes de acceso para usuario ID: %d", req.UserID)

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Error("Error de validación: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Validar permisos
	if !h.ValidatePermissions(req.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para listar sus solicitudes de acceso", req.UserID)
		return response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	// Validar y limpiar filtros
	transformedFilters := commonfilter.Transform(req.Filters, []psqlfilter.Operator{psqlfilter.EqualOperator, psqlfilter.ContainsOperator})
	transformedFilters.Purge(h.validFilters)

	// Ejecutar caso de uso
	accessRequests, total, err := h.listUserAccessRequestsUseCase.Execute(req.UserID, req.PaginationParams, req.Filters, lang)
	if err != nil {
		h.logger.Error("Error al obtener solicitudes de acceso: %v", err)
		return response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				),
			},
		}, nil
	}

	// Mapear a response
	var accessRequestResponses []response.UserAccessRequestResponse
	for _, accessRequest := range accessRequests {
		var accessRequestResponse response.UserAccessRequestResponse
		if accessRequestResponse, err = mapper.Map[entity.AccessRequest, response.UserAccessRequestResponse](&accessRequest); err != nil {
			h.logger.Error("Error al mapear solicitud de acceso: %v", err)
			continue
		}
		accessRequestResponses = append(accessRequestResponses, accessRequestResponse)
	}

	paginatedResponse := paginator.Page[response.UserAccessRequestResponse]{
		Data:         accessRequestResponses,
		TotalCount:   int64(total),
		Currentpage:  req.PaginationParams.Currentpage,
		ItemsPerpage: req.PaginationParams.ItemsPerpage,
	}

	h.logger.Info("Listado de solicitudes de acceso completado. Total: %d", total)
	return response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         paginatedResponse,
	}, nil
}
