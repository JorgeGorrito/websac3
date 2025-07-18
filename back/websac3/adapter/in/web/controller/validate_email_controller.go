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

type ValidateEmailController struct{}

var instanceValidateEmailController *ValidateEmailController = nil

func GetValidateEmailController() *ValidateEmailController {
	if instanceValidateEmailController == nil {
		instanceValidateEmailController = &ValidateEmailController{}
	}
	return instanceValidateEmailController
}

// ValidateEmail manejador para validar el correo electrónico mediante un token.
// @Summary Validar Correo Electrónico
// @Description Valida una solicitud de acceso utilizando un token previamente enviado al correo del usuario.
// @Description Si el token es válido, la solicitud de acceso se actualiza como verificada.
// @Description Si el token no existe o ha expirado, se devuelve un error 404.
// @Tags AccessRequest
// @Accept json
// @Produce json
// @Param request body request.ValidateEmailRequest true "Datos para validar el correo electrónico"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[string] "Correo verificado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Error de validación en el cuerpo de la solicitud"
// @Failure 404 {object} response.ApiResponse[string] "Token no válido o expirado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/access-request/email/validate [post]
func (v *ValidateEmailController) ValidateEmail(context *gin.Context) {
	var err error
	var validationEmailRequest request.ValidateEmailRequest
	var validationEmailCommand command.ValidateEmailCommand

	if err = context.ShouldBindJSON(&validationEmailRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	validationEmailCommand, err = mapper.Map[request.ValidateEmailRequest, command.ValidateEmailCommand](&validationEmailRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	result, _ := mediator.Send[command.ValidateEmailCommand, response.ApiResponse[string]](validationEmailCommand, context.GetString("lang"))
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
