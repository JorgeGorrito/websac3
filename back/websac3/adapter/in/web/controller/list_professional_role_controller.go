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

type ListProfessionalRoleController struct{}

var listProfessionalRoleControllerInstance *ListProfessionalRoleController = nil

func GetListProfessionalRoleController() *ListProfessionalRoleController {
	if listProfessionalRoleControllerInstance == nil {
		listProfessionalRoleControllerInstance = &ListProfessionalRoleController{}
	}
	return listProfessionalRoleControllerInstance
}

// Handle lista los roles profesionales con paginación y filtros
// @Summary Listar Roles Profesionales
// @Description Lista los roles profesionales existentes según filtros dinámicos y parámetros de paginación.
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Por ejemplo: `name[cont]=developer`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[eq]`: igual a (exact match)
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos:**
// @Description - `name[eq|cont]`
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags ProfessionalRole
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: name[cont]=developer)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListProfessionalRoleResponse]] "Se obtuvieron los roles profesionales exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "No se encontraron roles profesionales"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/professional-roles [get]
func (c *ListProfessionalRoleController) Handle(ctx *gin.Context) {
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
	var requestQuery = query.ListProfessionalRoleQuery{
		PaginationParams: paginationParams,
		Filters:          filters,
	}

	result, _ := mediator.Send[query.ListProfessionalRoleQuery, response.ApiResponse[paginator.Page[response.ListProfessionalRoleResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
