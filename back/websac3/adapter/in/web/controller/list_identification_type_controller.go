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

type ListIdentificationTypeController struct{}

var listIdentificationTypeControllerInstance *ListIdentificationTypeController = nil

func GetListIdentificationTypeController() *ListIdentificationTypeController {
	if listIdentificationTypeControllerInstance == nil {
		listIdentificationTypeControllerInstance = &ListIdentificationTypeController{}
	}
	return listIdentificationTypeControllerInstance
}

// Handle lista los tipos de identificación con paginación y filtros
// @Summary Listar Tipos de Identificación
// @Description Lista los tipos de identificación existentes según filtros dinámicos y parámetros de paginación.
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Por ejemplo: `name[cont]=cedula`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[eq]`: igual a (exact match)
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos:**
// @Description - `name[eq|cont]`
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags IdentificationType
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: name[cont]=cedula)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListIdentificationTypeResponse]] "Se obtuvieron los tipos de identificación exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "No se encontraron tipos de identificación"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/identification-types [get]
func (c *ListIdentificationTypeController) Handle(ctx *gin.Context) {
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
	var requestQuery = query.ListIdentificationTypeQuery{
		PaginationParams: paginationParams,
		Filters:          filters,
	}

	result, _ := mediator.Send[query.ListIdentificationTypeQuery, response.ApiResponse[paginator.Page[response.ListIdentificationTypeResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
	return
}
