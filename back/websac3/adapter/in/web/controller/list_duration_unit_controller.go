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

type ListDurationUnitController struct{ Authenticable }

var listDurationUnitController *ListDurationUnitController = nil

func GetListDurationUnitController() *ListDurationUnitController {
	if listDurationUnitController == nil {
		listDurationUnitController = &ListDurationUnitController{}
	}
	return listDurationUnitController
}

// Handle lista de unidades de duración con paginación y filtros
// @Summary Listar Unidades de Duración
// @Description Lista las unidades de duración existentes según filtros dinámicos y parámetros de paginación.
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Ejemplo: `name[cont]=hora`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos (ejemplos):**
// @Description - `name[cont]`
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags DurationUnit
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: name[cont]=hora)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListDurationUnitResponse]] "Se obtuvieron las unidades de duración exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "No se encontraron unidades de duración"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/duration-unit [get]
// @Security BearerAuth
func (c *ListDurationUnitController) Handle(ctx *gin.Context) {
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
	var requestQuery = query.ListDurationUnitQuery{
		PaginationParams: paginationParams,
		Filters:          filters,
	}

	result, _ := mediator.Send[query.ListDurationUnitQuery, response.ApiResponse[paginator.Page[response.ListDurationUnitResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
