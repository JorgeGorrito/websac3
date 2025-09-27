package command

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/validator"
)

type CreateReportFeedbackCommandHandler struct {
	handler.Authenticable

	validator                   validator.Validator
	logger                      logging.Logger
	createReportFeedbackUseCase usecase.CreateReportFeedbackUseCase
	msgProvider                 message.Provider
}

func NewCreateReportFeedbackCommandHandler(
	validator validator.Validator,
	logger logging.Logger,
	createReportFeedbackUseCase usecase.CreateReportFeedbackUseCase,
	msgProvider message.Provider,
) *CreateReportFeedbackCommandHandler {
	return &CreateReportFeedbackCommandHandler{
		validator:                   validator,
		logger:                      logger,
		createReportFeedbackUseCase: createReportFeedbackUseCase,
		msgProvider:                 msgProvider,

		Authenticable: handler.Authenticable{PermissionsRequired: []string{"create"}},
	}
}

func (h *CreateReportFeedbackCommandHandler) Handle(req command.CreateReportFeedbackCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Iniciando creación de retroalimentación para reporte ID: %d por auditor ID: %d", req.ReportID, req.UserID)

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Error("Error de validación al crear retroalimentación: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Mapear comando a entidad usando el mapper
	feedback, err := mapper.Map[command.CreateReportFeedbackCommand, entity.ReportFeedback](&req)
	if err != nil {
		h.logger.Error("Error al mapear comando a entidad: %v", err)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	// Ejecutar caso de uso
	if err := h.createReportFeedbackUseCase.Execute(req.ReportID, req.UserID, &feedback, lang); err != nil {
		h.logger.Error("Error al crear retroalimentación: %v", err)
		return response.ApiResponse[string]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				),
			},
		}, nil
	}

	h.logger.Info("Retroalimentación creada exitosamente para reporte ID: %d", req.ReportID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusCreated,
		Result:         h.msgProvider.WithLang(lang).GetMessage("report_feedback", "created_successfully"),
	}, nil
}
