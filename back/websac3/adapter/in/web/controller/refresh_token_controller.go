package controller

import (
	"net/http"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/jwt"
	"websac3/common/mapper"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct{}

var refreshTokenControllerInstance *RefreshTokenController = nil

func GetRefreshTokenController() *RefreshTokenController {
	if refreshTokenControllerInstance == nil {
		refreshTokenControllerInstance = &RefreshTokenController{}
	}
	return refreshTokenControllerInstance
}

// Handle refresco de token JWT
// @Summary Refrescar token JWT
// @Description Recibe un token de actualización válido y devuelve un nuevo token JWT.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.RefreshTokenRequest true "Token de actualización"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[jwt.TokenJWTResponse] "Token actualizado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Solicitud inválida (ejemplo: JSON mal formado)"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/auth/refresh [post]
func (c *RefreshTokenController) Handle(ctx *gin.Context) {
	var err error
	var refreshTokenRequest request.RefreshTokenRequest
	if err = ctx.ShouldBindJSON(&refreshTokenRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var refreshTokenCommand command.RefreshTokenCommand
	refreshTokenCommand, err = mapper.Map[request.RefreshTokenRequest, command.RefreshTokenCommand](&refreshTokenRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := mediator.Send[command.RefreshTokenCommand, response.ApiResponse[jwt.TokenJWTResponse]](refreshTokenCommand, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
	return
}
