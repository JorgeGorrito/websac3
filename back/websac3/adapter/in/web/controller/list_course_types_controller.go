package controller

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"
	"websac3/common/paginator"

	"github.com/gin-gonic/gin"
)

type ListCourseTypesController struct{ Authenticable }

var listCourseTypesController *ListCourseTypesController = nil

func GetListCourseTypesController() *ListCourseTypesController {
	if listCourseTypesController == nil {
		listCourseTypesController = &ListCourseTypesController{}
	}
	return listCourseTypesController
}

// Handle obtiene la lista paginada de tipos de curso
// @Summary Listar Tipos de Curso
// @Description Obtiene una lista paginada de tipos de curso con filtros opcionales.
// @Tags Course
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param page query int false "Número de página" default(1)
// @Param per_page query int false "Elementos por página" default(10)
// @Param name query string false "Filtro por nombre"
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListCourseTypesResponse]] "Lista de tipos de curso obtenida exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course-types [get]
// @Security BearerAuth
func (c *ListCourseTypesController) Handle(ctx *gin.Context) {
	token := c.GetToken(ctx)
	lang := ctx.Param("lang")

	var paginationParams paginator.PaginationParams
	if err := ctx.ShouldBindQuery(&paginationParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	var filters = util.ParseParamsFilter(ctx.Request.URL.Query())
	// Convert futil.Params to map[string]interface{}
	filtersMapInterface := make(map[string]interface{})
	for k, v := range filters {
		filtersMapInterface[k] = v
	}
	var requestQuery = query.ListCourseTypesQuery{
		PaginationParams: paginationParams,
		Filters:          filtersMapInterface,
		Permissions:      token.Permissions["courses"],
	}

	result, _ := mediator.Send[query.ListCourseTypesQuery, response.ApiResponse[paginator.Page[response.ListCourseTypesResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
