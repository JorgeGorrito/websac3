package handler

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	psqlfilter "websac3/app/port/out/persistence/filter"
	"websac3/common/filter"
	"websac3/common/mapper"
	"websac3/common/paginator"
	"websac3/common/validator"

	"websac3/app/port/out/message"
	"websac3/common/logging"
)

type ListReportFeedbacksQueryHandler struct {
	handler.Authenticable
	validator                  validator.Validator
	logger                     logging.Logger
	listReportFeedbacksUseCase usecase.ListReportFeedbacksUseCase
	msgProvider                message.Provider
}

func NewListReportFeedbacksQueryHandler(
	validator validator.Validator,
	logger logging.Logger,
	listReportFeedbacksUseCase usecase.ListReportFeedbacksUseCase,
	msgProvider message.Provider,
) *ListReportFeedbacksQueryHandler {
	return &ListReportFeedbacksQueryHandler{
		Authenticable:              handler.Authenticable{PermissionsRequired: []string{"list"}},
		validator:                  validator,
		logger:                     logger,
		listReportFeedbacksUseCase: listReportFeedbacksUseCase,
		msgProvider:                msgProvider,
	}
}

func (h *ListReportFeedbacksQueryHandler) Handle(req query.ListReportFeedbacksQuery, lang string) (response.ApiResponse[paginator.Page[response.ReportFeedbackResponse]], error) {
	h.logger.Info("Consultando retroalimentaciones de reportes con filtros: %v", req.Filters)

	// Validations request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Error("Error de validación: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ReportFeedbackResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Validations permisos
	if !h.ValidatePermissions(req.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para listar retroalimentaciones de reportes", req.UserID)
		return response.ApiResponse[paginator.Page[response.ReportFeedbackResponse]]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	// Transformar filtros
	filters := filter.Transform(req.Filters, []psqlfilter.Operator{psqlfilter.EqualOperator, psqlfilter.ContainsOperator})

	// Convertir filtros a map[string]interface{} para el use case
	filtersMap := make(map[string]interface{})
	for _, f := range filters {
		filtersMap[f.Field+"."+string(f.Operator)] = f.Value
	}

	// Ejecutar caso de uso
	feedbacks, total, err := h.listReportFeedbacksUseCase.Execute(filtersMap, req.PaginationParams.Currentpage, req.PaginationParams.ItemsPerpage, lang)
	if err != nil {
		h.logger.Error("Error al obtener retroalimentaciones: %v", err)
		return response.ApiResponse[paginator.Page[response.ReportFeedbackResponse]]{
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
	var feedbackResponses []response.ReportFeedbackResponse
	for _, feedback := range feedbacks {
		var feedbackResponse response.ReportFeedbackResponse
		if feedbackResponse, err = mapper.Map[entity.ReportFeedback, response.ReportFeedbackResponse](&feedback); err != nil {
			h.logger.Error("Error al mapear retroalimentación: %v", err)
			return response.ApiResponse[paginator.Page[response.ReportFeedbackResponse]]{
				HttpStatusCode: http.StatusInternalServerError,
				Errors: []string{
					h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				},
			}, nil
		}
		feedbackResponses = append(feedbackResponses, feedbackResponse)
	}

	// Crear página paginada
	page := paginator.Page[response.ReportFeedbackResponse]{
		Data:         feedbackResponses,
		TotalCount:   int64(total),
		Currentpage:  req.PaginationParams.Currentpage,
		ItemsPerpage: req.PaginationParams.ItemsPerpage,
	}

	h.logger.Info("Retroalimentaciones obtenidas exitosamente: %d elementos", len(feedbackResponses))
	return response.ApiResponse[paginator.Page[response.ReportFeedbackResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         page,
	}, nil
}
