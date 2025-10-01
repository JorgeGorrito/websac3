package command

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/validator"
)

type BulkCreateDegreeProgramCommandHandler struct {
	bulkCreateDegreeProgramUseCase usecase.BulkCreateDegreeProgramUseCase
	validator                      validator.Validator
	msgProvider                    message.Provider
	logger                         logging.Logger
}

func NewBulkCreateDegreeProgramCommandHandler(
	bulkCreateDegreeProgramUseCase usecase.BulkCreateDegreeProgramUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *BulkCreateDegreeProgramCommandHandler {
	return &BulkCreateDegreeProgramCommandHandler{
		bulkCreateDegreeProgramUseCase: bulkCreateDegreeProgramUseCase,
		validator:                      validator,
		msgProvider:                    msgProvider,
		logger:                         logger,
	}
}

func (h *BulkCreateDegreeProgramCommandHandler) Handle(request command.BulkCreateDegreeProgramCommand, lang string) (response.ApiResponse[response.BulkCreateDegreeProgramResponse], error) {
	h.logger.Info("Starting bulk creation of degree programs")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Validation warning for bulk degree program creation. Errors: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[response.BulkCreateDegreeProgramResponse]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("User with ID: %d does not have permissions for bulk degree program creation", request.CreatedBy)
		return response.ApiResponse[response.BulkCreateDegreeProgramResponse]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	result, err := h.bulkCreateDegreeProgramUseCase.Execute(request, lang)
	if err != nil {
		h.logger.Error("Error during bulk degree program creation. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[response.BulkCreateDegreeProgramResponse]{
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

	h.logger.Info("Bulk degree program creation completed. Total: %d, Successful: %d, Failed: %d",
		result.TotalProcessed, result.SuccessfulCount, result.FailedCount)

	return response.ApiResponse[response.BulkCreateDegreeProgramResponse]{
		HttpStatusCode: http.StatusOK,
		Result:         result,
	}, nil
}

func (h *BulkCreateDegreeProgramCommandHandler) ValidatePermissions(permissions []string) bool {
	// Check if user has permission to create degree programs
	for _, permission := range permissions {
		if permission == "create" || permission == "all" {
			return true
		}
	}
	return false
}
