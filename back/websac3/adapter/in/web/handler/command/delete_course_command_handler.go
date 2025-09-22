package command

import (
	"fmt"
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

type DeleteCourseCommandHandler struct {
	handler.Authenticable
	deleteCourseUseCase usecase.DeleteCourseUseCase
	msgProvider         message.Provider
	logger              logging.Logger
	validator           validator.Validator
}

func NewDeleteCourseCommandHandler(
	deleteCourseUseCase usecase.DeleteCourseUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *DeleteCourseCommandHandler {
	return &DeleteCourseCommandHandler{
		Authenticable:       handler.Authenticable{PermissionsRequired: []string{"delete"}},
		deleteCourseUseCase: deleteCourseUseCase,
		msgProvider:         msgProvider,
		logger:              logger,
		validator:           validator,
	}
}

func (h *DeleteCourseCommandHandler) Handle(request command.DeleteCourseCommand, lang string) (response.ApiResponse[string], error) {
	fmt.Println("Inicio de eliminación de curso")
	h.logger.Info("Inicio de eliminación de curso")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al eliminar curso. Errores: %v", err)
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
		h.logger.Warn("El usuario no tiene permisos para eliminar cursos")
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	err := h.deleteCourseUseCase.Execute(
		request.CourseID,
		request.UserID,
		lang,
	)

	if err != nil {
		h.logger.Error("Error al eliminar curso: %v", err)
		return response.ApiResponse[string]{
			HttpStatusCode: util.GetHttpStatusCodeByErr(err),
			Errors: []string{
				util.GetResultMessageByErr(err, h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error")),
			},
		}, nil
	}

	h.logger.Info("Curso eliminado exitosamente")
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         h.msgProvider.WithLang(lang).GetMessage("delete_course", "course_deleted_successfully"),
	}, nil
}
