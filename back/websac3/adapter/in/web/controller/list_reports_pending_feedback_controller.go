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

type ListReportsPendingFeedbackController struct {
	Authenticable
}

func NewListReportsPendingFeedbackController() *ListReportsPendingFeedbackController {
	return &ListReportsPendingFeedbackController{}
}

// ListReportsPendingFeedback godoc
// @Summary Listar reportes pendientes de retroalimentación
// @Description Permite a un auditor de ciberseguridad obtener la lista de reportes que no tienen retroalimentación
// @Tags Report Feedback
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param sort_by query string false "Campo por el cual ordenar (created_at)" default(created_at)
// @Param sort_order query string false "Orden de clasificación (asc, desc)" default(desc) Enums(asc, desc)
// @Param DegreeProgram.UserCreator.Person.HigherEducationInstitution.name[eq] query string false "Filtro por nombre exacto de institución educativa"
// @Param DegreeProgram.UserCreator.Person.HigherEducationInstitution.name[cont] query string false "Filtro por nombre de institución educativa que contiene"
// @Param DegreeProgram.UserCreator.Person.HigherEducationInstitution.snies[eq] query int false "Filtro por SNIES exacto de institución educativa"
// @Param DegreeProgram.UserCreator.Person.HigherEducationInstitution.snies[cont] query string false "Filtro por SNIES que contiene (búsqueda parcial)"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ReportResponse]] "Lista de reportes pendientes"
// @Failure 400 {object} response.ApiResponse[paginator.Page[response.ReportResponse]] "Error de validación"
// @Failure 401 {object} response.ApiResponse[paginator.Page[response.ReportResponse]] "No autorizado"
// @Failure 403 {object} response.ApiResponse[paginator.Page[response.ReportResponse]] "Permisos insuficientes"
// @Failure 500 {object} response.ApiResponse[paginator.Page[response.ReportResponse]] "Error interno del servidor"
// @Router /api/v1/{lang}/report-feedbacks/pending-reports [get]
// @Security BearerAuth
func (c *ListReportsPendingFeedbackController) Handle(ctx *gin.Context) {
	var err error
	var paginationParams paginator.PaginationParams

	if err = ctx.ShouldBindQuery(&paginationParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Obtener el usuario autenticado
	token := c.GetToken(ctx)
	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Parsear filtros de la query string
	allParams := ctx.Request.URL.Query()
	filtersMap := make(url.Values)

	// Filtrar solo los parámetros que contienen operadores (filtros)
	for key, values := range allParams {
		if strings.Contains(key, "[") && strings.HasSuffix(key, "]") {
			filtersMap[key] = values
		}
	}
	var filters = util.ParseParamsFilter(filtersMap)

	// Obtener parámetros de ordenamiento
	sortBy := ctx.DefaultQuery("sort_by", "created_at")
	sortOrder := ctx.DefaultQuery("sort_order", "desc")

	// Validar sort_order
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	// Crear el query con filtros y ordenamiento
	queryDto := query.ListReportsPendingFeedbackQuery{
		PaginationParams: paginationParams,
		UserID:           token.Sub,
		Filters:          filters,
		SortBy:           sortBy,
		SortOrder:        sortOrder,
		Permissions:      token.Permissions["reporting-feedback"],
	}

	result, err := mediator.Send[query.ListReportsPendingFeedbackQuery, response.ApiResponse[paginator.Page[response.ReportResponse]]](queryDto, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
