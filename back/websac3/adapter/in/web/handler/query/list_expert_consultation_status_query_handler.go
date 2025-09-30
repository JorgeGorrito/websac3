package handler

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/paginator"
	"websac3/common/validator"
)

type ListExpertConsultationStatusQueryHandler struct {
	validator                           validator.Validator
	logger                              logging.Logger
	listExpertConsultationStatusUseCase usecase.ListExpertConsultationStatusUseCase
	msgProvider                         message.Provider
}

func NewListExpertConsultationStatusQueryHandler(
	validator validator.Validator,
	logger logging.Logger,
	listExpertConsultationStatusUseCase usecase.ListExpertConsultationStatusUseCase,
	msgProvider message.Provider,
) *ListExpertConsultationStatusQueryHandler {
	return &ListExpertConsultationStatusQueryHandler{
		validator:                           validator,
		logger:                              logger,
		listExpertConsultationStatusUseCase: listExpertConsultationStatusUseCase,
		msgProvider:                         msgProvider,
	}
}

func (h *ListExpertConsultationStatusQueryHandler) Handle(req query.ListExpertConsultationStatusQuery, lang string) (response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]], error) {
	h.logger.Info("Inicio de consulta de estados de asesorías de experto")

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al obtener estados de asesorías de experto. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Ejecutar caso de uso
	result, err := h.listExpertConsultationStatusUseCase.Execute(req, lang)
	if err != nil {
		h.logger.Error("Error al obtener estados de asesorías de experto: %v", err)
		return response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Consulta de estados de asesorías de experto completada exitosamente")
	return result, nil
}
