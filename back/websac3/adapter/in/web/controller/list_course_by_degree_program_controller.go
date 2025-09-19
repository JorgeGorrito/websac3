package controller

import (
	"net/http"
	"net/url"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"
	"websac3/common/paginator"

	"github.com/gin-gonic/gin"
)

type ListCourseByDegreeProgramController struct{ Authenticable }

var listCourseByDegreeProgramController *ListCourseByDegreeProgramController = nil

func GetListCourseByDegreeProgramController() *ListCourseByDegreeProgramController {
	if listCourseByDegreeProgramController == nil {
		listCourseByDegreeProgramController = &ListCourseByDegreeProgramController{}
	}
	return listCourseByDegreeProgramController
}

// Handle lista de cursos asociados a un programa de grado con paginación y filtros
// @Summary Listar Cursos por Programa de Grado
// @Description Lista los cursos asociados a un programa de grado específico según filtros dinámicos y parámetros de paginación.
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Ejemplo: `name[cont]=seguridad`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos (ejemplos):**
// @Description - `name[cont]`
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags DegreeProgram
// @Accept json
// @Produce json
// @Param degree_program_id path int true "ID del programa de grado"
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: name[cont]=criptografía)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListCourseByDegreeProgramResponse]] "Se obtuvieron los cursos exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 404 {object} response.ApiResponse[string] "No se encontraron cursos para el programa de grado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program/{degree_program_id}/courses [get]
// @Security BearerAuth
func (c *ListCourseByDegreeProgramController) Handle(ctx *gin.Context) {
	var err error
	var paginationParams paginator.PaginationParams

	if err = ctx.ShouldBindQuery(&paginationParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Get degree program ID from path parameter
	degreeProgramIDStr := ctx.Param("degree_program_id")
	degreeProgramID, err := strconv.ParseUint(degreeProgramIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid degree program ID"})
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
	// Convert futil.Params to map[string]interface{}
	filtersMapInterface := make(map[string]interface{})
	for k, v := range filters {
		filtersMapInterface[k] = v
	}
	var requestQuery = query.ListCourseByDegreeProgramQuery{
		PaginationParams: paginationParams,
		DegreeProgramID:  uint(degreeProgramID),
		Filters:          filtersMapInterface,
		Permissions:      token.Permissions["courses"],
	}

	result, _ := mediator.Send[query.ListCourseByDegreeProgramQuery, response.ApiResponse[paginator.Page[response.ListCourseByDegreeProgramResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
