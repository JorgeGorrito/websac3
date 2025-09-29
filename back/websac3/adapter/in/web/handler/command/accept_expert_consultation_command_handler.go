package command

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/validator"
)

type AcceptExpertConsultationCommandHandler struct {
	handler.Authenticable
	validator                       validator.Validator
	logger                          logging.Logger
	acceptExpertConsultationUseCase usecase.AcceptExpertConsultationUseCase
	messageProvider                 message.Provider
}

func NewAcceptExpertConsultationCommandHandler(
	validator validator.Validator,
	logger logging.Logger,
	acceptExpertConsultationUseCase usecase.AcceptExpertConsultationUseCase,
	messageProvider message.Provider,
) *AcceptExpertConsultationCommandHandler {
	return &AcceptExpertConsultationCommandHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"update"},
		},
		validator:                       validator,
		logger:                          logger,
		acceptExpertConsultationUseCase: acceptExpertConsultationUseCase,
		messageProvider:                 messageProvider,
	}
}

func (h *AcceptExpertConsultationCommandHandler) Handle(req command.AcceptExpertConsultationCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Aceptando solicitud de asesoría ID: %d por experto ID: %d", req.ConsultationID, req.ExpertID)

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Error("Error de validación: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Ejecutar caso de uso
	err := h.acceptExpertConsultationUseCase.Execute(req.ConsultationID, req.ExpertResponse, req.ExpertID, lang)
	if err != nil {
		h.logger.Error("Error al aceptar solicitud de asesoría: %v", err)
		return response.ApiResponse[string]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.messageProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				),
			},
		}, nil
	}

	h.logger.Info("Solicitud de asesoría aceptada exitosamente. ID: %d", req.ConsultationID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result: h.messageProvider.
			WithLang(lang).
			GetMessage("accept_expert_consultation", "success"),
	}, nil
}
