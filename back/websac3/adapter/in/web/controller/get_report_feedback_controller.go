package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type GetReportFeedbackController struct {
	Authenticable
}

func NewGetReportFeedbackController() *GetReportFeedbackController {
	return &GetReportFeedbackController{}
}

// GetReportFeedback godoc
// @Summary Obtener retroalimentación de un reporte
// @Description Obtiene la retroalimentación completa de un reporte específico
// @Tags Report Feedback
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param reportId path int true "ID del reporte"
// @Success 200 {object} response.ApiResponse[response.ReportFeedbackResponse] "Retroalimentación del reporte"
// @Failure 400 {object} response.ApiResponse[response.ReportFeedbackResponse] "Error de validación"
// @Failure 401 {object} response.ApiResponse[response.ReportFeedbackResponse] "No autorizado"
// @Failure 404 {object} response.ApiResponse[response.ReportFeedbackResponse] "Retroalimentación no encontrada"
// @Failure 500 {object} response.ApiResponse[response.ReportFeedbackResponse] "Error interno del servidor"
// @Router /api/v1/{lang}/report-feedbacks/report/{reportId} [get]
// @Security BearerAuth
func (c *GetReportFeedbackController) Handle(ctx *gin.Context) {
	// Obtener el ID del reporte de los parámetros de la URL
	reportIDStr := ctx.Param("reportId")
	reportID, err := strconv.ParseUint(reportIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de reporte inválido"})
		return
	}

	// Obtener el usuario autenticado
	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	queryDto := query.GetReportFeedbackQuery{
		ReportID:    uint(reportID),
		Permissions: token.Permissions["reporting-feedback"],
	}

	result, err := mediator.Send[query.GetReportFeedbackQuery, response.ApiResponse[response.ReportFeedbackResponse]](queryDto, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
