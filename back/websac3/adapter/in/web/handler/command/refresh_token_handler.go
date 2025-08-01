package command

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/jwt"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/validator"
)

type RefreshTokenHandler struct {
	getUserByIDUseCase usecase.GetUserByIDUseCase
	jwtGenerator       jwt.Generator
	msgProvider        message.Provider
	validator          validator.Validator
	logger             logging.Logger
}

func NewRefreshTokenHandler(
	getUserByIDUseCase usecase.GetUserByIDUseCase,
	jwtGenerator jwt.Generator,
	msgProvider message.Provider,
	validator validator.Validator,
	logger logging.Logger,
) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		getUserByIDUseCase: getUserByIDUseCase,
		jwtGenerator:       jwtGenerator,
		msgProvider:        msgProvider,
		validator:          validator,
		logger:             logger,
	}
}

func (h *RefreshTokenHandler) Handle(request command.RefreshTokenCommand, lang string) (response.ApiResponse[jwt.TokenJWTResponse], error) {
	h.logger.Info("Inicio la petición para refrescar un token JWT: %s", request.RefreshToken)

	if validationErrors := h.validator.ValidateFields(&request, lang); validationErrors != nil {
		validationErrorsMsg := func() []string {
			var validationErrorsMsg []string
			for _, err := range validationErrors {
				validationErrorsMsg = append(validationErrorsMsg, err.Error())
			}
			return validationErrorsMsg
		}()
		h.logger.Warn("Advertencia de validación datos de entrada al refrescar token. Errores: %v", validationErrorsMsg)
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrorsMsg,
		}, nil
	}

	refreshToken, err := h.jwtGenerator.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		h.logger.Error("Error al validar el token de refresco: %v", err)
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusUnauthorized,
			Errors:         []string{h.msgProvider.WithLang(lang).GetMessage("refresh_token", "invalid_refresh_token")},
		}, nil
	}

	user, err := h.getUserByIDUseCase.Execute(refreshToken.Sub, lang)
	if err != nil {
		h.logger.Error("Error al obtener el usuario por ID: %v", err)
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors:         []string{h.msgProvider.WithLang(lang).GetMessage("internal_error", "internal_server_error")},
		}, nil
	}

	accessTokenClaims, err := mapper.Map[entity.User, jwt.AccessTokenClaims](&user)
	if err != nil {
		h.logger.Error("Error al mapear el usuario a la respuesta de token: %v", err)
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors:         []string{h.msgProvider.WithLang(lang).GetMessage("internal_error", "internal_server_error")},
		}, nil
	}

	accessToken, err := h.jwtGenerator.GenerateAccessToken(accessTokenClaims)
	if err != nil {
		h.logger.Error("Error al generar el token de acceso: %v", err)
		return response.ApiResponse[jwt.TokenJWTResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors:         []string{h.msgProvider.WithLang(lang).GetMessage("internal_error", "internal_server_error")},
		}, nil
	}

	h.logger.Info("Token de refresco validado correctamente para el usuario con ID: %d", refreshToken.Sub)
	return response.ApiResponse[jwt.TokenJWTResponse]{
		HttpStatusCode: http.StatusOK,
		Result: jwt.TokenJWTResponse{
			AccessToken:  accessToken,
			RefreshToken: request.RefreshToken,
		},
	}, nil

}
