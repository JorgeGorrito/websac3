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

type CreateUserFromTokenCommandHandler struct {
	handler.Authenticable
	createUserFromTokenUseCase usecase.CreateUserFromTokenUseCase
	msgProvider                message.Provider
	validator                  validator.Validator
	logger                     logging.Logger
}

func NewCreateUserFromTokenCommandHandler(
	createUserFromTokenUseCase usecase.CreateUserFromTokenUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *CreateUserFromTokenCommandHandler {
	return &CreateUserFromTokenCommandHandler{
		Authenticable:              handler.Authenticable{PermissionsRequired: []string{}}, // No requiere permisos
		createUserFromTokenUseCase: createUserFromTokenUseCase,
		msgProvider:                msgProvider,
		validator:                  validator,
		logger:                     logger,
	}
}

func (h *CreateUserFromTokenCommandHandler) Handle(request command.CreateUserFromTokenCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la creación de usuario desde token")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al crear usuario. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Validar que las contraseñas coincidan
	if request.Password != request.ConfirmPassword {
		h.logger.Warn("Las contraseñas no coinciden")
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("create_user", "passwords_mismatch"),
			},
		}, nil
	}

	if err := h.createUserFromTokenUseCase.Execute(request.CreateUserToken, request.Password, lang); err != nil {
		h.logger.Error("Error al crear usuario desde token. Error: %s", err.Error())
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

	h.logger.Info("Usuario creado exitosamente desde token")
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusCreated,
		Result:         h.msgProvider.WithLang(lang).GetMessage("create_user", "user_created_successfully"),
	}, nil
}
