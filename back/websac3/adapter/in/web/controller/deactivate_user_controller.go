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

type DeactivateUserController struct{ Authenticable }

var deactivateUserController *DeactivateUserController = nil

func GetDeactivateUserController() *DeactivateUserController {
	if deactivateUserController == nil {
		deactivateUserController = &DeactivateUserController{}
	}
	return deactivateUserController
}

// Handle desactiva un usuario por su ID
// @Summary Desactivar Usuario
// @Description Desactiva un usuario del sistema por su ID.
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path uint true "ID del usuario a desactivar"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[string] "Usuario desactivado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 403 {object} response.ApiResponse[string] "No tiene permisos para realizar esta acción"
// @Failure 404 {object} response.ApiResponse[string] "Usuario no encontrado"
// @Failure 409 {object} response.ApiResponse[string] "Usuario ya está desactivado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/users/{user_id}/deactivate [put]
// @Security BearerAuth
func (c *DeactivateUserController) Handle(ctx *gin.Context) {
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
	cmd, err := mapper.Map[mapper.UserActionInput, command.DeactivateUserCommand](input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating command: " + err.Error()})
		return
	}

	result, _ := mediator.Send[command.DeactivateUserCommand, response.ApiResponse[string]](cmd, lang)
	ctx.JSON(result.HttpStatusCode, result)
}
