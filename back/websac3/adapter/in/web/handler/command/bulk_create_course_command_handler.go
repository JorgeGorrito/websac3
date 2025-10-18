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

type BulkCreateCourseCommandHandler struct {
	handler.Authenticable
	bulkCreateCourseUseCase usecase.BulkCreateCourseUseCase
	validator               validator.Validator
	msgProvider             message.Provider
	logger                  logging.Logger
}

func NewBulkCreateCourseCommandHandler(
	bulkCreateCourseUseCase usecase.BulkCreateCourseUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *BulkCreateCourseCommandHandler {
	return &BulkCreateCourseCommandHandler{
		Authenticable:           handler.Authenticable{PermissionsRequired: []string{"create"}},
		bulkCreateCourseUseCase: bulkCreateCourseUseCase,
		validator:               validator,
		msgProvider:             msgProvider,
		logger:                  logger,
	}
}

func (h *BulkCreateCourseCommandHandler) Handle(
	command command.BulkCreateCourseCommand,
	lang string,
) (response.ApiResponse[response.BulkCreateCourseResponse], error) {
	h.logger.Info("Inicio la carga masiva de cursos para programa de grado ID: %d", command.DegreeProgramID)

	if err := h.validator.ValidateFields(&command, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al cargar cursos. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.BulkCreateCourseResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(command.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para crear cursos", command.CreatedBy)
		return response.ApiResponse[response.BulkCreateCourseResponse]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	result, err := h.bulkCreateCourseUseCase.Execute(command, lang)
	if err != nil {
		h.logger.Error("Error al realizar la carga masiva de cursos. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[response.BulkCreateCourseResponse]{
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

	h.logger.Info("Bulk course creation completed. Total: %d, Successful: %d, Failed: %d",
		result.TotalProcessed, result.SuccessfulCount, result.FailedCount)

	return response.ApiResponse[response.BulkCreateCourseResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         result,
	}, nil
}







