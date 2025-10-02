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

type BulkCreateCourseTopicCommandHandler struct {
	handler.Authenticable
	bulkCreateCourseTopicUseCase usecase.BulkCreateCourseTopicUseCase
	validator                    validator.Validator
	msgProvider                  message.Provider
	logger                       logging.Logger
}

func NewBulkCreateCourseTopicCommandHandler(
	bulkCreateCourseTopicUseCase usecase.BulkCreateCourseTopicUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *BulkCreateCourseTopicCommandHandler {
	return &BulkCreateCourseTopicCommandHandler{
		Authenticable:                handler.Authenticable{PermissionsRequired: []string{"update"}},
		bulkCreateCourseTopicUseCase: bulkCreateCourseTopicUseCase,
		validator:                    validator,
		msgProvider:                  msgProvider,
		logger:                       logger,
	}
}

func (h *BulkCreateCourseTopicCommandHandler) Handle(
	command command.BulkCreateCourseTopicCommand,
	lang string,
) (response.ApiResponse[response.BulkCreateCourseTopicResponse], error) {
	h.logger.Info("Inicio la carga masiva de tópicos para curso ID: %d", command.CourseID)

	if err := h.validator.ValidateFields(&command, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al cargar tópicos. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.BulkCreateCourseTopicResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(command.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para actualizar cursos", command.UserID)
		return response.ApiResponse[response.BulkCreateCourseTopicResponse]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	result, err := h.bulkCreateCourseTopicUseCase.Execute(command, lang)
	if err != nil {
		h.logger.Error("Error al realizar la carga masiva de tópicos. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[response.BulkCreateCourseTopicResponse]{
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

	h.logger.Info("Bulk course topic creation completed. Total: %d, Successful: %d, Failed: %d",
		result.TotalProcessed, result.SuccessfulCount, result.FailedCount)

	return response.ApiResponse[response.BulkCreateCourseTopicResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         result,
	}, nil
}
