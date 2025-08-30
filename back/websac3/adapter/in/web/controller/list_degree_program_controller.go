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

type ListDegreeProgramController struct{ Authenticable }

var listDegreeProgramController *ListDegreeProgramController = nil

func GetListDegreeProgramController() *ListDegreeProgramController {
	if listDegreeProgramController == nil {
		listDegreeProgramController = &ListDegreeProgramController{}
	}
	return listDegreeProgramController
}

// Handle lista de programas de grado con paginación y filtros
// @Summary Listar Programas de Grado
// @Description Lista los programas de grado existentes según filtros dinámicos y parámetros de paginación.
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Ejemplo: `name[cont]=ingeniería`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos (ejemplos):**
// @Description - `name[cont]`
// @Description - `snies[cont]`
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags DegreeProgram
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: name[cont]=ingeniería)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]] "Se obtuvieron los programas de grado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "No se encontraron programas de grado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program [get]
// @Security BearerAuth
func (c *ListDegreeProgramController) Handle(ctx *gin.Context) {
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
	var requestQuery = query.ListDegreeProgramQuery{
		PaginationParams: paginationParams,
		Filters:          filters,
		Permissions:      token.Permissions["degree-programs"],
		UserRole:         token.Role,
		UserID:           token.Sub,
	}

	result, _ := mediator.Send[query.ListDegreeProgramQuery, response.ApiResponse[paginator.Page[response.ListDegreeProgramResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
