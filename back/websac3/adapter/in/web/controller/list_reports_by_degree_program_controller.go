package controller

import (
	"net/http"
	"strconv"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/common/mediator"
	"websac3/common/paginator"

	"github.com/gin-gonic/gin"
)

type ListReportsByDegreeProgramController struct {
	Authenticable
}

var listReportsByDegreeProgramController *ListReportsByDegreeProgramController = nil

func GetListReportsByDegreeProgramController() *ListReportsByDegreeProgramController {
	if listReportsByDegreeProgramController == nil {
		listReportsByDegreeProgramController = &ListReportsByDegreeProgramController{}
	}
	return listReportsByDegreeProgramController
}

// Handle lista de reportes generados para un programa de grado específico con paginación
// @Summary Listar Reportes por Programa de Grado
// @Description Lista los reportes de evaluación generados para un programa de grado específico con paginación, ordenados por fecha de creación (más recientes primero).
// @Tags Report
// @Accept json
// @Produce json
// @Param degree_program_id path uint true "ID del programa de grado"
// @Param lang path string true "Código de idioma" default(en) Enums(en, es)
// @Param current_page query int false "Número de página (por defecto: 1)" default(1)
// @Param items_per_page query int false "Items por página (por defecto: 10)" default(10)
// @Success 200 {object} response.ApiResponse[paginator.Page[response.ListReportSummaryResponse]] "Se obtuvieron los reportes exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "ID de programa de grado inválido o parámetros de paginación incorrectos"
// @Failure 401 {object} response.ApiResponse[string] "Token de autenticación requerido o inválido"
// @Failure 403 {object} response.ApiResponse[string] "Sin permisos para acceder a este recurso"
// @Failure 404 {object} response.ApiResponse[string] "Programa de grado no encontrado"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/degree-program/{degree_program_id}/reports [get]
// @Security BearerAuth
func (c *ListReportsByDegreeProgramController) Handle(ctx *gin.Context) {
	var err error
	var paginationParams paginator.PaginationParams

	lang := ctx.Param("lang")
	degreeProgramIDStr := ctx.Param("degree_program_id")

	// Validar que el ID sea un número válido
	degreeProgramID, err := strconv.ParseUint(degreeProgramIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors: []string{
				"Invalid degree program ID format",
			},
		})
		return
	}

	// Obtener parámetros de paginación
	paginationParams = paginator.PaginationParams{
		Currentpage:  1,
		ItemsPerpage: 10,
	}

	if currentPageStr := ctx.Query("current_page"); currentPageStr != "" {
		if currentPage, parseErr := strconv.ParseUint(currentPageStr, 10, 32); parseErr == nil {
			paginationParams.Currentpage = uint(currentPage)
		}
	}

	if itemsPerPageStr := ctx.Query("items_per_page"); itemsPerPageStr != "" {
		if itemsPerPage, parseErr := strconv.ParseUint(itemsPerPageStr, 10, 32); parseErr == nil {
			paginationParams.ItemsPerpage = uint(itemsPerPage)
		}
	}

	queryRequest := query.ListReportsByDegreeProgramQuery{
		DegreeProgramID:  uint(degreeProgramID),
		PaginationParams: paginationParams,
	}

	result, err := mediator.Send[query.ListReportsByDegreeProgramQuery, response.ApiResponse[paginator.Page[response.ListReportSummaryResponse]]](queryRequest, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				"Internal server error",
			},
		})
		return
	}

	ctx.JSON(result.HttpStatusCode, result)
}
