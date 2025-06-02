package command

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/port/out/logging"
	"websac3/common/mapper"
	"websac3/common/validator"
)

type UseCase[E any] interface {
	Execute(entity E) error
}

type LoggerMessages struct {
	InitMessage          string
	InputValidationError string
	MappingError         string
	SuccessMessage       string
	UseCaseError         string
}

type ApiMessages struct {
	SuccessMessage string
	MappingError   string
	UseCaseError   string
}

type MessageBaseApiHandler struct {
	Logger LoggerMessages
	Api    ApiMessages
}

func BaseApiHandler[R any, E any](
	request R,
	useCase UseCase[E],
	logger logging.Logger,
	messages MessageBaseApiHandler,
	successHttpStatusCode int,
) response.ApiResponse[string] {
	logger.Info(messages.Logger.InitMessage)
	if err := validator.ValidateFields(&request); err != nil {
		logger.Error(messages.Logger.InputValidationError+" %s", err.Error())
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Error:          "Error de validación: " + err.Error(),
		}
	}

	var entity E
	if err := mapper.Map(&request, &entity); err != nil {
		logger.Error(messages.Logger.MappingError+" %s", err.Error())
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Error:          messages.Api.MappingError,
		}
	}

	if err := useCase.Execute(entity); err != nil {
		logger.Error(messages.Logger.UseCaseError+" %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[string]{
			HttpStatusCode: httpStatusCode,
			Error:          util.GetResultMessageByErr(err, messages.Api.UseCaseError),
		}
	}
	logger.Info(messages.Logger.SuccessMessage)
	return response.ApiResponse[string]{
		HttpStatusCode: successHttpStatusCode,
		Result:         messages.Api.SuccessMessage,
	}
}
