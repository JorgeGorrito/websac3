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

type UpdateCourseCommandHandler struct {
	handler.Authenticable
	updateCourseUseCase usecase.UpdateCourseUseCase
	msgProvider         message.Provider
	logger              logging.Logger
	validator           validator.Validator
}

func NewUpdateCourseCommandHandler(
	updateCourseUseCase usecase.UpdateCourseUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *UpdateCourseCommandHandler {
	return &UpdateCourseCommandHandler{
		Authenticable:       handler.Authenticable{PermissionsRequired: []string{"update"}},
		updateCourseUseCase: updateCourseUseCase,
		msgProvider:         msgProvider,
		logger:              logger,
		validator:           validator,
	}
}

func (h *UpdateCourseCommandHandler) Handle(request command.UpdateCourseCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio de actualización de curso")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al actualizar curso. Errores: %v", err)
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
		h.logger.Warn("El usuario no tiene permisos para actualizar cursos")
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	course, err := mapper.Map[command.UpdateCourseCommand, entity.Course](&request)
	if err != nil {
		h.logger.Error("Error al mapear datos de entrada a entidad Course. Error: %s", err.Error())
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err := h.updateCourseUseCase.Execute(course, request.UserID, lang); err != nil {
		h.logger.Error("Error al actualizar curso %s. Error: %s", request.Name, err.Error())
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
	h.logger.Info("Curso actualizado exitosamente: %s", request.Name)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         h.msgProvider.WithLang(lang).GetMessage("update_course", "course_updated_successfully"),
	}, nil
}
