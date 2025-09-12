package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type GetReportByIDController struct {
	Authenticable
}

var getReportByIDController *GetReportByIDController = nil

func GetGetReportByIDController() *GetReportByIDController {
	if getReportByIDController == nil {
		getReportByIDController = &GetReportByIDController{}
	}
	return getReportByIDController
}

// Handle obtiene un reporte específico por su ID
// @Summary Obtener Reporte por ID
// @Description Obtiene toda la información detallada de un reporte específico por su ID, incluyendo programa de grado, rol profesional y reportes de áreas de conocimiento.
// @Tags Report
// @Accept json
// @Produce json
// @Param report_id path uint true "ID del reporte"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[response.ListReportResponse] "Se obtuvo el reporte exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "ID de reporte inválido"
// @Failure 401 {object} response.ApiResponse[string] "Token de autenticación requerido o inválido"
// @Failure 403 {object} response.ApiResponse[string] "Sin permisos para acceder a este recurso"
// @Failure 404 {object} response.ApiResponse[string] "Reporte no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/report/{report_id} [get]
// @Security BearerAuth
func (c *GetReportByIDController) Handle(ctx *gin.Context) {
	lang := ctx.Param("lang")
	reportIDStr := ctx.Param("report_id")

	// Validar que el ID sea un número válido
	reportID, err := strconv.ParseUint(reportIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors: []string{
				"Invalid report ID format",
			},
		})
		return
	}

	queryRequest := query.GetReportByIDQuery{
		ReportID: uint(reportID),
	}

	result, err := mediator.Send[query.GetReportByIDQuery, response.ApiResponse[response.ListReportResponse]](queryRequest, lang)
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
