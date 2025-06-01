package command

import (
	"net/http"

	"websac3/adapter/in/web/response"
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
	h.logger.Info("Inicio la creación de solicitud de acceso para el usuario: %s", request.Person.User.Email)
	if err := h.createAccessRequestUseCase.Execute(request); err != nil {
		h.logger.Error("Error al crear solicitud de acceso. %s", err.Error())
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Result:         "Algo salió mal al registrar la solicitud de acceso",
		}, err
	}
	h.logger.Info("Solicitud de acceso creada exitosamente para el usuario: %s", request.Person.User.Email)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusCreated,
		Result:         "Solicitud de acceso creada exitosamente",
	}, nil
}
