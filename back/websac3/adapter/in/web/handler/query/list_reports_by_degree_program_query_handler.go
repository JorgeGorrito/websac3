package handler

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/paginator"
	"websac3/common/validator"
)

type ListReportsByDegreeProgramQueryHandler struct {
	handler.Authenticable
	listReportsByDegreeProgramUseCase usecase.ListReportsByDegreeProgramUseCase
	msgProvider                       message.Provider
	logger                            logging.Logger
	validator                         validator.Validator
}

func NewListReportsByDegreeProgramQueryHandler(
	listReportsByDegreeProgramUseCase usecase.ListReportsByDegreeProgramUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *ListReportsByDegreeProgramQueryHandler {
	return &ListReportsByDegreeProgramQueryHandler{
		Authenticable:                     handler.Authenticable{PermissionsRequired: []string{"list"}},
		listReportsByDegreeProgramUseCase: listReportsByDegreeProgramUseCase,
		msgProvider:                       msgProvider,
		logger:                            logger,
		validator:                         validator,
	}
}

func (h *ListReportsByDegreeProgramQueryHandler) Handle(request query.ListReportsByDegreeProgramQuery, lang string) (response.ApiResponse[paginator.Page[response.ListReportSummaryResponse]], error) {
	h.logger.Info("Inicio de consulta de reportes para el programa de grado ID: %d", request.DegreeProgramID)

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar reportes. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListReportSummaryResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	pagination := &request.PaginationParams
	reports, total, err := h.listReportsByDegreeProgramUseCase.Execute(request.DegreeProgramID, pagination.Currentpage, pagination.ItemsPerpage, lang)
	if err != nil {
		h.logger.Error("Error al obtener reportes para el programa de grado ID: %d. Error: %v", request.DegreeProgramID, err)
		return response.ApiResponse[paginator.Page[response.ListReportSummaryResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	var reportResponses []response.ListReportSummaryResponse
	for _, report := range reports {
		reportResponse := response.ListReportSummaryResponse{
			ID:                   report.ID,
			Lang:                 lang,
			Score:                report.Score,
			ProfessionalRoleName: report.ProfessionalRole.Name,
			CreatedAt:            report.CreatedAt,
		}
		reportResponses = append(reportResponses, reportResponse)
	}

	// Crear página paginada
	page := paginator.Page[response.ListReportSummaryResponse]{
		Data:         reportResponses,
		TotalCount:   total,
		Currentpage:  pagination.Currentpage,
		ItemsPerpage: pagination.ItemsPerpage,
	}

	h.logger.Info("Consulta de reportes completada exitosamente. Se encontraron %d reportes de %d totales", len(reportResponses), total)
	return response.ApiResponse[paginator.Page[response.ListReportSummaryResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         page,
	}, nil
}
