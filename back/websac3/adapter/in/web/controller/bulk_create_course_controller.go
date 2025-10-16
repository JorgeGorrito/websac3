package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type BulkCreateCourseController struct {
	Authenticable
}

func GetBulkCreateCourseController() *BulkCreateCourseController {
	return &BulkCreateCourseController{}
}

// BulkCreateCourse maneja la creación masiva de cursos desde archivo CSV
// @Summary Creación Masiva de Cursos
// @Description Sube un archivo CSV y crea múltiples cursos en lote para un programa de grado específico.
// @Description El archivo CSV debe seguir el formato de la plantilla descargable y contener todas las columnas requeridas.
// @Description La operación procesa cada fila individualmente y reporta éxitos y errores por separado.
// @Tags Course
// @Accept multipart/form-data
// @Produce json
// @Param lang path string true "Código de idioma" default(es) Enums(en, es)
// @Param degree_program_id formData int true "ID del programa de grado"
// @Param file formData file true "Archivo CSV con datos de cursos"
// @Success 200 {object} response.ApiResponse[response.BulkCreateCourseResponse] "Resultado de la creación masiva"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (archivo CSV mal formateado)"
// @Failure 401 {object} response.ApiResponse[string] "Usuario no autenticado"
// @Failure 403 {object} response.ApiResponse[string] "No tiene permisos para realizar esta acción"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course/bulk [post]
// @Security BearerAuth
func (c *BulkCreateCourseController) BulkCreateCourse(context *gin.Context) {
	var err error
	var bulkCreateRequest request.BulkCreateCourseRequest
	if err = context.ShouldBind(&bulkCreateRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	// Parse degree_program_id from form
	degreeProgramIDStr := context.PostForm("degree_program_id")
	degreeProgramID, err := strconv.ParseUint(degreeProgramIDStr, 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": "Invalid degree_program_id"})
		return
	}

	var bulkCreateCommand command.BulkCreateCourseCommand
	bulkCreateCommand.File = bulkCreateRequest.File
	bulkCreateCommand.DegreeProgramID = uint(degreeProgramID)

	token := c.GetToken(context)
	lang := context.Param("lang")
	bulkCreateCommand.CreatedBy = token.Sub
	bulkCreateCommand.Permissions = token.Permissions["courses"]

	result, _ := mediator.Send[command.BulkCreateCourseCommand, response.ApiResponse[response.BulkCreateCourseResponse]](bulkCreateCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}





