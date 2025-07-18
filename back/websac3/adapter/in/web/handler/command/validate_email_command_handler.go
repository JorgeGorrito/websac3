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

type ValidateEmailCommandHandler struct {
	validateEmailUseCase usecase.ValidateEmailUseCase
	msgProvider          message.Provider
	validator            validator.Validator
	logger               logging.Logger
}

func NewValidateEmailCommandHandler(
	validateEmailUseCase usecase.ValidateEmailUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *ValidateEmailCommandHandler {
	return &ValidateEmailCommandHandler{
		validateEmailUseCase: validateEmailUseCase,
		msgProvider:          msgProvider,
		validator:            validator,
		logger:               logger,
	}
}

func (h *ValidateEmailCommandHandler) Handle(request command.ValidateEmailCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la validación de correo electrónico con token: " + request.ValidationToken)

	if validationErrors := h.validator.ValidateFields(&request, lang); validationErrors != nil {
		validationErrorsMsg := func() []string {
			var validationErrorsMsg []string
			for _, err := range validationErrors {
				validationErrorsMsg = append(validationErrorsMsg, err.Error())
			}
			return validationErrorsMsg
		}()
		h.logger.Warn("Advertencia de validación datos de entrada al validar correo electrónico. Errores: %v", validationErrorsMsg)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrorsMsg,
		}, nil
	}

	if err := h.validateEmailUseCase.Execute(request.ValidationToken, lang); err != nil {
		h.logger.Error("Error al validar el correo electrónico con token: %s. Error: %s", request.ValidationToken, err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[string]{
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

	h.logger.Info("Correo electrónico validado exitosamente con token: " + request.ValidationToken)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result: h.msgProvider.
			WithLang(lang).
			GetMessage("validate_email", "email_validated"),
	}, nil
}
