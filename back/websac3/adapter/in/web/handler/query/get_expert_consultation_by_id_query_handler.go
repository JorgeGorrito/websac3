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

type GetExpertConsultationByIDQueryHandler struct {
	handler.Authenticable
	validator                        validator.Validator
	logger                           logging.Logger
	getExpertConsultationByIDUseCase usecase.GetExpertConsultationByIDUseCase
	msgProvider                      message.Provider
}

func NewGetExpertConsultationByIDQueryHandler(
	validator validator.Validator,
	logger logging.Logger,
	getExpertConsultationByIDUseCase usecase.GetExpertConsultationByIDUseCase,
	msgProvider message.Provider,
) *GetExpertConsultationByIDQueryHandler {
	return &GetExpertConsultationByIDQueryHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"read"},
		},
		validator:                        validator,
		logger:                           logger,
		getExpertConsultationByIDUseCase: getExpertConsultationByIDUseCase,
		msgProvider:                      msgProvider,
	}
}

func (h *GetExpertConsultationByIDQueryHandler) Handle(req query.GetExpertConsultationByIDQuery, lang string) (response.ApiResponse[response.UserExpertConsultationResponse], error) {
	h.logger.Info("Obteniendo asesoría de experto con ID: %d", req.ConsultationID)

	// Validar request
	if err := h.validator.ValidateFields(&req, lang); err != nil {
		h.logger.Error("Error de validación: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.UserExpertConsultationResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	// Validar permisos
	if !h.ValidatePermissions(req.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para ver la asesoría con ID: %d", req.UserID, req.ConsultationID)
		return response.ApiResponse[response.UserExpertConsultationResponse]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	// Ejecutar caso de uso
	expertConsultation, err := h.getExpertConsultationByIDUseCase.Execute(req.ConsultationID, lang)
	if err != nil {
		h.logger.Error("Error al obtener asesoría de experto con ID: %d. Error: %v", req.ConsultationID, err)
		return response.ApiResponse[response.UserExpertConsultationResponse]{
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
	var expertConsultationResponse response.UserExpertConsultationResponse
	expertConsultationResponse, err = mapper.Map[entity.ExpertConsultation, response.UserExpertConsultationResponse](&expertConsultation)
	if err != nil {
		h.logger.Error("Error al mapear asesoría de experto: %v", err)
		return response.ApiResponse[response.UserExpertConsultationResponse]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	h.logger.Info("Asesoría de experto con ID: %d obtenida exitosamente", req.ConsultationID)
	return response.ApiResponse[response.UserExpertConsultationResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         expertConsultationResponse,
	}, nil
}
