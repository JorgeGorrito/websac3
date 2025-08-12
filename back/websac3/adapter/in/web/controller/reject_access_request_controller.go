package controller

import (
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type RejectAccessRequestController struct{ Authenticable }

var rejectAccessRequestControllerInstance *RejectAccessRequestController = nil

func GetRejectAccessRequestController() *RejectAccessRequestController {
	if rejectAccessRequestControllerInstance == nil {
		rejectAccessRequestControllerInstance = &RejectAccessRequestController{}
	}
	return rejectAccessRequestControllerInstance
}

// Handle rechaza una solicitud de acceso
// @Summary Rechazar Solicitud de Acceso
// @Description Rechaza una solicitud de acceso previamente registrada.
// @Description Esta acción finaliza el proceso, impidiendo la creación de usuario para el solicitante.
// @Tags AccessRequest
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param accessRequestID path int true "ID de la solicitud de acceso a rechazar"
// @Success 200 {object} response.ApiResponse[string] "La solicitud fue rechazada correctamente"
// @Failure 400 {object} response.ApiResponse[string] "ID inválido o formato incorrecto"
// @Failure 404 {object} response.ApiResponse[string] "No se encontró la solicitud de acceso"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/access-request/{accessRequestID}/reject [post]
// @Security BearerAuth
func (c *RejectAccessRequestController) Handle(ctx *gin.Context) {
	var AccessRequestToRejectID uint

	idStr := ctx.Param("accessRequestID")
	if idStr == "" {
		idStr = "0"
	}

	AccessRequestToRejectID, _ = func() (uint, error) {
		ID, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(ID), nil
	}()

	token := c.GetToken(ctx)
	lang := ctx.Param("lang")
	var request command.RejectAccessRequestCommand = command.RejectAccessRequestCommand{
		AccessRequestID: AccessRequestToRejectID,
		UserID:          token.Sub,
		Permissions:     token.Permissions["access-request"],
	}

	result, _ := mediator.Send[command.RejectAccessRequestCommand, response.ApiResponse[string]](request, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
