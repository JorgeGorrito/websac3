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

type ListAnsweredExpertConsultationsController struct {
	Authenticable
}

func NewListAnsweredExpertConsultationsController() *ListAnsweredExpertConsultationsController {
	return &ListAnsweredExpertConsultationsController{}
}

// ListAnsweredExpertConsultations godoc
// @Summary Listar solicitudes de asesoría respondidas por el experto autenticado
// @Description Permite a los expertos obtener la lista de solicitudes de asesoría que han respondido (aceptadas o rechazadas)
// @Tags ExpertConsultation
// @Accept json
// @Produce json
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Param Requester.Person.name[eq] query string false "Filtro por nombre exacto del solicitante"
// @Param Requester.Person.name[cont] query string false "Filtro por nombre del solicitante que contiene"
// @Param Requester.Person.lastname[eq] query string false "Filtro por apellido exacto del solicitante"
// @Param Requester.Person.lastname[cont] query string false "Filtro por apellido del solicitante que contiene"
// @Param DegreeProgram.name[eq] query string false "Filtro por nombre exacto de programa de grado"
// @Param DegreeProgram.name[cont] query string false "Filtro por nombre de programa de grado que contiene"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.AnsweredExpertConsultationResponse]] "Lista de solicitudes de asesoría respondidas"
// @Failure 400 {object} response.ApiResponse[paginator.Page[response.AnsweredExpertConsultationResponse]] "Error de validación"
// @Failure 401 {object} response.ApiResponse[paginator.Page[response.AnsweredExpertConsultationResponse]] "No autorizado"
// @Failure 403 {object} response.ApiResponse[paginator.Page[response.AnsweredExpertConsultationResponse]] "Permisos insuficientes"
// @Failure 500 {object} response.ApiResponse[paginator.Page[response.AnsweredExpertConsultationResponse]] "Error interno del servidor"
// @Router /api/v1/{lang}/expert-consultations/answered [get]
// @Security BearerAuth
func (c *ListAnsweredExpertConsultationsController) Handle(ctx *gin.Context) {
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

	// Crear el query con filtros
	queryDto := query.ListAnsweredExpertConsultationsQuery{
		PaginationParams: paginationParams,
		UserID:           token.Sub,
		Filters:          filters,
		Permissions:      token.Permissions["expert-advisory"],
	}

	result, err := mediator.Send[query.ListAnsweredExpertConsultationsQuery, response.ApiResponse[paginator.Page[response.AnsweredExpertConsultationResponse]]](queryDto, ctx.GetString("lang"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(result.HttpStatusCode, result.ToResponseFormat())
}


