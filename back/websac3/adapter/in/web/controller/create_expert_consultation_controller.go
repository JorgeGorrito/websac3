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

type CreateExpertConsultationController struct{}

var instanceCreateExpertConsultationController *CreateExpertConsultationController = nil

func GetCreateExpertConsultationController() *CreateExpertConsultationController {
	if instanceCreateExpertConsultationController == nil {
		instanceCreateExpertConsultationController = &CreateExpertConsultationController{}
	}
	return instanceCreateExpertConsultationController
}

// CreateExpertConsultation manejador para crear una consulta de experto
// @Summary Crear Consulta de Experto
// @Description Crea una nueva consulta de asesoría con un experto en ciberseguridad para un programa de grado y reporte específico.
// @Description La consulta se crea en estado "pendiente" y será evaluada por un experto.
// @Tags ExpertConsultation
// @Accept json
// @Produce json
// @Param request body request.CreateExpertConsultationRequest true "CreateExpertConsultationRequest"
// @Param lang path string true "Código de idioma" default(en) English Enums(en, es)
// @Success 201 {object} response.ApiResponse[string] "Consulta de experto creada exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "Usuario, programa de grado o reporte no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/expert-consultation [post]
func (c *CreateExpertConsultationController) CreateExpertConsultation(context *gin.Context) {
	var err error
	var createExpertConsultationRequest request.CreateExpertConsultationRequest
	if err = context.ShouldBindJSON(&createExpertConsultationRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	var createExpertConsultationCommand command.CreateExpertConsultationCommand
	createExpertConsultationCommand, err = mapper.Map[request.CreateExpertConsultationRequest, command.CreateExpertConsultationCommand](&createExpertConsultationRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	lang := context.Param("lang")

	result, _ := mediator.Send[command.CreateExpertConsultationCommand, response.ApiResponse[string]](createExpertConsultationCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
