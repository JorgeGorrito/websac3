package controller

import (
	"fmt"
	"net/http"

	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mapper"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type ActivateUserController struct{ Authenticable }

var activateUserController *ActivateUserController = nil

func GetActivateUserController() *ActivateUserController {
	if activateUserController == nil {
		activateUserController = &ActivateUserController{}
	}
	return activateUserController
}

// Handle activa un usuario por su ID
// @Summary Activar Usuario
// @Description Activa un usuario del sistema por su ID.
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path uint true "ID del usuario a activar"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[string] "Usuario activado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 403 {object} response.ApiResponse[string] "No tiene permisos para realizar esta acción"
// @Failure 404 {object} response.ApiResponse[string] "Usuario no encontrado"
// @Failure 409 {object} response.ApiResponse[string] "Usuario ya está activo"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/users/{user_id}/activate [put]
// @Security BearerAuth
func (c *ActivateUserController) Handle(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	if userIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var userID uint
	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	permissions := c.GetPermissionsByModuleName(token.Permissions, "user-management")
	lang := ctx.Param("lang")

	input := &mapper.UserActionInput{
		UserID:      userID,
		Permissions: permissions,
	}
	cmd, err := mapper.Map[mapper.UserActionInput, command.ActivateUserCommand](input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating command"})
		return
	}

	result, _ := mediator.Send[command.ActivateUserCommand, response.ApiResponse[string]](cmd, lang)
	ctx.JSON(result.HttpStatusCode, result)
}
