package controller

import (
	"net/http"
	"net/url"
	"strings"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"
	"websac3/common/paginator"

	"github.com/gin-gonic/gin"
)

type ListReportFeedbacksController struct {
	Authenticable
}

func NewListReportFeedbacksController() *ListReportFeedbacksController {
	return &ListReportFeedbacksController{}
}

// ListReportFeedbacks godoc
// @Summary Listar retroalimentaciones de reportes
// @Description Permite a un auditor de ciberseguridad obtener la lista de todas las retroalimentaciones de reportes con paginación y filtros
// @Description Los filtros deben enviarse como query params con formato: `campo[operador]=valor`. Por ejemplo: `Auditor.Person.name[cont]=juan`.
// @Description
// @Description **Operadores disponibles:**
// @Description - `[eq]`: igual a (exact match)
// @Description - `[cont]`: contiene (subcadena, case-insensitive)
// @Description
// @Description **Filtros válidos:**
// @Description - `Report.DegreeProgram.name[eq|cont]`
// @Description - `Report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.name[eq|cont]`
// @Description - `Auditor.Person.name[eq|cont]`
// @Description - `Auditor.Person.lastname[eq|cont]`
// @Description - `AuditorRating[eq]`
// @Description
// @Description Cualquier filtro no listado será ignorado automáticamente.
// @Tags Report Feedback
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param filters query string false "Filtros dinámicos: formato campo[operador]=valor (ej: Auditor.Person.name[cont]=juan)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ReportFeedbackResponse]] "Lista de retroalimentaciones obtenida exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido"
// @Failure 401 {object} response.ApiResponse[string] "No autorizado"
// @Failure 403 {object} response.ApiResponse[string] "Permisos insuficientes"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/report-feedbacks [get]
// @Security BearerAuth
func (c *ListReportFeedbacksController) Handle(ctx *gin.Context) {
	var err error
	var paginationParams paginator.PaginationParams

	if err = ctx.ShouldBindQuery(&paginationParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Los filtros vienen como parámetros individuales en la query string
	allParams := ctx.Request.URL.Query()
	filtersMap := make(url.Values)

	// Filtrar solo los parámetros que contienen operadores (filtros)
	for key, values := range allParams {
		if strings.Contains(key, "[") && strings.HasSuffix(key, "]") {
			filtersMap[key] = values
		}
	}

	var filters = util.ParseParamsFilter(filtersMap)
	var lang string = ctx.Param("lang")
	var token = c.GetToken(ctx)

	var requestQuery query.ListReportFeedbacksQuery = query.ListReportFeedbacksQuery{
		PaginationParams: paginationParams,
		Filters:          filters,
		UserID:           token.Sub,
		Permissions:      token.Permissions["reporting-feedback"],
	}

	result, _ := mediator.Send[query.ListReportFeedbacksQuery, response.ApiResponse[paginator.Page[response.ReportFeedbackResponse]]](requestQuery, lang)
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
