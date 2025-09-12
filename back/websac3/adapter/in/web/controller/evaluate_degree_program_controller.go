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

type EvaluateDegreeProgramController struct {
	Authenticable
}

var evaluateDegreeProgramControllerInstance *EvaluateDegreeProgramController = nil

func GetEvaluateDegreeProgramController() *EvaluateDegreeProgramController {
	if evaluateDegreeProgramControllerInstance == nil {
		evaluateDegreeProgramControllerInstance = &EvaluateDegreeProgramController{}
	}
	return evaluateDegreeProgramControllerInstance
}

// @Summary Evaluar programa de grado
// @Tags DegreeProgram
// @Accept json
// @Produce json
// @Param request body request.EvaluateDegreeProgramRequest true "Datos del programa de grado y rol profesional a evaluar"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[string]
// @Failure 400 {object} response.ApiResponse[string]
// @Router /api/v1/{lang}/degree-program/evaluate [post]
// @Security BearerAuth
func (c *EvaluateDegreeProgramController) Handle(ctx *gin.Context) {
	var err error
	var req request.EvaluateDegreeProgramRequest
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

	var cmd command.EvaluateDegreeProgramCommand
	cmd, err = mapper.Map[request.EvaluateDegreeProgramRequest, command.EvaluateDegreeProgramCommand](&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Agregar el UserID del token al comando
	cmd.UserID = token.Sub

	result, err := mediator.Send[command.EvaluateDegreeProgramCommand, response.ApiResponse[string]](cmd, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
