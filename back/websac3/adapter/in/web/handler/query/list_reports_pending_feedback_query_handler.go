package handler

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/paginator"
	"websac3/common/validator"
)

type ListReportsPendingFeedbackQueryHandler struct {
	handler.Authenticable
	validator                         validator.Validator
	logger                            logging.Logger
	listReportsPendingFeedbackUseCase usecase.ListReportsPendingFeedbackUseCase
	msgProvider                       message.Provider
}

func NewListReportsPendingFeedbackQueryHandler(
	validator validator.Validator,
	logger logging.Logger,
	listReportsPendingFeedbackUseCase usecase.ListReportsPendingFeedbackUseCase,
	msgProvider message.Provider,
) *ListReportsPendingFeedbackQueryHandler {
	return &ListReportsPendingFeedbackQueryHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"list"},
		},
		validator:                         validator,
		logger:                            logger,
		listReportsPendingFeedbackUseCase: listReportsPendingFeedbackUseCase,
		msgProvider:                       msgProvider,
	}
}

func (h *ListReportsPendingFeedbackQueryHandler) Handle(req query.ListReportsPendingFeedbackQuery, lang string) (response.ApiResponse[paginator.Page[response.ReportResponse]], error) {
	h.logger.Info("Listando reportes pendientes de retroalimentación para auditor ID: %d", req.UserID)

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Error("Error de validación: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ReportResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Validar permisos
	if !h.ValidatePermissions(req.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para listar reportes pendientes de retroalimentación", req.UserID)
		return response.ApiResponse[paginator.Page[response.ReportResponse]]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	// Ejecutar caso de uso
	reports, total, err := h.listReportsPendingFeedbackUseCase.Execute(req.UserID, req.PaginationParams, lang)
	if err != nil {
		h.logger.Error("Error al obtener reportes pendientes: %v", err)
		return response.ApiResponse[paginator.Page[response.ReportResponse]]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				),
			},
		}, nil
	}

	// Mapear a response
	var reportResponses []response.ReportResponse
	for _, report := range reports {
		var reportResponse response.ReportResponse
		if reportResponse, err = mapper.Map[entity.Report, response.ReportResponse](&report); err != nil {
			h.logger.Error("Error al mapear reporte: %v", err)
			continue
		}
		reportResponses = append(reportResponses, reportResponse)
	}

	paginatedResponse := paginator.Page[response.ReportResponse]{
		Data:         reportResponses,
		TotalCount:   int64(total),
		Currentpage:  req.PaginationParams.Currentpage,
		ItemsPerpage: req.PaginationParams.ItemsPerpage,
	}

	h.logger.Info("Listado de reportes pendientes completado. Total: %d", total)
	return response.ApiResponse[paginator.Page[response.ReportResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         paginatedResponse,
	}, nil
}
