package controller

import (
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type GetUserStatisticsController struct{ Authenticable }

var instanceGetUserStatisticsController *GetUserStatisticsController = nil

func GetGetUserStatisticsController() *GetUserStatisticsController {
	if instanceGetUserStatisticsController == nil {
		instanceGetUserStatisticsController = &GetUserStatisticsController{}
	}
	return instanceGetUserStatisticsController
}

// GetUserStatistics obtiene las estadísticas del usuario autenticado según su rol
// @Summary Obtener Estadísticas del Usuario
// @Description Obtiene estadísticas personalizadas del usuario autenticado según su rol.
// @Description **Para Program Lead o Guest**: Programas Registrados, Reportes Generados, Asesorías Activas
// @Description **Para Cybersecurity Auditor**: Reportes Retroalimentados, Reportes Pendientes, Solicitudes de Asesoría
// @Description **Para Admin**: Usuarios Activos, Reportes Totales, Consultas Totales, Solicitudes de Acceso Totales
// @Tags Statistics
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[response.GetUserStatisticsResponse] "Estadísticas obtenidas exitosamente"
// @Failure 401 {object} response.ApiResponse[string] "Usuario no autenticado"
// @Failure 404 {object} response.ApiResponse[string] "Usuario no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/statistics [get]
// @Security BearerAuth
func (c *GetUserStatisticsController) GetUserStatistics(context *gin.Context) {
	token := c.GetToken(context)
	lang := context.Param("lang")

	getUserStatisticsQuery := query.GetUserStatisticsQuery{
		UserID: token.Sub,
	}

	result, _ := mediator.Send[query.GetUserStatisticsQuery, response.ApiResponse[response.GetUserStatisticsResponse]](getUserStatisticsQuery, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
