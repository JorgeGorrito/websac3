package controller

import (
	"net/http"
	"net/url"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"
	"websac3/common/paginator"

	"github.com/gin-gonic/gin"
)

type ListRejectedAccessRequestController struct{ Authenticable }

var listRejectedAccessRequestControllerInstance *ListRejectedAccessRequestController = nil

func GetListRejectedAccessRequestController() *ListRejectedAccessRequestController {
	if listRejectedAccessRequestControllerInstance == nil {
		listRejectedAccessRequestControllerInstance = &ListRejectedAccessRequestController{}
	}
	return listRejectedAccessRequestControllerInstance
}

// Handle lista solicitudes de acceso rechazadas con paginación y filtros
// @Summary Listar Solicitudes de Acceso Rechazadas
// @Description Lista las solicitudes de acceso rechazadas existentes según filtros dinámicos y parámetros de paginación.
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Por ejemplo: `Applicant.name[cont]=juan`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[eq]`: igual a (exact match)
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos:**
// @Description - `Applicant.name[eq|cont]`
// @Description - `Applicant.lastname[eq|cont]`
// @Description - `Applicant.HigherEducationInstitution.name[eq|cont]`
// @Description - `Applicant.HigherEducationInstitution.Municipality.name[eq|cont]`
// @Description - `Applicant.HigherEducationInstitution.Department.name[eq|cont]`
// @Description - `Status.name[eq|cont]`
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags AccessRequest
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: Applicant.name[cont]=juan)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListRejectedAccessRequestResponse]] "Se obtuvieron las solicitudes de acceso rechazadas exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 401 {object} response.ApiResponse[string] "Token de autenticación inválido o faltante"
// @Failure 403 {object} response.ApiResponse[string] "No tiene permisos para realizar esta acción"
// @Failure 404 {object} response.ApiResponse[string] "No se encontraron solicitudes de acceso rechazadas"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/access-requests/rejected [get]
// @Security BearerAuth
func (c *ListRejectedAccessRequestController) Handle(ctx *gin.Context) {
	var err error
	var paginationParams paginator.PaginationParams

	if err = ctx.ShouldBindQuery(&paginationParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	rawFilters := ctx.Request.URL.Query().Get("filters")
	filtersMap, err := url.ParseQuery(rawFilters)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filters format"})
		return
	}
	var filters = util.ParseParamsFilter(filtersMap)
	var lang string = ctx.Param("lang")

	// Obtener el usuario autenticado
	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Convertir permisos de map a slice
	var permissions []string
	for _, perms := range token.Permissions {
		permissions = append(permissions, perms...)
	}

	var requestQuery = query.ListRejectedAccessRequestQuery{
		PaginationParams: paginationParams,
		Filters:          filters,
		UserID:           token.Sub,
		Permissions:      permissions,
	}

	result, _ := mediator.Send[query.ListRejectedAccessRequestQuery, response.ApiResponse[paginator.Page[response.ListRejectedAccessRequestResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
