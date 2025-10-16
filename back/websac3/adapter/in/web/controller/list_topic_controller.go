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

type ListTopicController struct{ Authenticable }

var listTopicController *ListTopicController = nil

func GetListTopicController() *ListTopicController {
	if listTopicController == nil {
		listTopicController = &ListTopicController{}
	}
	return listTopicController
}

// Handle lista de temas con paginación y filtros
// @Summary Listar Temas
// @Description Lista los temas existentes según filtros dinámicos y parámetros de paginación.
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Ejemplo: `name[cont]=seguridad`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos (ejemplos):**
// @Description - `name[cont]` - busca por nombre del tema
// @Description - `id[cont]` - busca por ID del tema
// @Description
// @Description **Nota:** Los filtros de name e id se aplican con operador OR, es decir, se devolverán los temas que coincidan con cualquiera de los dos criterios.
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags Topic
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: name[cont]=criptografía, id[cont]=12)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListTopicResponse]] "Se obtuvieron los temas exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "No se encontraron temas"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/topic [get]
// @Security BearerAuth
func (c *ListTopicController) Handle(ctx *gin.Context) {
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
	lang := ctx.Param("lang")
	var filters = util.ParseParamsFilter(filtersMap)
	var requestQuery = query.ListTopicQuery{
		PaginationParams: paginationParams,
		Filters:          filters,
		Permissions:      token.Permissions["topics-degree-programs"],
	}

	result, _ := mediator.Send[query.ListTopicQuery, response.ApiResponse[paginator.Page[response.ListTopicResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
