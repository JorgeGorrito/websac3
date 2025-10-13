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

type BulkCreateCourseTopicController struct {
	Authenticable
}

func GetBulkCreateCourseTopicController() *BulkCreateCourseTopicController {
	return &BulkCreateCourseTopicController{}
}

// BulkCreateCourseTopic maneja la creación masiva de tópicos de curso desde archivo CSV
// @Summary Creación Masiva de Tópicos de Curso
// @Description Sube un archivo CSV y agrega múltiples tópicos en lote a un curso específico.
// @Description El archivo CSV debe seguir el formato de la plantilla descargable y contener las columnas requeridas (topic_id, study_hours).
// @Description La operación procesa cada fila individualmente y reporta éxitos y errores por separado.
// @Tags Course
// @Accept multipart/form-data
// @Produce json
// @Param lang path string true "Código de idioma" default(es) Enums(en, es)
// @Param course_id formData int true "ID del curso"
// @Param file formData file true "Archivo CSV con datos de tópicos"
// @Success 200 {object} response.ApiResponse[response.BulkCreateCourseTopicResponse] "Resultado de la creación masiva"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (archivo CSV mal formateado)"
// @Failure 401 {object} response.ApiResponse[string] "Usuario no autenticado"
// @Failure 403 {object} response.ApiResponse[string] "No tiene permisos para realizar esta acción"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course/topic/bulk [post]
// @Security BearerAuth
func (c *BulkCreateCourseTopicController) BulkCreateCourseTopic(context *gin.Context) {
	var err error
	var bulkCreateRequest request.BulkCreateCourseTopicRequest
	if err = context.ShouldBind(&bulkCreateRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	// Parse course_id from form
	courseIDStr := context.PostForm("course_id")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": "Invalid course_id"})
		return
	}

	var bulkCreateCommand command.BulkCreateCourseTopicCommand
	bulkCreateCommand.File = bulkCreateRequest.File
	bulkCreateCommand.CourseID = uint(courseID)

	token := c.GetToken(context)
	lang := context.Param("lang")
	bulkCreateCommand.UserID = token.Sub
	bulkCreateCommand.Permissions = token.Permissions["courses"]

	result, _ := mediator.Send[command.BulkCreateCourseTopicCommand, response.ApiResponse[response.BulkCreateCourseTopicResponse]](bulkCreateCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}


