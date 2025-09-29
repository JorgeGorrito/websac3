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

type RejectExpertConsultationCommandHandler struct {
	handler.Authenticable
	validator                       validator.Validator
	logger                          logging.Logger
	rejectExpertConsultationUseCase usecase.RejectExpertConsultationUseCase
	messageProvider                 message.Provider
}

func NewRejectExpertConsultationCommandHandler(
	validator validator.Validator,
	logger logging.Logger,
	rejectExpertConsultationUseCase usecase.RejectExpertConsultationUseCase,
	messageProvider message.Provider,
) *RejectExpertConsultationCommandHandler {
	return &RejectExpertConsultationCommandHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"update"},
		},
		validator:                       validator,
		logger:                          logger,
		rejectExpertConsultationUseCase: rejectExpertConsultationUseCase,
		messageProvider:                 messageProvider,
	}
}

func (h *RejectExpertConsultationCommandHandler) Handle(req command.RejectExpertConsultationCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Rechazando solicitud de asesoría ID: %d por experto ID: %d", req.ConsultationID, req.ExpertID)

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
	err := h.rejectExpertConsultationUseCase.Execute(req.ConsultationID, req.ExpertResponse, req.ExpertID, lang)
	if err != nil {
		h.logger.Error("Error al rechazar solicitud de asesoría: %v", err)
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

	h.logger.Info("Solicitud de asesoría rechazada exitosamente. ID: %d", req.ConsultationID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result: h.messageProvider.
			WithLang(lang).
			GetMessage("reject_expert_consultation", "success"),
	}, nil
}
