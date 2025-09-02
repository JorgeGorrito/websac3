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

type CreateCourseCommandHandler struct {
	handler.Authenticable
	createCourseUseCase usecase.CreateCourseUseCase
	msgProvider         message.Provider
	validator           validator.Validator
	logger              logging.Logger
}

func NewCreateCourseCommandHandler(
	createCourseUseCase usecase.CreateCourseUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *CreateCourseCommandHandler {
	return &CreateCourseCommandHandler{
		Authenticable:      handler.Authenticable{PermissionsRequired: []string{"create"}},
		createCourseUseCase: createCourseUseCase,
		msgProvider:         msgProvider,
		validator:           validator,
		logger:              logger,
	}
}

func (h *CreateCourseCommandHandler) Handle(request command.CreateCourseCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la creación de curso: %s", request.Name)
	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al crear curso. Errores: %v", err)
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
		h.logger.Warn("El usuario con ID: %d no tiene permisos para crear cursos", request.CreatedBy)
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	course, err := mapper.Map[command.CreateCourseCommand, entity.Course](&request)
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

	if err := h.createCourseUseCase.Execute(course, lang); err != nil {
		h.logger.Error("Error al crear curso %s. Error: %s", request.Name, err.Error())
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
	h.logger.Info("Curso creado exitosamente: %s", request.Name)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusCreated,
		Result: h.msgProvider.
			WithLang(lang).
			GetMessage("create_course", "course_created"),
	}, nil
}
