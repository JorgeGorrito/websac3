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

type LoginController struct{}

var loginControllerInstance *LoginController = nil

func GetLoginController() *LoginController {
	if loginControllerInstance == nil {
		loginControllerInstance = &LoginController{}
	}
	return loginControllerInstance
}

// Handle login de usuario
// @Summary Iniciar sesión
// @Description Autentica a un usuario con sus credenciales y devuelve un token JWT.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Credenciales de inicio de sesión"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[jwt.TokenJWTResponse] "Inicio de sesión exitoso"
// @Failure 400 {object} response.ApiResponse[string] "Solicitud inválida (ejemplo: JSON mal formado)"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/auth/ [post]
func (c *LoginController) Handle(ctx *gin.Context) {
	var err error
	var loginRequest request.LoginRequest

	if err = ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loginCommand command.LoginCommand
	loginCommand, err = mapper.Map[request.LoginRequest, command.LoginCommand](&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := mediator.Send[command.LoginCommand, response.ApiResponse[jwt.TokenJWTResponse]](loginCommand, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
	return
}
