package controller

import (
	"net/http"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type BulkCreateDegreeProgramController struct {
	Authenticable
}

func GetBulkCreateDegreeProgramController() *BulkCreateDegreeProgramController {
	return &BulkCreateDegreeProgramController{}
}

// BulkCreateDegreeProgram maneja la creación masiva de programas de grado desde archivo CSV
// @Summary Creación Masiva de Programas de Grado
// @Description Sube un archivo CSV y crea múltiples programas de grado en lote.
// @Description El archivo CSV debe seguir el formato de la plantilla descargable y contener todas las columnas requeridas.
// @Description La operación procesa cada fila individualmente y reporta éxitos y errores por separado.
// @Tags DegreeProgram
// @Accept multipart/form-data
// @Produce json
// @Param lang path string true "Código de idioma" default(es) Enums(en, es)
// @Param file formData file true "Archivo CSV con datos de programas de grado"
// @Success 200 {object} response.ApiResponse[response.BulkCreateDegreeProgramResponse] "Resultado de la creación masiva"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (archivo CSV mal formateado)"
// @Failure 401 {object} response.ApiResponse[string] "Usuario no autenticado"
// @Failure 403 {object} response.ApiResponse[string] "No tiene permisos para realizar esta acción"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program/bulk [post]
// @Security BearerAuth
func (c *BulkCreateDegreeProgramController) BulkCreateDegreeProgram(context *gin.Context) {
	var err error
	var bulkCreateRequest request.BulkCreateDegreeProgramRequest
	if err = context.ShouldBind(&bulkCreateRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	var bulkCreateCommand command.BulkCreateDegreeProgramCommand
	bulkCreateCommand.File = bulkCreateRequest.File

	token := c.GetToken(context)
	lang := context.Param("lang")
	bulkCreateCommand.CreatedBy = token.Sub
	bulkCreateCommand.Permissions = token.Permissions["degree-programs"]

	result, _ := mediator.Send[command.BulkCreateDegreeProgramCommand, response.ApiResponse[response.BulkCreateDegreeProgramResponse]](bulkCreateCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
