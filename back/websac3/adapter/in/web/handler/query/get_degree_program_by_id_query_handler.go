package handler

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/validator"
)

type GetDegreeProgramByIDQueryHandler struct {
	validator                   validator.Validator
	logger                      logging.Logger
	getDegreeProgramByIDUseCase usecase.GetDegreeProgramByIDUseCase
	msgProvider                 message.Provider
}

func NewGetDegreeProgramByIDQueryHandler(
	validator validator.Validator,
	logger logging.Logger,
	getDegreeProgramByIDUseCase usecase.GetDegreeProgramByIDUseCase,
	msgProvider message.Provider,
) *GetDegreeProgramByIDQueryHandler {
	return &GetDegreeProgramByIDQueryHandler{
		validator:                   validator,
		logger:                      logger,
		getDegreeProgramByIDUseCase: getDegreeProgramByIDUseCase,
		msgProvider:                 msgProvider,
	}
}

func (h *GetDegreeProgramByIDQueryHandler) Handle(req query.GetDegreeProgramByIDQuery, lang string) (response.ApiResponse[response.GetDegreeProgramByIDResponse], error) {
	h.logger.Info("Inicio de consulta de detalle de programa de grado ID: %d", req.DegreeProgramID)

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al obtener detalle de programa de grado. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.GetDegreeProgramByIDResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Ejecutar caso de uso
	result, err := h.getDegreeProgramByIDUseCase.Execute(req.DegreeProgramID, lang)
	if err != nil {
		h.logger.Error("Error al obtener detalle de programa de grado ID %d: %v", req.DegreeProgramID, err)
		return response.ApiResponse[response.GetDegreeProgramByIDResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Consulta de detalle de programa de grado ID %d completada exitosamente", req.DegreeProgramID)
	return result, nil
}
