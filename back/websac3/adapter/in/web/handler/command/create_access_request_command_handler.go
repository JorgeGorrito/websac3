package command

import (
	"net/http"
	"websac3/adapter/in/web/response"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/logging"
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

func (h *CreateAccessRequestCommandHandler) Handle(request command.CreateAccessRequestCommand) (response.ApiResponse[string], error) {
	return BaseApiHandler[command.CreateAccessRequestCommand, entity.AccessRequest](
		request,
		h.createAccessRequestUseCase,
		h.logger,
		MessageBaseApiHandler{
			Logger: LoggerMessages{
				InitMessage:          "Inicio la creación de solicitud de acceso para el usuario con CC: " + request.Person.IdentificationNumber,
				InputValidationError: "Error de validación datos de entrada al crear solicitud de acceso.",
				MappingError:         "Error al mapear datos de entrada a entidad AccessRequest.",
				SuccessMessage:       "Solicitud de acceso creada exitosamente para el usuario con CC: " + request.Person.IdentificationNumber,
				UseCaseError:         "Error al crear solicitud de acceso para el usuario con CC " + request.Person.IdentificationNumber + ".",
			},
			Api: ApiMessages{
				SuccessMessage: "Solicitud de acceso creada exitosamente",
				MappingError:   "Error al procesar los datos de la solicitud de acceso",
				UseCaseError:   "Algo salió mal al registrar la solicitud de acceso",
			},
		},
		http.StatusCreated,
	), nil
}
