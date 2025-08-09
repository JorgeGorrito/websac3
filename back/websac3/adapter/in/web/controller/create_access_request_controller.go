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

type CreateAccessRequestController struct{}

var instanceCreateAccessRequestController *CreateAccessRequestController = nil

func GetCreateAccessRequestController() *CreateAccessRequestController {
	if instanceCreateAccessRequestController == nil {
		instanceCreateAccessRequestController = &CreateAccessRequestController{}
	}
	return instanceCreateAccessRequestController
}

// CreateAccessRequest manejador para crear una solicitud de acceso sin verificación de correo electrónico
// @Summary Crear Solicitud de Acceso Pendiente de Verificación de Correo Electrónico
// @Description Crea una solicitud de acceso pendiente de verificación de correo electrónico.
// @Description Toda solicitud creada genera un token de validación y lo concatena al enlace de redireccion de validacion de correo electrónico que se envía al usuario para verificar el correo electrónico.
// @Description Toda solicitud de acceso sin verificación de correo electrónico no se tendrá en cuenta en el proceso de aprobación de acceso al aplicativo.
// @Tags AccessRequest
// @Accept json
// @Produce json
// @Param request body request.CreateAccessRequestRequest true "CreateAccessRequestRequest"
// @Param lang path string true "Código de idioma" default(en) English Enums(en, es)
// @Success 201 {object} response.ApiResponse[string] "Solicitud de acceso creada exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 409 {object} response.ApiResponse[string] "Solicitud de acceso ya existe para el usuario"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/access-request [post]
func (c *CreateAccessRequestController) CreateAccessRequest(context *gin.Context) {
	var err error
	var createAccessRequestRequest request.CreateAccessRequestRequest
	if err = context.ShouldBindJSON(&createAccessRequestRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	var createAccessRequestCommand command.CreateAccessRequestCommand
	createAccessRequestCommand, err = mapper.Map[request.CreateAccessRequestRequest, command.CreateAccessRequestCommand](&createAccessRequestRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	result, _ := mediator.Send[command.CreateAccessRequestCommand, response.ApiResponse[string]](createAccessRequestCommand, context.GetString("lang"))
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
