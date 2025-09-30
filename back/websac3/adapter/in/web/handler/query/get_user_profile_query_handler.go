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

type GetUserProfileQueryHandler struct {
	getUserProfileUseCase usecase.GetUserProfileUseCase
	msgProvider           message.Provider
	logger                logging.Logger
}

func NewGetUserProfileQueryHandler(
	getUserProfileUseCase usecase.GetUserProfileUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
) *GetUserProfileQueryHandler {
	return &GetUserProfileQueryHandler{
		getUserProfileUseCase: getUserProfileUseCase,
		msgProvider:           msgProvider,
		logger:                logger,
	}
}

func (h *GetUserProfileQueryHandler) Handle(req query.GetUserProfileQuery, lang string) (response.ApiResponse[response.GetUserProfileResponse], error) {
	h.logger.Info("Inicio de consulta de perfil de usuario con ID: %d", req.UserID)

	// Ejecutar caso de uso
	user, err := h.getUserProfileUseCase.Execute(req.UserID, lang)
	if err != nil {
		h.logger.Error("Error al obtener perfil de usuario ID %d: %v", req.UserID, err)
		httpStatusCode := util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[response.GetUserProfileResponse]{
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
	userProfileResponse, err := mapper.Map[entity.User, response.GetUserProfileResponse](&user)
	if err != nil {
		h.logger.Error("Error al mapear usuario a response: %v", err)
		return response.ApiResponse[response.GetUserProfileResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Consulta de perfil de usuario ID %d completada exitosamente", req.UserID)
	return response.ApiResponse[response.GetUserProfileResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         userProfileResponse,
	}, nil
}
