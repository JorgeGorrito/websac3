package command

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/validator"
)

type EvaluateDegreeProgramCommandHandler struct {
	validator                    validator.Validator
	logger                       logging.Logger
	evaluateDegreeProgramUseCase usecase.EvaluateDegreeProgramUseCase
	msgProvider                  message.Provider
}

func NewEvaluateDegreeProgramCommandHandler(
	validator validator.Validator,
	logger logging.Logger,
	evaluateDegreeProgramUseCase usecase.EvaluateDegreeProgramUseCase,
	msgProvider message.Provider,
) *EvaluateDegreeProgramCommandHandler {
	return &EvaluateDegreeProgramCommandHandler{
		validator:                    validator,
		logger:                       logger,
		evaluateDegreeProgramUseCase: evaluateDegreeProgramUseCase,
		msgProvider:                  msgProvider,
	}
}

func (h *EvaluateDegreeProgramCommandHandler) Handle(req command.EvaluateDegreeProgramCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la evaluación del programa de grado con ID: %d para el rol profesional ID: %d", req.DegreeProgramID, req.ProfessionalRoleID)

	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al evaluar programa de grado. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{HttpStatusCode: http.StatusBadRequest, Errors: validationErrors}, nil
	}

	// Ejecutar evaluación y envío de reporte
	if err := h.evaluateDegreeProgramUseCase.Execute(req.DegreeProgramID, req.ProfessionalRoleID, req.UserID, lang); err != nil {
		h.logger.Error("Error al evaluar programa de grado con ID: %d para el rol profesional ID: %d y usuario ID: %d. Error: %s", req.DegreeProgramID, req.ProfessionalRoleID, req.UserID, err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[string]{
			HttpStatusCode: httpStatusCode,
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				),
			},
		}, nil
	}

	h.logger.Info("Evaluación del programa de grado con ID: %d completada exitosamente", req.DegreeProgramID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         h.msgProvider.WithLang(lang).GetMessage("evaluate_degree_program", "report_sent_successfully"),
	}, nil
}
