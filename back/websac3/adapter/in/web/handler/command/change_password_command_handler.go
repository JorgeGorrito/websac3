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

type ChangePasswordCommandHandler struct {
	handler.Authenticable
	changePasswordUseCase usecase.ChangePasswordUseCase
	validator             validator.Validator
	logger                logging.Logger
	msgProvider           message.Provider
}

func NewChangePasswordCommandHandler(
	changePasswordUseCase usecase.ChangePasswordUseCase,
	validator validator.Validator,
	logger logging.Logger,
	msgProvider message.Provider,
) *ChangePasswordCommandHandler {
	return &ChangePasswordCommandHandler{
		Authenticable:         handler.Authenticable{PermissionsRequired: []string{}},
		changePasswordUseCase: changePasswordUseCase,
		validator:             validator,
		logger:                logger,
		msgProvider:           msgProvider,
	}
}

func (h *ChangePasswordCommandHandler) Handle(request command.ChangePasswordCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio cambio de contraseña para el usuario con ID: %d", request.UserID)

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación de datos de entrada al cambiar contraseña. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Validar que las contraseñas nuevas coincidan
	if request.NewPassword != request.ConfirmNewPassword {
		h.logger.Warn("Las contraseñas nuevas no coinciden para el usuario ID: %d", request.UserID)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("change_password", "passwords_mismatch"),
			},
		}, nil
	}

	if err := h.changePasswordUseCase.Execute(request, lang); err != nil {
		h.logger.Error("Error al cambiar contraseña del usuario ID %d. Error: %s", request.UserID, err.Error())
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

	h.logger.Info("Contraseña cambiada exitosamente para el usuario con ID: %d", request.UserID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         h.msgProvider.WithLang(lang).GetMessage("change_password", "password_changed_successfully"),
	}, nil
}
