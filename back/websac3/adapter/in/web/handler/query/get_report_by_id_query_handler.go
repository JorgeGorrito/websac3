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

type GetReportByIDQueryHandler struct {
	handler.Authenticable
	getReportByIDUseCase usecase.GetReportByIDUseCase
	msgProvider          message.Provider
	logger               logging.Logger
	validator            validator.Validator
}

func NewGetReportByIDQueryHandler(
	getReportByIDUseCase usecase.GetReportByIDUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *GetReportByIDQueryHandler {
	return &GetReportByIDQueryHandler{
		Authenticable:        handler.Authenticable{PermissionsRequired: []string{"list"}},
		getReportByIDUseCase: getReportByIDUseCase,
		msgProvider:          msgProvider,
		logger:               logger,
		validator:            validator,
	}
}

func (h *GetReportByIDQueryHandler) Handle(request query.GetReportByIDQuery, lang string) (response.ApiResponse[response.ListReportResponse], error) {
	h.logger.Info("Inicio de consulta de reporte por ID: %d", request.ReportID)

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al obtener reporte. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.ListReportResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	report, err := h.getReportByIDUseCase.Execute(request.ReportID, lang)
	if err != nil {
		h.logger.Error("Error al obtener reporte por ID: %d. Error: %v", request.ReportID, err)
		return response.ApiResponse[response.ListReportResponse]{
			HttpStatusCode: http.StatusNotFound,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("report", "not_found"),
			},
		}, nil
	}

	reportResponse, err := mapper.Map[entity.Report, response.ListReportResponse](&report)
	if err != nil {
		h.logger.Error("Error al mapear reporte a response. Error: %v", err)
		return response.ApiResponse[response.ListReportResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Consulta de reporte por ID completada exitosamente")
	return response.ApiResponse[response.ListReportResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         reportResponse,
	}, nil
}
