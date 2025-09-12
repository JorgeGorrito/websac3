package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type ListReportsByDegreeProgramController struct {
	Authenticable
}

var listReportsByDegreeProgramController *ListReportsByDegreeProgramController = nil

func GetListReportsByDegreeProgramController() *ListReportsByDegreeProgramController {
	if listReportsByDegreeProgramController == nil {
		listReportsByDegreeProgramController = &ListReportsByDegreeProgramController{}
	}
	return listReportsByDegreeProgramController
}

// Handle lista de reportes generados para un programa de grado específico
// @Summary Listar Reportes por Programa de Grado
// @Description Lista todos los reportes de evaluación generados para un programa de grado específico, ordenados por fecha de creación (más recientes primero).
// @Tags Report
// @Accept json
// @Produce json
// @Param degree_program_id path uint true "ID del programa de grado"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[[]response.ListReportResponse] "Se obtuvieron los reportes exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "ID de programa de grado inválido"
// @Failure 401 {object} response.ApiResponse[string] "Token de autenticación requerido o inválido"
// @Failure 403 {object} response.ApiResponse[string] "Sin permisos para acceder a este recurso"
// @Failure 404 {object} response.ApiResponse[string] "Programa de grado no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program/{degree_program_id}/reports [get]
// @Security BearerAuth
func (c *ListReportsByDegreeProgramController) Handle(ctx *gin.Context) {
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

	queryRequest := query.ListReportsByDegreeProgramQuery{
		DegreeProgramID: uint(degreeProgramID),
	}

	result, err := mediator.Send[query.ListReportsByDegreeProgramQuery, response.ApiResponse[[]response.ListReportResponse]](queryRequest, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				"Internal server error",
			},
		})
		return
	}

	ctx.JSON(result.HttpStatusCode, result)
}
