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

type UpdateCourseController struct{ Authenticable }

var updateCourseController *UpdateCourseController = nil

func GetUpdateCourseController() *UpdateCourseController {
	if updateCourseController == nil {
		updateCourseController = &UpdateCourseController{}
	}
	return updateCourseController
}

// Handle actualiza un curso
// @Summary Actualizar Curso
// @Description Actualiza un curso específico. Solo el usuario que creó el curso puede actualizarlo.
// @Tags Course
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param course_id path int true "ID del curso a actualizar"
// @Param request body request.UpdateCourseRequest true "Datos del curso a actualizar"
// @Success 200 {object} response.ApiResponse[string] "Curso actualizado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 403 {object} response.ApiResponse[string] "No autorizado para actualizar este curso"
// @Failure 404 {object} response.ApiResponse[string] "Curso no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course/{course_id} [put]
// @Security BearerAuth
func (c *UpdateCourseController) Handle(ctx *gin.Context) {
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
	var req request.UpdateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = uint(courseID)

	// Map request to command using mappers.Map
	var cmd command.UpdateCourseCommand
	cmd, err = mapper.Map[request.UpdateCourseRequest, command.UpdateCourseCommand](&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add additional fields from token
	cmd.UserID = token.Sub
	cmd.Permissions = token.Permissions["courses"]

	result, _ := mediator.Send[command.UpdateCourseCommand, response.ApiResponse[string]](cmd, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
