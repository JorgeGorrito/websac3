package handler

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/mapper"
)

type GetUserStatisticsQueryHandler struct {
	getUserStatisticsUseCase usecase.GetUserStatisticsUseCase
	msgProvider              message.Provider
	logger                   logging.Logger
}

func NewGetUserStatisticsQueryHandler(
	getUserStatisticsUseCase usecase.GetUserStatisticsUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
) *GetUserStatisticsQueryHandler {
	return &GetUserStatisticsQueryHandler{
		getUserStatisticsUseCase: getUserStatisticsUseCase,
		msgProvider:              msgProvider,
		logger:                   logger,
	}
}

func (h *GetUserStatisticsQueryHandler) Handle(req query.GetUserStatisticsQuery, lang string) (response.ApiResponse[response.GetUserStatisticsResponse], error) {
	h.logger.Info("Inicio de consulta de estadísticas para el usuario con ID: %d", req.UserID)

	// Ejecutar caso de uso
	stats, err := h.getUserStatisticsUseCase.Execute(req.UserID, lang)
	if err != nil {
		h.logger.Error("Error al obtener estadísticas del usuario ID %d: %v", req.UserID, err)
		httpStatusCode := util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[response.GetUserStatisticsResponse]{
			HttpStatusCode: httpStatusCode,
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				),
			},
		}, nil
	}

	// Convertir entity a response usando mapper
	statsResponse, err := mapper.Map[entity.UserStatistics, response.GetUserStatisticsResponse](&stats)
	if err != nil {
		h.logger.Error("Error al mapear estadísticas a response: %v", err)
		return response.ApiResponse[response.GetUserStatisticsResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Consulta de estadísticas del usuario ID %d completada exitosamente", req.UserID)
	return response.ApiResponse[response.GetUserStatisticsResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         statsResponse,
	}, nil
}
