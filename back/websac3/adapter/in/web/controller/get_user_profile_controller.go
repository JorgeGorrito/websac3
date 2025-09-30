package controller

import (
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type GetUserProfileController struct{ Authenticable }

var instanceGetUserProfileController *GetUserProfileController = nil

func GetGetUserProfileController() *GetUserProfileController {
	if instanceGetUserProfileController == nil {
		instanceGetUserProfileController = &GetUserProfileController{}
	}
	return instanceGetUserProfileController
}

// GetUserProfile obtiene la información del perfil del usuario autenticado
// @Summary Obtener Perfil del Usuario
// @Description Obtiene la información del perfil del usuario autenticado actualmente.
// @Description Devuelve información completa del usuario incluyendo datos personales, rol, institución, etc.
// @Tags User
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[response.GetUserProfileResponse] "Perfil del usuario obtenido exitosamente"
// @Failure 401 {object} response.ApiResponse[string] "Usuario no autenticado"
// @Failure 404 {object} response.ApiResponse[string] "Usuario no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/profile [get]
// @Security BearerAuth
func (c *GetUserProfileController) GetUserProfile(context *gin.Context) {
	token := c.GetToken(context)
	lang := context.Param("lang")

	getUserProfileQuery := query.GetUserProfileQuery{
		UserID: token.Sub,
	}

	result, _ := mediator.Send[query.GetUserProfileQuery, response.ApiResponse[response.GetUserProfileResponse]](getUserProfileQuery, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
