package command

import (
	"errors"
	"net/http"

	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/app/domain/errs"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/validator"
)

type DeactivateUserCommandHandler struct {
	handler.Authenticable
	deactivateUserUseCase usecase.DeactivateUserUseCase
	msgProvider           message.Provider
	logger                logging.Logger
	validator             validator.Validator
}

func NewDeactivateUserCommandHandler(
	deactivateUserUseCase usecase.DeactivateUserUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *DeactivateUserCommandHandler {
	return &DeactivateUserCommandHandler{
		Authenticable:         handler.Authenticable{PermissionsRequired: []string{"deactivate"}},
		deactivateUserUseCase: deactivateUserUseCase,
		msgProvider:           msgProvider,
		logger:                logger,
		validator:             validator,
	}
}

func (h *DeactivateUserCommandHandler) Handle(cmd command.DeactivateUserCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio de desactivación de usuario con ID: %d", cmd.UserID)

	if err := h.validator.ValidateFields(&cmd, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al desactivar usuario. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(cmd.Permissions) {
		h.logger.Warn("El usuario no tiene permisos para desactivar usuarios")
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	err := h.deactivateUserUseCase.Execute(cmd.UserID, lang)
	if err != nil {
		h.logger.Error("Error al desactivar usuario con ID %d: %v", cmd.UserID, err)

		// Manejar diferentes tipos de errores
		if errors.Is(err, errs.NotFoundError) {
			return response.ApiResponse[string]{
				HttpStatusCode: http.StatusNotFound,
				Errors:         []string{err.Error()},
			}, nil
		}
		if errors.Is(err, errs.ConflictError) {
			return response.ApiResponse[string]{
				HttpStatusCode: http.StatusConflict,
				Errors:         []string{err.Error()},
			}, nil
		}

		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Usuario con ID %d desactivado exitosamente", cmd.UserID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         h.msgProvider.WithLang(lang).GetMessage("deactivate_user", "user_deactivated_successfully"),
		Errors:         []string{},
	}, nil
}
