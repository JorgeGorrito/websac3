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

type CreateDegreeProgramController struct{ Authenticable }

var instanceCreateDegreeProgramController *CreateDegreeProgramController = nil

func GetCreateDegreeProgramController() *CreateDegreeProgramController {
	if instanceCreateDegreeProgramController == nil {
		instanceCreateDegreeProgramController = &CreateDegreeProgramController{}
	}
	return instanceCreateDegreeProgramController
}

// CreateDegreeProgram manejador para crear un programa de grado
// @Summary Crear Programa de Grado
// @Description Crea un nuevo programa de grado con toda la información académica necesaria.
// @Tags DegreeProgram
// @Accept json
// @Produce json
// @Param request body request.CreateDegreeProgramRequest true "CreateDegreeProgramRequest"
// @Param lang path string true "Código de idioma" default(en) English Enums(en, es)
// @Success 201 {object} response.ApiResponse[string] "Programa de grado creado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 409 {object} response.ApiResponse[string] "Programa de grado ya existe"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program [post]
// @Security BearerAuth
func (c *CreateDegreeProgramController) CreateDegreeProgram(context *gin.Context) {
	var err error
	var createDegreeProgramRequest request.CreateDegreeProgramRequest
	if err = context.ShouldBindJSON(&createDegreeProgramRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	var createDegreeProgramCommand command.CreateDegreeProgramCommand
	createDegreeProgramCommand, err = mapper.Map[request.CreateDegreeProgramRequest, command.CreateDegreeProgramCommand](&createDegreeProgramRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	token := c.GetToken(context)
	lang := context.Param("lang")
	createDegreeProgramCommand.CreatedBy = token.Sub
	createDegreeProgramCommand.Permissions = token.Permissions["degree-programs"]

	result, _ := mediator.Send[command.CreateDegreeProgramCommand, response.ApiResponse[string]](createDegreeProgramCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
