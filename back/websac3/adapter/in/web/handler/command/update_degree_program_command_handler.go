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

type UpdateDegreeProgramCommandHandler struct {
	handler.Authenticable
	updateDegreeProgramUseCase usecase.UpdateDegreeProgramUseCase
	validator                  validator.Validator
	logger                     logging.Logger
	msgProvider                message.Provider
}

func NewUpdateDegreeProgramCommandHandler(
	validator validator.Validator,
	logger logging.Logger,
	updateDegreeProgramUseCase usecase.UpdateDegreeProgramUseCase,
	msgProvider message.Provider,
) *UpdateDegreeProgramCommandHandler {
	return &UpdateDegreeProgramCommandHandler{
		Authenticable:              handler.Authenticable{PermissionsRequired: []string{"update"}},
		updateDegreeProgramUseCase: updateDegreeProgramUseCase,
		validator:                  validator,
		logger:                     logger,
		msgProvider:                msgProvider,
	}
}

func (h *UpdateDegreeProgramCommandHandler) Handle(request command.UpdateDegreeProgramCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la actualización de programa de grado: %s (ID: %d)", request.Name, request.ID)
	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al actualizar programa de grado. Errores: %v", err)
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
		h.logger.Warn("El usuario con ID: %d no tiene permisos para actualizar programas de grado", request.UpdatedBy)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	if err := h.updateDegreeProgramUseCase.Execute(request, lang); err != nil {
		h.logger.Error("Error al actualizar programa de grado %s (ID: %d). Error: %s", request.Name, request.ID, err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[string]{
			HttpStatusCode: httpStatusCode,
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.
						WithLang(lang).
						GetMessage("base_error", "internal_error"),
				),
			},
		}, nil
	}
	h.logger.Info("Programa de grado actualizado exitosamente: %s (ID: %d)", request.Name, request.ID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result: h.msgProvider.
			WithLang(lang).
			GetMessage("update_degree_program", "degree_program_updated"),
	}, nil
}
