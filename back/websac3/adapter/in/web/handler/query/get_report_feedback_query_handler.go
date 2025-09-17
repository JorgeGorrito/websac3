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
	"websac3/common/validator"
)

type GetReportFeedbackQueryHandler struct {
	handler.Authenticable
	validator                validator.Validator
	logger                   logging.Logger
	getReportFeedbackUseCase usecase.GetReportFeedbackUseCase
	msgProvider              message.Provider
}

func NewGetReportFeedbackQueryHandler(
	validator validator.Validator,
	logger logging.Logger,
	getReportFeedbackUseCase usecase.GetReportFeedbackUseCase,
	msgProvider message.Provider,
) *GetReportFeedbackQueryHandler {
	return &GetReportFeedbackQueryHandler{
		Authenticable:            handler.Authenticable{PermissionsRequired: []string{"read"}},
		validator:                validator,
		logger:                   logger,
		getReportFeedbackUseCase: getReportFeedbackUseCase,
		msgProvider:              msgProvider,
	}
}

func (h *GetReportFeedbackQueryHandler) Handle(req query.GetReportFeedbackQuery, lang string) (response.ApiResponse[response.ReportFeedbackResponse], error) {
	h.logger.Info("Consultando retroalimentación para reporte ID: %d", req.ReportID)

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Error("Error de validación: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.ReportFeedbackResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(req.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para obtener retroalimentación de reporte", req.ReportID)
		return response.ApiResponse[response.ReportFeedbackResponse]{
			HttpStatusCode: http.StatusForbidden,
			Errors:         []string{h.msgProvider.WithLang(lang).GetMessage("base_error", "forbidden")},
		}, nil
	}

	// Ejecutar caso de uso
	feedback, err := h.getReportFeedbackUseCase.Execute(req.ReportID, lang)
	if err != nil {
		h.logger.Error("Error al obtener retroalimentación: %v", err)
		return response.ApiResponse[response.ReportFeedbackResponse]{
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
	var feedbackResponse response.ReportFeedbackResponse
	if feedbackResponse, err = mapper.Map[entity.ReportFeedback, response.ReportFeedbackResponse](feedback); err != nil {
		h.logger.Error("Error al mapear retroalimentación: %v", err)
		return response.ApiResponse[response.ReportFeedbackResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Retroalimentación obtenida exitosamente para reporte ID: %d", req.ReportID)
	return response.ApiResponse[response.ReportFeedbackResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         feedbackResponse,
	}, nil
}
