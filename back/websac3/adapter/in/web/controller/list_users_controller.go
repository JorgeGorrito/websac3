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

type ListUsersController struct{ Authenticable }

var listUsersController *ListUsersController = nil

func GetListUsersController() *ListUsersController {
	if listUsersController == nil {
		listUsersController = &ListUsersController{}
	}
	return listUsersController
}

// Handle lista usuarios con paginación y filtros
// @Summary Listar Usuarios
// @Description Lista los usuarios registrados según filtros dinámicos y parámetros de paginación.
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Por ejemplo: `email[cont]=juan`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[eq]`: igual a (exact match)
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos:**
// @Description - `email[eq|cont]`: filtrar por email
// @Description - `Person.name[eq|cont]`: filtrar por nombre de la persona
// @Description - `Person.lastname[eq|cont]`: filtrar por apellido de la persona
// @Description - `Role.name[eq|cont]`: filtrar por nombre del rol
// @Description - `is_active[eq]`: filtrar por estado activo (true/false)
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags User
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: email[cont]=juan)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListUsersResponse]] "Se obtuvieron los usuarios exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "No se encontraron usuarios"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/users [get]
// @Security BearerAuth
func (c *ListUsersController) Handle(ctx *gin.Context) {
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

	token := c.GetToken(ctx)

	permissions := c.GetPermissionsByModuleName(token.Permissions, "users")
	lang := ctx.Param("lang")
	var filters = util.ParseParamsFilter(filtersMap)
	var requestQuery = query.ListUsersQuery{
		UserID:           token.Sub,
		PaginationParams: paginationParams,
		Filters:          filters,
		Permissions:      permissions,
	}

	result, _ := mediator.Send[query.ListUsersQuery, response.ListUsersApiResponse](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result)
}
