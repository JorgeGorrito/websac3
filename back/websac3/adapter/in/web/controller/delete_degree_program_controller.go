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

type DeleteDegreeProgramController struct{ Authenticable }

var deleteDegreeProgramController *DeleteDegreeProgramController = nil

func GetDeleteDegreeProgramController() *DeleteDegreeProgramController {
	if deleteDegreeProgramController == nil {
		deleteDegreeProgramController = &DeleteDegreeProgramController{}
	}
	return deleteDegreeProgramController
}

// Handle elimina un programa de grado
// @Summary Eliminar Programa de Grado
// @Description Elimina un programa de grado específico. Solo el usuario que creó el programa o un administrador puede eliminarlo.
// @Description La eliminación es lógica (soft delete), por lo que el programa no se borra físicamente de la base de datos.
// @Tags DegreeProgram
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(es) Enums(en, es)
// @Param degree_program_id path int true "ID del programa de grado a eliminar"
// @Success 200 {object} response.ApiResponse[string] "Programa de grado eliminado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ID inválido)"
// @Failure 401 {object} response.ApiResponse[string] "Usuario no autenticado"
// @Failure 403 {object} response.ApiResponse[string] "No autorizado para eliminar este programa de grado"
// @Failure 404 {object} response.ApiResponse[string] "Programa de grado no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program/{degree_program_id} [delete]
// @Security BearerAuth
func (c *DeleteDegreeProgramController) Handle(ctx *gin.Context) {
	token := c.GetToken(ctx)
	lang := ctx.Param("lang")

	// Get degree program ID from URL parameter
	degreeProgramIDStr := ctx.Param("degree_program_id")
	degreeProgramID, err := strconv.ParseUint(degreeProgramIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid degree program ID"})
		return
	}

	// Create request DTO
	var req request.DeleteDegreeProgramRequest
	req.DegreeProgramID = uint(degreeProgramID)

	// Map request to command
	var cmd command.DeleteDegreeProgramCommand
	cmd, err = mapper.Map[request.DeleteDegreeProgramRequest, command.DeleteDegreeProgramCommand](&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add additional fields from token
	cmd.UserID = token.Sub
	cmd.Permissions = token.Permissions["degree-programs"]

	result, _ := mediator.Send[command.DeleteDegreeProgramCommand, response.ApiResponse[string]](cmd, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}


