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

type CreateDegreeProgramCommandHandler struct {
	handler.Authenticable
	createDegreeProgramUseCase usecase.CreateDegreeProgramUseCase
	msgProvider                message.Provider
	validator                  validator.Validator
	logger                     logging.Logger
}

func NewCreateDegreeProgramCommandHandler(
	createDegreeProgramUseCase usecase.CreateDegreeProgramUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *CreateDegreeProgramCommandHandler {
	return &CreateDegreeProgramCommandHandler{
		Authenticable:              handler.Authenticable{PermissionsRequired: []string{"create"}},
		createDegreeProgramUseCase: createDegreeProgramUseCase,
		msgProvider:                msgProvider,
		validator:                  validator,
		logger:                     logger,
	}
}

func (h *CreateDegreeProgramCommandHandler) Handle(request command.CreateDegreeProgramCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la creación de programa de grado: %s", request.Name)
	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al crear programa de grado. Errores: %v", err)
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
		h.logger.Warn("El usuario con ID: %d no tiene permisos para crear programas de grado", request.CreatedBy)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	if err := h.createDegreeProgramUseCase.Execute(request, lang); err != nil {
		h.logger.Error("Error al crear programa de grado %s. Error: %s", request.Name, err.Error())
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
	h.logger.Info("Programa de grado creado exitosamente: %s", request.Name)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusCreated,
		Result: h.msgProvider.
			WithLang(lang).
			GetMessage("create_degree_program", "degree_program_created"),
	}, nil
}
