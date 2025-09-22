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

type DeleteCourseController struct{ Authenticable }

var deleteCourseController *DeleteCourseController = nil

func GetDeleteCourseController() *DeleteCourseController {
	if deleteCourseController == nil {
		deleteCourseController = &DeleteCourseController{}
	}
	return deleteCourseController
}

// Handle elimina un curso
// @Summary Eliminar Curso
// @Description Elimina un curso específico. Solo el usuario que creó el curso puede eliminarlo.
// @Tags Course
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param course_id path int true "ID del curso a eliminar"
// @Success 200 {object} response.ApiResponse[string] "Curso eliminado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 403 {object} response.ApiResponse[string] "No autorizado para eliminar este curso"
// @Failure 404 {object} response.ApiResponse[string] "Curso no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course/{course_id} [delete]
// @Security BearerAuth
func (c *DeleteCourseController) Handle(ctx *gin.Context) {
	token := c.GetToken(ctx)
	lang := ctx.Param("lang")

	// Get course ID from URL parameter
	courseIDStr := ctx.Param("course_id")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Create request DTO
	var req request.DeleteCourseRequest
	req.CourseID = uint(courseID)

	// Map request to command using mappers.Map
	var cmd command.DeleteCourseCommand
	cmd, err = mapper.Map[request.DeleteCourseRequest, command.DeleteCourseCommand](&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add additional fields from token
	cmd.UserID = token.Sub
	cmd.Permissions = token.Permissions["courses"]

	result, _ := mediator.Send[command.DeleteCourseCommand, response.ApiResponse[string]](cmd, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
