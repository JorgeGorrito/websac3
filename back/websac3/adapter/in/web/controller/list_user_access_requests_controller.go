package controller

import (
	"net/http"
	"net/url"
	"strings"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"
	"websac3/common/paginator"

	"github.com/gin-gonic/gin"
)

type ListUserAccessRequestsController struct {
	Authenticable
}

func NewListUserAccessRequestsController() *ListUserAccessRequestsController {
	return &ListUserAccessRequestsController{}
}

// ListUserAccessRequests godoc
// @Summary Listar solicitudes de acceso del usuario autenticado
// @Description Permite a un usuario autenticado obtener la lista de sus propias solicitudes de acceso
// @Tags Access Request
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param Status.name[eq] query string false "Filtro por estado exacto"
// @Param Status.name[cont] query string false "Filtro por estado que contiene"
// @Param Applicant.HigherEducationInstitution.name[eq] query string false "Filtro por nombre exacto de institución educativa"
// @Param Applicant.HigherEducationInstitution.name[cont] query string false "Filtro por nombre de institución educativa que contiene"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]] "Lista de solicitudes de acceso del usuario"
// @Failure 400 {object} response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]] "Error de validación"
// @Failure 401 {object} response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]] "No autorizado"
// @Failure 403 {object} response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]] "Permisos insuficientes"
// @Failure 500 {object} response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]] "Error interno del servidor"
// @Router /api/v1/{lang}/user/access-requests [get]
// @Security BearerAuth
func (c *ListUserAccessRequestsController) Handle(ctx *gin.Context) {
	var err error
	var paginationParams paginator.PaginationParams

	if err = ctx.ShouldBindQuery(&paginationParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Obtener el usuario autenticado
	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Parsear filtros de la query string
	allParams := ctx.Request.URL.Query()
	filtersMap := make(url.Values)

	// Filtrar solo los parámetros que contienen operadores (filtros)
	for key, values := range allParams {
		if strings.Contains(key, "[") && strings.HasSuffix(key, "]") {
			filtersMap[key] = values
		}
	}
	var filters = util.ParseParamsFilter(filtersMap)

	// Crear el query con filtros
	queryDto := query.ListUserAccessRequestsQuery{
		PaginationParams: paginationParams,
		UserID:           token.Sub,
		Filters:          filters,
		Permissions:      token.Permissions["access-request"],
	}

	result, err := mediator.Send[query.ListUserAccessRequestsQuery, response.ApiResponse[paginator.Page[response.UserAccessRequestResponse]]](queryDto, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
