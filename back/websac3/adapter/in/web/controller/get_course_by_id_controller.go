package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type GetCourseByIDController struct{ Authenticable }

var getCourseByIDController *GetCourseByIDController = nil

func GetGetCourseByIDController() *GetCourseByIDController {
	if getCourseByIDController == nil {
		getCourseByIDController = &GetCourseByIDController{}
	}
	return getCourseByIDController
}

// Handle obtiene el detalle de un curso por ID
// @Summary Obtener Detalle de Curso
// @Description Obtiene la información detallada de un curso específico incluyendo sus tópicos asociados.
// @Tags Course
// @Accept json
// @Produce json
// @Param course_id path int true "ID del curso"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[response.GetCourseByIDResponse] "Se obtuvo el detalle del curso exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 404 {object} response.ApiResponse[string] "Curso no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course/{course_id} [get]
// @Security BearerAuth
func (c *GetCourseByIDController) Handle(ctx *gin.Context) {
	// Get course ID from path parameter
	courseIDStr := ctx.Param("course_id")
	courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	token := c.GetToken(ctx)
	lang := ctx.Param("lang")
	var requestQuery = query.GetCourseByIDQuery{
		CourseID:    uint(courseID),
		Permissions: token.Permissions["courses"],
	}

	result, _ := mediator.Send[query.GetCourseByIDQuery, response.ApiResponse[response.GetCourseByIDResponse]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
