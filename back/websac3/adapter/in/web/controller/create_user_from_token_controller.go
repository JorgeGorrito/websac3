package controller

import (
	"net/http"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mapper"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type CreateUserFromTokenController struct{}

var instanceCreateUserFromTokenController *CreateUserFromTokenController = nil

func GetCreateUserFromTokenController() *CreateUserFromTokenController {
	if instanceCreateUserFromTokenController == nil {
		instanceCreateUserFromTokenController = &CreateUserFromTokenController{}
	}
	return instanceCreateUserFromTokenController
}

// CreateUserFromToken manejador para crear un usuario desde un token de creación
// @Summary Crear Usuario desde Token
// @Description Crea un usuario utilizando un token de creación válido y una contraseña.
// @Description El token debe corresponder a una solicitud de acceso aprobada y verificada.
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.CreateUserFromTokenRequest true "CreateUserFromTokenRequest"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 201 {object} response.ApiResponse[string] "Usuario creado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido o contraseñas no coinciden"
// @Failure 404 {object} response.ApiResponse[string] "Token de creación no encontrado"
// @Failure 409 {object} response.ApiResponse[string] "Usuario ya existe"
// @Failure 422 {object} response.ApiResponse[string] "Solicitud no aprobada o email no verificado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/user/create-from-token [post]
func (c *CreateUserFromTokenController) CreateUserFromToken(context *gin.Context) {
	var err error
	var createUserFromTokenRequest request.CreateUserFromTokenRequest

	if err = context.ShouldBindJSON(&createUserFromTokenRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createUserFromTokenCommand command.CreateUserFromTokenCommand
	createUserFromTokenCommand, err = mapper.Map[request.CreateUserFromTokenRequest, command.CreateUserFromTokenCommand](&createUserFromTokenRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lang := context.Param("lang")

	result, _ := mediator.Send[command.CreateUserFromTokenCommand, response.ApiResponse[string]](createUserFromTokenCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
