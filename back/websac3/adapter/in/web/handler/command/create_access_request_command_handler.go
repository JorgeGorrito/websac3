package command

import (
	"net/http"
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

type CreateAccessRequestCommandHandler struct {
	createAccessRequestUseCase usecase.CreateAccessRequestUseCase
	msgProvider                message.Provider
	validator                  validator.Validator
	logger                     logging.Logger
}

func NewCreateAccessRequestCommandHandler(
	createAccessRequestUseCase usecase.CreateAccessRequestUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *CreateAccessRequestCommandHandler {
	return &CreateAccessRequestCommandHandler{
		createAccessRequestUseCase: createAccessRequestUseCase,
		msgProvider:                msgProvider,
		validator:                  validator,
		logger:                     logger,
	}
}

func (h *CreateAccessRequestCommandHandler) Handle(request command.CreateAccessRequestCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la creación de solicitud de acceso para el usuario con CC: " + request.IdentificationNumber)
	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al crear solicitud de acceso. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	accessRequest, err := mapper.Map[command.CreateAccessRequestCommand, entity.AccessRequest](&request)
	if err != nil {
		h.logger.Error("Error al mapear datos de entrada a entidad AccessRequest. Error: %s", err.Error())
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err := h.createAccessRequestUseCase.Execute(accessRequest, lang); err != nil {
		h.logger.Error("Error al crear solicitud de acceso para el usuario con CC "+request.IdentificationNumber+". Error: %s", err.Error())
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
	h.logger.Info("Solicitud de acceso creada exitosamente para el usuario con CC: " + request.IdentificationNumber)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result: h.msgProvider.
			WithLang(lang).
			GetMessage("create_access_request", "access_request_created"),
	}, nil
}
