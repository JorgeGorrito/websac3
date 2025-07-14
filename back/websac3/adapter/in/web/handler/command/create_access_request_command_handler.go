package command

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/validator"
)

type CreateAccessRequestCommandHandler struct {
	createAccessRequestUseCase usecase.CreateAccessRequestUseCase
	logger                     logging.Logger
}

func NewCreateAccessRequestCommandHandler(
	createAccessRequestUseCase usecase.CreateAccessRequestUseCase,
	logger logging.Logger,
) *CreateAccessRequestCommandHandler {
	return &CreateAccessRequestCommandHandler{
		createAccessRequestUseCase: createAccessRequestUseCase,
		logger:                     logger,
	}
}

func (h *CreateAccessRequestCommandHandler) Handle(request command.CreateAccessRequestCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la creación de solicitud de acceso para el usuario con CC: " + request.Person.IdentificationNumber)
	if err := validator.ValidateFields(&request); err != nil {
		h.logger.Error("Error de validación datos de entrada al crear solicitud de acceso. Errores: %v", err)
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
			Errors:         []string{"Error al procesar los datos de la solicitud de acceso"},
		}, nil
	}

	if err := h.createAccessRequestUseCase.Execute(accessRequest, request.RedirectUrlTo, lang); err != nil {
		h.logger.Error("Error al crear solicitud de acceso para el usuario con CC "+request.Person.IdentificationNumber+". Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[string]{
			HttpStatusCode: httpStatusCode,
			Errors:         []string{util.GetResultMessageByErr(err, "Algo salió mal al registrar la solicitud de acceso")},
		}, nil
	}
	h.logger.Info("Solicitud de acceso creada exitosamente para el usuario con CC: " + request.Person.IdentificationNumber)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusOK,
		Result:         "Solicitud de acceso creada exitosamente",
	}, nil
}
