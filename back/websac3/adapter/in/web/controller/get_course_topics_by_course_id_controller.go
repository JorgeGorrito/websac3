package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type GetCourseTopicsByCourseIDController struct{ Authenticable }

var getCourseTopicsByCourseIDController *GetCourseTopicsByCourseIDController = nil

func GetGetCourseTopicsByCourseIDController() *GetCourseTopicsByCourseIDController {
	if getCourseTopicsByCourseIDController == nil {
		getCourseTopicsByCourseIDController = &GetCourseTopicsByCourseIDController{}
	}
	return getCourseTopicsByCourseIDController
}

// Handle obtiene los tópicos de un curso por ID
// @Summary Obtener Tópicos de Curso
// @Description Obtiene la lista de tópicos asociados a un curso específico con información detallada.
// @Tags Course
// @Accept json
// @Produce json
// @Param course_id path int true "ID del curso"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[response.GetCourseTopicsByCourseIDResponse] "Se obtuvieron los tópicos del curso exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 404 {object} response.ApiResponse[string] "Tópicos no encontrados para el curso"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course/{course_id}/topics [get]
// @Security BearerAuth
func (c *GetCourseTopicsByCourseIDController) Handle(ctx *gin.Context) {
	// Get course ID from path parameter
	courseIDStr := ctx.Param("course_id")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	token := c.GetToken(ctx)
	lang := ctx.Param("lang")
	var requestQuery = query.GetCourseTopicsByCourseIDQuery{
		CourseID:    uint(courseID),
		Permissions: token.Permissions["courses"],
	}

	result, _ := mediator.Send[query.GetCourseTopicsByCourseIDQuery, response.ApiResponse[response.GetCourseTopicsByCourseIDResponse]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
