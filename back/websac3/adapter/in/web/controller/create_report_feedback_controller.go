package controller

import (
	"net/http"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mapper"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type CreateReportFeedbackController struct {
	Authenticable
}

var instanceCreateReportFeedbackController *CreateReportFeedbackController = nil

func GetCreateReportFeedbackController() *CreateReportFeedbackController {
	if instanceCreateReportFeedbackController == nil {
		instanceCreateReportFeedbackController = &CreateReportFeedbackController{}
	}
	return instanceCreateReportFeedbackController
}

// CreateReportFeedback godoc
// @Summary Crear retroalimentación de reporte
// @Description Permite a un auditor de ciberseguridad crear retroalimentación para un reporte de evaluación
// @Tags Report Feedback
// @Accept json
// @Produce json
// @Param request body request.CreateReportFeedbackRequest true "Datos de la retroalimentación"
// @Param lang path string true "Código de idioma" default(en) English Enums(en, es)
// @Success 201 {object} response.ApiResponse[string] "Retroalimentación creada exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Error de validación"
// @Failure 401 {object} response.ApiResponse[string] "No autorizado"
// @Failure 403 {object} response.ApiResponse[string] "Permisos insuficientes"
// @Failure 409 {object} response.ApiResponse[string] "Ya existe retroalimentación para este reporte"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/report-feedbacks [post]
// @Security BearerAuth
func (c *CreateReportFeedbackController) Handle(ctx *gin.Context) {
	var err error
	var req request.CreateReportFeedbackRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener el usuario autenticado
	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	var cmd command.CreateReportFeedbackCommand
	cmd, err = mapper.Map[request.CreateReportFeedbackRequest, command.CreateReportFeedbackCommand](&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Agregar el UserID del token al comando
	cmd.UserID = token.Sub

	result, err := mediator.Send[command.CreateReportFeedbackCommand, response.ApiResponse[string]](cmd, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
