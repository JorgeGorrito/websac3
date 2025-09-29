package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type RejectExpertConsultationController struct {
	Authenticable
}

func NewRejectExpertConsultationController() *RejectExpertConsultationController {
	return &RejectExpertConsultationController{}
}

// RejectExpertConsultation godoc
// @Summary Rechazar solicitud de asesoría
// @Description Permite a un experto rechazar una solicitud de asesoría con un mensaje de respuesta
// @Tags ExpertConsultation
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param consultation_id path int true "ID de la solicitud de asesoría"
// @Param request body request.RejectExpertConsultationRequest true "Datos para rechazar la solicitud"
// @Success 200 {object} response.ApiResponse[string] "Solicitud rechazada exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Error de validación"
// @Failure 401 {object} response.ApiResponse[string] "No autorizado"
// @Failure 403 {object} response.ApiResponse[string] "Permisos insuficientes"
// @Failure 409 {object} response.ApiResponse[string] "Conflicto - solicitud ya procesada"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/expert-consultations/{consultation_id}/reject [put]
// @Security BearerAuth
func (c *RejectExpertConsultationController) Handle(ctx *gin.Context) {
	// Obtener el ID de la consulta desde la URL
	consultationIDStr := ctx.Param("consultation_id")
	consultationID, err := strconv.ParseUint(consultationIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de consulta inválido"})
		return
	}

	// Obtener el usuario autenticado
	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Parsear el request body
	var rejectRequest request.RejectExpertConsultationRequest
	if err := ctx.ShouldBindJSON(&rejectRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de request inválido"})
		return
	}

	// Mapear request a command
	rejectCommand := command.RejectExpertConsultationCommand{
		ConsultationID: uint(consultationID),
		ExpertResponse: rejectRequest.ExpertResponse,
		ExpertID:       token.Sub,
	}

	result, err := mediator.Send[command.RejectExpertConsultationCommand, response.ApiResponse[string]](rejectCommand, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
