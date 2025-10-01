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

type DeleteDegreeProgramCommandHandler struct {
	handler.Authenticable
	deleteDegreeProgram usecase.DeleteDegreeProgramUseCase
	msgProvider         message.Provider
	logger              logging.Logger
	validator           validator.Validator
}

func NewDeleteDegreeProgramCommandHandler(
	deleteDegreeProgram usecase.DeleteDegreeProgramUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *DeleteDegreeProgramCommandHandler {
	return &DeleteDegreeProgramCommandHandler{
		Authenticable:       handler.Authenticable{PermissionsRequired: []string{"delete"}},
		deleteDegreeProgram: deleteDegreeProgram,
		msgProvider:         msgProvider,
		logger:              logger,
		validator:           validator,
	}
}

func (h *DeleteDegreeProgramCommandHandler) Handle(request command.DeleteDegreeProgramCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio de eliminación de programa de grado con ID: %d", request.DegreeProgramID)

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al eliminar programa de grado. Errores: %v", err)
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
		h.logger.Warn("El usuario con ID %d no tiene permisos para eliminar programas de grado", request.UserID)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	err := h.deleteDegreeProgram.Execute(
		request.DegreeProgramID,
		request.UserID,
		lang,
	)

	if err != nil {
		h.logger.Error("Error al eliminar programa de grado con ID %d: %v", request.DegreeProgramID, err)
		return response.ApiResponse[string]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(err, h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error")),
			},
		}, nil
	}

	h.logger.Info("Programa de grado con ID %d eliminado exitosamente", request.DegreeProgramID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         h.msgProvider.WithLang(lang).GetMessage("delete_degree_program", "degree_program_deleted_successfully"),
	}, nil
}
