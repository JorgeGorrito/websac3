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

type ListCourseNaturesController struct{ Authenticable }

var listCourseNaturesController *ListCourseNaturesController = nil

func GetListCourseNaturesController() *ListCourseNaturesController {
	if listCourseNaturesController == nil {
		listCourseNaturesController = &ListCourseNaturesController{}
	}
	return listCourseNaturesController
}

// Handle obtiene la lista paginada de naturalezas de curso
// @Summary Listar Naturalezas de Curso
// @Description Obtiene una lista paginada de naturalezas de curso con filtros opcionales.
// @Tags Course
// @Accept json
// @Produce json
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param page query int false "Número de página" default(1)
// @Param per_page query int false "Elementos por página" default(10)
// @Param name query string false "Filtro por nombre"
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListCourseNaturesResponse]] "Lista de naturalezas de curso obtenida exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course-natures [get]
// @Security BearerAuth
func (c *ListCourseNaturesController) Handle(ctx *gin.Context) {
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
	var requestQuery = query.ListCourseNaturesQuery{
		PaginationParams: paginationParams,
		Filters:          filtersMapInterface,
		Permissions:      token.Permissions["courses"],
	}

	result, _ := mediator.Send[query.ListCourseNaturesQuery, response.ApiResponse[paginator.Page[response.ListCourseNaturesResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
