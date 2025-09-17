package controller

import (
	"net/http"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type ListReportsPendingFeedbackController struct {
	Authenticable
}

func NewListReportsPendingFeedbackController() *ListReportsPendingFeedbackController {
	return &ListReportsPendingFeedbackController{}
}

// ListReportsPendingFeedback godoc
// @Summary Listar reportes pendientes de retroalimentación
// @Description Permite a un auditor de ciberseguridad obtener la lista de reportes que no tienen retroalimentación
// @Tags Report Feedback
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[response.PaginatedResponse[response.ReportResponse]] "Lista de reportes pendientes"
// @Failure 400 {object} response.ApiResponse[response.PaginatedResponse[response.ReportResponse]] "Error de validación"
// @Failure 401 {object} response.ApiResponse[response.PaginatedResponse[response.ReportResponse]] "No autorizado"
// @Failure 403 {object} response.ApiResponse[response.PaginatedResponse[response.ReportResponse]] "Permisos insuficientes"
// @Failure 500 {object} response.ApiResponse[response.PaginatedResponse[response.ReportResponse]] "Error interno del servidor"
// @Router /api/v1/{lang}/report-feedbacks/pending-reports [get]
// @Security BearerAuth
func (c *ListReportsPendingFeedbackController) Handle(ctx *gin.Context) {
	var err error
	var req request.ListReportsPendingFeedbackRequest
	if err = ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Obtener el usuario autenticado
	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Valores por defecto
	if req.Currentpage == 0 {
		req.Currentpage = 1
	}
	if req.ItemsPerpage == 0 {
		req.ItemsPerpage = 10
	}

	// Crear el query directamente
	queryDto := query.ListReportsPendingFeedbackQuery{
		PaginationParams: req.PaginationParams,
		UserID:           token.Sub,
		Permissions:      token.Permissions["reporting-feedback"],
	}

	result, err := mediator.Send[query.ListReportsPendingFeedbackQuery, response.ApiResponse[response.PaginatedResponse[response.ReportResponse]]](queryDto, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
