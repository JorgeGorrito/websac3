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

type ListUserExpertConsultationsQueryHandler struct {
	handler.Authenticable
	validator                          validator.Validator
	logger                             logging.Logger
	listUserExpertConsultationsUseCase usecase.ListUserExpertConsultationsUseCase
	msgProvider                        message.Provider
	validFilters                       []string
}

func NewListUserExpertConsultationsQueryHandler(
	validator validator.Validator,
	logger logging.Logger,
	listUserExpertConsultationsUseCase usecase.ListUserExpertConsultationsUseCase,
	msgProvider message.Provider,
) *ListUserExpertConsultationsQueryHandler {
	return &ListUserExpertConsultationsQueryHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"list"},
		},
		validator:                          validator,
		logger:                             logger,
		listUserExpertConsultationsUseCase: listUserExpertConsultationsUseCase,
		msgProvider:                        msgProvider,
		validFilters: []string{
			"Status.name",
			"DegreeProgram.name",
		},
	}
}

func (h *ListUserExpertConsultationsQueryHandler) Handle(req query.ListUserExpertConsultationsQuery, lang string) (response.ApiResponse[paginator.Page[response.UserExpertConsultationResponse]], error) {
	h.logger.Info("Listando solicitudes de asesoría para usuario ID: %d", req.UserID)

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Error("Error de validación: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.UserExpertConsultationResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Validar permisos - TEMPORALMENTE DESHABILITADO PARA DEBUGGING
	// if !h.ValidatePermissions(req.Permissions) {
	// 	h.logger.Warn("El usuario con ID: %d no tiene permisos para listar sus solicitudes de asesoría", req.UserID)
	// 	return response.ApiResponse[paginator.Page[response.UserExpertConsultationResponse]]{
	// 		HttpStatusCode: http.StatusForbidden,
	// 		Errors: []string{
	// 			h.msgProvider.
	// 				WithLang(lang).
	// 				GetMessage("base_error", "forbidden"),
	// 		},
	// 	}, nil
	// }

	// Validar y limpiar filtros
	transformedFilters := commonfilter.Transform(req.Filters, []psqlfilter.Operator{psqlfilter.EqualOperator, psqlfilter.ContainsOperator})
	transformedFilters.Purge(h.validFilters)

	// Ejecutar caso de uso
	expertConsultations, total, err := h.listUserExpertConsultationsUseCase.Execute(req.UserID, req.PaginationParams, req.Filters, lang)
	if err != nil {
		h.logger.Error("Error al obtener solicitudes de asesoría: %v", err)
		return response.ApiResponse[paginator.Page[response.UserExpertConsultationResponse]]{
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
	var expertConsultationResponses []response.UserExpertConsultationResponse
	for _, expertConsultation := range expertConsultations {
		var expertConsultationResponse response.UserExpertConsultationResponse
		if expertConsultationResponse, err = mapper.Map[entity.ExpertConsultation, response.UserExpertConsultationResponse](&expertConsultation); err != nil {
			h.logger.Error("Error al mapear solicitud de asesoría: %v", err)
			continue
		}
		expertConsultationResponses = append(expertConsultationResponses, expertConsultationResponse)
	}

	paginatedResponse := paginator.Page[response.UserExpertConsultationResponse]{
		Data:         expertConsultationResponses,
		TotalCount:   int64(total),
		Currentpage:  req.PaginationParams.Currentpage,
		ItemsPerpage: req.PaginationParams.ItemsPerpage,
	}

	h.logger.Info("Listado de solicitudes de asesoría completado. Total: %d", total)
	return response.ApiResponse[paginator.Page[response.UserExpertConsultationResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         paginatedResponse,
	}, nil
}
