package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type GetExpertConsultationByIDController struct {
	Authenticable
}

func NewGetExpertConsultationByIDController() *GetExpertConsultationByIDController {
	return &GetExpertConsultationByIDController{}
}

// GetExpertConsultationByID godoc
// @Summary Obtener asesoría de experto por ID
// @Description Permite obtener los detalles de una asesoría de experto específica por su ID
// @Tags ExpertConsultation
// @Accept json
// @Produce json
// @Param consultation_id path int true "ID de la asesoría"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[response.UserExpertConsultationResponse] "Detalles de la asesoría"
// @Failure 400 {object} response.ApiResponse[response.UserExpertConsultationResponse] "Error de validación"
// @Failure 401 {object} response.ApiResponse[response.UserExpertConsultationResponse] "No autorizado"
// @Failure 403 {object} response.ApiResponse[response.UserExpertConsultationResponse] "Permisos insuficientes"
// @Failure 404 {object} response.ApiResponse[response.UserExpertConsultationResponse] "Asesoría no encontrada"
// @Failure 500 {object} response.ApiResponse[response.UserExpertConsultationResponse] "Error interno del servidor"
// @Router /api/v1/{lang}/expert-consultations/{consultation_id} [get]
// @Security BearerAuth
func (c *GetExpertConsultationByIDController) Handle(ctx *gin.Context) {
	// Obtener consultation_id del path
	consultationIDStr := ctx.Param("consultation_id")
	consultationID, err := strconv.ParseUint(consultationIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid consultation ID format"})
		return
	}

	// Obtener el usuario autenticado
	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	lang := ctx.Param("lang")

	// Crear el query
	queryDto := query.GetExpertConsultationByIDQuery{
		ConsultationID: uint(consultationID),
		UserID:         token.Sub,
		Permissions:    token.Permissions["expert-advisory"],
	}

	result, err := mediator.Send[query.GetExpertConsultationByIDQuery, response.ApiResponse[response.UserExpertConsultationResponse]](queryDto, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
