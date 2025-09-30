package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type GetDegreeProgramByIDController struct{}

var getDegreeProgramByIDControllerInstance *GetDegreeProgramByIDController = nil

func GetGetDegreeProgramByIDController() *GetDegreeProgramByIDController {
	if getDegreeProgramByIDControllerInstance == nil {
		getDegreeProgramByIDControllerInstance = &GetDegreeProgramByIDController{}
	}
	return getDegreeProgramByIDControllerInstance
}

// Handle obtiene el detalle de un programa de grado por ID
// @Summary Obtener Detalle de Programa de Grado
// @Description Obtiene la información detallada de un programa de grado específico incluyendo roles profesionales, institución y todos los datos relacionados. Para consultar los cursos del programa, use el endpoint específico de cursos.
// @Tags DegreeProgram
// @Accept json
// @Produce json
// @Param degree_program_id path int true "ID del programa de grado"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[response.GetDegreeProgramByIDResponse] "Se obtuvo el detalle del programa de grado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "ID de programa de grado inválido"
// @Failure 404 {object} response.ApiResponse[string] "Programa de grado no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program/{degree_program_id} [get]
func (c *GetDegreeProgramByIDController) Handle(ctx *gin.Context) {
	lang := ctx.Param("lang")
	degreeProgramIDStr := ctx.Param("degree_program_id")

	// Validar que el ID sea un número válido
	degreeProgramID, err := strconv.ParseUint(degreeProgramIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors: []string{
				"Invalid degree program ID format",
			},
		})
		return
	}

	queryRequest := query.GetDegreeProgramByIDQuery{
		DegreeProgramID: uint(degreeProgramID),
	}

	result, err := mediator.Send[query.GetDegreeProgramByIDQuery, response.ApiResponse[response.GetDegreeProgramByIDResponse]](queryRequest, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				"Internal server error",
			},
		})
		return
	}

	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
