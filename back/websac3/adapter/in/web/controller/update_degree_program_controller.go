package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mapper"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type UpdateDegreeProgramController struct{ Authenticable }

var instanceUpdateDegreeProgramController *UpdateDegreeProgramController = nil

func GetUpdateDegreeProgramController() *UpdateDegreeProgramController {
	if instanceUpdateDegreeProgramController == nil {
		instanceUpdateDegreeProgramController = &UpdateDegreeProgramController{}
	}
	return instanceUpdateDegreeProgramController
}

// UpdateDegreeProgram manejador para actualizar un programa de grado
// @Summary Actualizar Programa de Grado
// @Description Actualiza la información de un programa de grado existente.
// @Description El programa de grado incluye información sobre duración, nivel de formación, enfoque del programa,
// @Description perfiles de entrada y egreso, y roles profesionales asociados.
// @Description Los roles profesionales deben ser IDs válidos de roles existentes en el sistema.
// @Tags DegreeProgram
// @Accept json
// @Produce json
// @Param degree_program_id path int true "ID del programa de grado"
// @Param request body request.UpdateDegreeProgramRequest true "UpdateDegreeProgramRequest"
// @Param lang path string true "Código de idioma" default(en) English Enums(en, es)
// @Success 200 {object} response.ApiResponse[string] "Programa de grado actualizado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "Programa de grado no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program/{degree_program_id} [put]
// @Security BearerAuth
func (c *UpdateDegreeProgramController) UpdateDegreeProgram(context *gin.Context) {
	var err error
	var updateDegreeProgramRequest request.UpdateDegreeProgramRequest
	if err = context.ShouldBindJSON(&updateDegreeProgramRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	// Obtener el ID del programa de grado desde la URL
	degreeProgramIDStr := context.Param("degree_program_id")
	degreeProgramID, err := strconv.ParseUint(degreeProgramIDStr, 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": "Invalid degree program ID format"})
		return
	}

	var updateDegreeProgramCommand command.UpdateDegreeProgramCommand
	updateDegreeProgramCommand, err = mapper.Map[request.UpdateDegreeProgramRequest, command.UpdateDegreeProgramCommand](&updateDegreeProgramRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	// Asignar el ID del programa de grado
	updateDegreeProgramCommand.ID = uint(degreeProgramID)

	token := c.GetToken(context)
	lang := context.Param("lang")
	updateDegreeProgramCommand.UpdatedBy = token.Sub
	updateDegreeProgramCommand.Permissions = token.Permissions["degree-programs"]

	result, _ := mediator.Send[command.UpdateDegreeProgramCommand, response.ApiResponse[string]](updateDegreeProgramCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
