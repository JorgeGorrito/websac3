package controller

import (
	"strconv"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type ApproveAccessRequestController struct{ Authenticable }

var approveAccessRequestControllerInstance *ApproveAccessRequestController = nil

func GetApproveAccessRequestController() *ApproveAccessRequestController {
	if approveAccessRequestControllerInstance == nil {
		approveAccessRequestControllerInstance = &ApproveAccessRequestController{}
	}
	return approveAccessRequestControllerInstance
}

// Handle aprueba una solicitud de acceso
// @Summary Aprobar Solicitud de Acceso
// @Description Aprueba una solicitud de acceso previamente registrada.
// @Description Esta acción activa el proceso para la creación de usuario por parte del solicitante.
// @Tags AccessRequest
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param accessRequestID path int true "ID de la solicitud de acceso a aprobar"
// @Param request body request.ApproveAccessRequestRequest true "Cuerpo de la solicitud con RoleID"
// @Success 200 {object} response.ApiResponse[string] "La solicitud fue aprobada correctamente"
// @Failure 400 {object} response.ApiResponse[string] "ID inválido o formato incorrecto"
// @Failure 404 {object} response.ApiResponse[string] "No se encontró la solicitud de acceso"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/access-request/{accessRequestID}/approve [post]
// @Security BearerAuth
func (c *ApproveAccessRequestController) Handle(ctx *gin.Context) {
	var AccessRequestToApproveID uint

	idStr := ctx.Param("accessRequestID")
	if idStr == "" {
		idStr = "0"
	}

	AccessRequestToApproveID, _ = func() (uint, error) {
		ID, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(ID), nil
	}()

	// Obtener el RoleID del body
	var requestBody request.ApproveAccessRequestRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(400, response.ApiResponse[string]{
			HttpStatusCode: 400,
			Errors:         []string{"Invalid request body"},
		})
		return
	}

	token := c.GetToken(ctx)
	lang := ctx.Param("lang")
	var request command.ApproveAccessRequestCommand = command.ApproveAccessRequestCommand{
		AccessRequestID: AccessRequestToApproveID,
		UserID:          token.Sub,
		RoleID:          requestBody.RoleID,
		Permissions:     token.Permissions["access-request"],
	}

	result, _ := mediator.Send[command.ApproveAccessRequestCommand, response.ApiResponse[string]](request, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
