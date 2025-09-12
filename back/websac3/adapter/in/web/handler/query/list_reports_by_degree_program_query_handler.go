package handler

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/mapper"
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

func (h *ListReportsByDegreeProgramQueryHandler) Handle(request query.ListReportsByDegreeProgramQuery, lang string) (response.ApiResponse[[]response.ListReportResponse], error) {
	h.logger.Info("Inicio de consulta de reportes para el programa de grado ID: %d", request.DegreeProgramID)

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar reportes. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[[]response.ListReportResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	reports, err := h.listReportsByDegreeProgramUseCase.Execute(request.DegreeProgramID, lang)
	if err != nil {
		h.logger.Error("Error al obtener reportes para el programa de grado ID: %d. Error: %v", request.DegreeProgramID, err)
		return response.ApiResponse[[]response.ListReportResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	var reportResponses []response.ListReportResponse
	for _, report := range reports {
		reportResponse, err := mapper.Map[entity.Report, response.ListReportResponse](&report)
		if err != nil {
			h.logger.Error("Error al mapear reporte a response. Error: %v", err)
			return response.ApiResponse[[]response.ListReportResponse]{
				HttpStatusCode: http.StatusInternalServerError,
				Errors: []string{
					h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				},
			}, nil
		}
		reportResponses = append(reportResponses, reportResponse)
	}

	h.logger.Info("Consulta de reportes completada exitosamente. Se encontraron %d reportes", len(reportResponses))
	return response.ApiResponse[[]response.ListReportResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         reportResponses,
	}, nil
}
