package command

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/jwt"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/validator"
)

type LoginHandler struct {
	loginUseCase usecase.LoginUseCase
	jwtGenerator jwt.Generator
	msgProvider  message.Provider
	validator    validator.Validator
	logger       logging.Logger
}

func NewLoginHandler(
	loginUseCase usecase.LoginUseCase,
	jwtGenerator jwt.Generator,
	msgProvider message.Provider,
	validator validator.Validator,
	logger logging.Logger,
) *LoginHandler {
	return &LoginHandler{
		loginUseCase: loginUseCase,
		jwtGenerator: jwtGenerator,
		msgProvider:  msgProvider,
		validator:    validator,
		logger:       logger,
	}
}

func (h *LoginHandler) Handle(request command.LoginCommand, lang string) (response.ApiResponse[jwt.TokenJWTResponse], error) {
	h.logger.Info("Inicio la petición para iniciar sesión al usuario con email: " + request.Email)

	if validationErrors := h.validator.ValidateFields(&request, lang); validationErrors != nil {
		validationErrorsMsg := func() []string {
			var validationErrorsMsg []string
			for _, err := range validationErrors {
				validationErrorsMsg = append(validationErrorsMsg, err.Error())
			}
			return validationErrorsMsg
		}()
		h.logger.Warn("Advertencia de validación datos de entrada al iniciar sesión. Errores: %v", validationErrorsMsg)
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrorsMsg,
		}, nil
	}

	user, err := mapper.Map[command.LoginCommand, entity.User](&request)
	if err != nil {
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	user, err = h.loginUseCase.Execute(user, lang)
	if err != nil {
		h.logger.Error("Error con las credenciales del usuario con email `%s`. Error: %s", request.Email, err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[jwt.TokenJWTResponse]{
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

	accessTokenClaims, err := mapper.Map[entity.User, jwt.AccessTokenClaims](&user)
	if err != nil {
		h.logger.Error("Error al mapear el usuario a la respuesta del token. Error: %s", err.Error())
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	refreshTokenclaims, err := mapper.Map[entity.User, jwt.RefreshTokenClaims](&user)
	if err != nil {
		h.logger.Error("Error al mapear el usuario a la respuesta del token de refresco. Error: %s", err.Error())
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	token, err := h.jwtGenerator.GenerateTokens(
		accessTokenClaims,
		refreshTokenclaims,
	)
	if err != nil {
		h.logger.Error("Error al generar los tokens JWT. Error: %s", err.Error())
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Finalizó inicio de sesión exitoso para el usuario con email: " + user.Email)
	return response.ApiResponse[jwt.TokenJWTResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         token,
	}, nil
}
