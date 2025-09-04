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

type ApproveAccessRequestCommandHandler struct {
	handler.Authenticable
	approveAccessRequestUseCase usecase.ApproveAccessRequestUseCase
	msgProvider                 message.Provider
	validator                   validator.Validator
	logger                      logging.Logger
}

func NewApproveAccessRequestCommandHandler(
	approveAccessRequestUseCase usecase.ApproveAccessRequestUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *ApproveAccessRequestCommandHandler {
	return &ApproveAccessRequestCommandHandler{
		Authenticable:               handler.Authenticable{PermissionsRequired: []string{"approve"}},
		approveAccessRequestUseCase: approveAccessRequestUseCase,
		msgProvider:                 msgProvider,
		validator:                   validator,
		logger:                      logger,
	}
}

func (h *ApproveAccessRequestCommandHandler) Handle(request command.ApproveAccessRequestCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la aprobación de solicitud de acceso con ID: %d, el usuario con ID: %d", request.AccessRequestID, request.UserID)
	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al aprobar solicitud de acceso. Errores: %v", err)
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
		h.logger.Warn("El usuario con ID: %d no tiene permisos para aprobar solicitudes de acceso", request.UserID)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	err := h.approveAccessRequestUseCase.Execute(request.AccessRequestID, request.RoleID, lang)
	if err != nil {
		h.logger.Error("Error al aprobar solicitud de acceso con ID: %d, el usuario con ID: %d. Error: %s", request.AccessRequestID, request.UserID, err.Error())
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

	h.logger.Info("Finalizó exitosamente la aprobación de solicitud de acceso con ID: %d, el usuario con ID: %d", request.AccessRequestID, request.UserID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         h.msgProvider.WithLang(lang).GetMessage("approve_access_request", "approve_success"),
	}, nil
}
