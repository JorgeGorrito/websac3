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

type CreateExpertConsultationCommandHandler struct {
	createExpertConsultationUseCase usecase.CreateExpertConsultationUseCase
	msgProvider                     message.Provider
	validator                       validator.Validator
	logger                          logging.Logger
}

func NewCreateExpertConsultationCommandHandler(
	createExpertConsultationUseCase usecase.CreateExpertConsultationUseCase,
	validator validator.Validator,
	msgProvider message.Provider,
	logger logging.Logger,
) *CreateExpertConsultationCommandHandler {
	return &CreateExpertConsultationCommandHandler{
		createExpertConsultationUseCase: createExpertConsultationUseCase,
		msgProvider:                     msgProvider,
		validator:                       validator,
		logger:                          logger,
	}
}

func (h *CreateExpertConsultationCommandHandler) Handle(request command.CreateExpertConsultationCommand, lang string) (response.ApiResponse[string], error) {
	h.logger.Info("Inicio la creación de consulta de experto para el usuario ID: %d", request.RequesterID)

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación datos de entrada al crear consulta de experto. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	expertConsultation, err := mapper.Map[command.CreateExpertConsultationCommand, entity.ExpertConsultation](&request)
	if err != nil {
		h.logger.Error("Error al mapear datos de entrada a entidad ExpertConsultation. Error: %s", err.Error())
		return response.ApiResponse[string]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if err := h.createExpertConsultationUseCase.Execute(expertConsultation, lang); err != nil {
		h.logger.Error("Error al crear consulta de experto para el usuario ID %d. Error: %s", request.RequesterID, err.Error())
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

	h.logger.Info("Consulta de experto creada exitosamente para el usuario ID: %d", request.RequesterID)
	return response.ApiResponse[string]{
		HttpStatusCode: http.StatusCreated,
		Result: h.msgProvider.
			WithLang(lang).
			GetMessage("create_expert_consultation", "expert_consultation_created"),
	}, nil
}
