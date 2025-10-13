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

type ChangePasswordController struct{ Authenticable }

var instanceChangePasswordController *ChangePasswordController = nil

func GetChangePasswordController() *ChangePasswordController {
	if instanceChangePasswordController == nil {
		instanceChangePasswordController = &ChangePasswordController{}
	}
	return instanceChangePasswordController
}

// ChangePassword cambia la contraseña del usuario autenticado
// @Summary Cambiar Contraseña
// @Description Permite al usuario autenticado cambiar su contraseña.
// @Description Requiere la contraseña actual, la nueva contraseña y su confirmación.
// @Description La nueva contraseña debe cumplir con los siguientes requisitos:
// @Description - Tener un mínimo de 8 caracteres
// @Description - Ser diferente a la contraseña actual
// @Description - Incluir al menos una letra
// @Description - Incluir al menos un número
// @Description - Incluir al menos un símbolo especial
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.ChangePasswordRequest true "ChangePasswordRequest"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[string] "Contraseña cambiada exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Datos de entrada inválidos (contraseña actual incorrecta, contraseñas no coinciden, no cumple requisitos de complejidad, etc.)"
// @Failure 401 {object} response.ApiResponse[string] "Usuario no autenticado"
// @Failure 404 {object} response.ApiResponse[string] "Usuario no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/change-password [put]
// @Security BearerAuth
func (c *ChangePasswordController) ChangePassword(context *gin.Context) {
	var err error
	var changePasswordRequest request.ChangePasswordRequest
	if err = context.ShouldBindJSON(&changePasswordRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	var changePasswordCommand command.ChangePasswordCommand
	changePasswordCommand, err = mapper.Map[request.ChangePasswordRequest, command.ChangePasswordCommand](&changePasswordRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	token := c.GetToken(context)
	lang := context.Param("lang")
	changePasswordCommand.UserID = token.Sub

	result, _ := mediator.Send[command.ChangePasswordCommand, response.ApiResponse[string]](changePasswordCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
