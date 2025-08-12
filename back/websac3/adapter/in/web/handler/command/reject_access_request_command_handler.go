package command

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/validator"
)

type RejectAccessRequestCommandHandler struct {
	handler.Authenticable
	rejectAccessRequestUseCase usecase.RejectAccessRequestUseCase
	msgProvider                message.Provider
	validator                  validator.Validator
	logger                     logging.Logger
}

func NewRejectAccessRequestCommandHandler(
	rejectAccessRequestUseCase usecase.RejectAccessRequestUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *RejectAccessRequestCommandHandler {
	return &RejectAccessRequestCommandHandler{
		Authenticable:              handler.Authenticable{PermissionsRequired: []string{"reject"}},
		rejectAccessRequestUseCase: rejectAccessRequestUseCase,
		msgProvider:                msgProvider,
		validator:                  validator,
		logger:                     logger,
	}
}

func (h *RejectAccessRequestCommandHandler) Handle(request command.RejectAccessRequestCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio el rechazo de solicitud de acceso con ID: %d, el usuario con ID: %d", request.AccessRequestID, request.UserID)
	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al rechazar solicitud de acceso. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para rechazar solicitudes de acceso", request.UserID)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	err := h.rejectAccessRequestUseCase.Execute(request.AccessRequestID, lang)
	if err != nil {
		h.logger.Error("Error al rechazar solicitud de acceso con ID: %d, el usuario con ID: %d. Error: %s", request.AccessRequestID, request.UserID, err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[string]{
			HttpStatusCode: httpStatusCode,
			Errors: []string{
				util.GetResultMessageByErr(err, h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error")),
			},
		}, nil
	}

	h.logger.Info("Se ha rechazado la solicitud de acceso con ID: %d, el usuario con ID: %d", request.AccessRequestID, request.UserID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         h.msgProvider.WithLang(lang).GetMessage("reject_access_request", "reject_success"),
	}, nil
}
