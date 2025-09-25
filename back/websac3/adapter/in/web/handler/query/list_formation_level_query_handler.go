package handler

import (
	"net/http"
	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/adapter/in/web/util"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence/filter"
	futil "websac3/common/filter"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/paginator"
	"websac3/common/validator"
)

type ListFormationLevelQueryHandler struct {
	handler.Authenticable
	listFormationLevelUseCase usecase.ListFormationLevelUseCase
	msgProvider               message.Provider
	logger                    logging.Logger
	validator                 validator.Validator

	validFilters []string
}

func NewListFormationLevelQueryHandler(
	listFormationLevelUseCase usecase.ListFormationLevelUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
	validator validator.Validator,
) *ListFormationLevelQueryHandler {
	return &ListFormationLevelQueryHandler{
		Authenticable: handler.Authenticable{
			PermissionsRequired: []string{"list"},
		},
		listFormationLevelUseCase: listFormationLevelUseCase,
		msgProvider:               msgProvider,
		logger:                    logger,
		validator:                 validator,

		validFilters: []string{"name"},
	}
}

func (h *ListFormationLevelQueryHandler) Handle(request query.ListFormationLevelQuery, lang string) (response.ApiResponse[paginator.Page[response.ListFormationLevelsResponse]], error) {
	h.logger.Info("Inicio de consulta de niveles de formación")

	if err := h.validator.ValidateFields(&request, lang); err != nil {
		h.logger.Warn("Advertencia de validación para los datos de entrada al listar niveles de formación. Errores: %v", err)
		var validationErrors []string
		for n := range err {
			validationErrors = append(validationErrors, err[n].Error())
		}
		return response.ApiResponse[paginator.Page[response.ListFormationLevelsResponse]]{
			HttpStatusCode: http.StatusBadRequest,
			Errors:         validationErrors,
		}, nil
	}

	if !h.ValidatePermissions(request.Permissions) {
		h.logger.Warn("El usuario no tiene permisos para listar niveles de formación")
		return response.ApiResponse[paginator.Page[response.ListFormationLevelsResponse]]{
			HttpStatusCode: http.StatusForbidden,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	pagination := &request.PaginationParams
	filters := futil.Transform(request.Filters, []filter.Operator{filter.EqualOperator, filter.ContainsOperator})
	filters.Purge(h.validFilters)
	var name string = ""
	if len(filters) > 0 {
		name, _ = filters[0].Value.(string)
	}
	results, total, err := h.listFormationLevelUseCase.Execute(pagination.Currentpage, pagination.ItemsPerpage, name, lang)

	var resultsMapped []response.ListFormationLevelsResponse
	var errMap error
	for _, result := range results {
		var resultMapped response.ListFormationLevelsResponse
		resultMapped, errMap = mapper.Map[entity.FormationLevel, response.ListFormationLevelsResponse](&result)
		if errMap != nil {
			h.logger.Error("Error al mapear el nivel de formación: %v", errMap)
			return response.ApiResponse[paginator.Page[response.ListFormationLevelsResponse]]{
				HttpStatusCode: http.StatusInternalServerError,
				Errors: []string{
					h.msgProvider.
						WithLang(lang).
						GetMessage("base_error", "internal_error"),
				},
			}, errMap
		}
		resultsMapped = append(resultsMapped, resultMapped)
	}

	resultPaginated, errPag := paginator.New[response.ListFormationLevelsResponse]().
		SetData(resultsMapped).
		SetCurrentPage(pagination.Currentpage).
		SetTotalCount(total).
		SetItemsPerPage(pagination.ItemsPerpage).
		GetPage()
	if errPag != nil {
		h.logger.Error("Error al paginar los resultados de la consulta de niveles de formación. Error: %s", errPag.Error())
		return response.ApiResponse[paginator.Page[response.ListFormationLevelsResponse]]{
			HttpStatusCode: http.StatusInternalServerError,
			Errors: []string{
				h.msgProvider.
					WithLang(lang).
					GetMessage("base_error", "internal_error"),
			},
		}, errPag
	}
	if err != nil {
		h.logger.Error("Error al obtener los resultados de la consulta de niveles de formación. Error: %s", err.Error())
		var httpStatusCode int = util.GetHttpStatusCodeByErr(err)
		return response.ApiResponse[paginator.Page[response.ListFormationLevelsResponse]]{
			HttpStatusCode: httpStatusCode,
			Errors: []string{
				util.GetResultMessageByErr(
					err,
					h.msgProvider.
						WithLang(lang).
						GetMessage("base_error", "internal_error"),
				),
			},
		}, err
	}

	h.logger.Info("Consulta de niveles de formación finalizada con éxito")
	return response.ApiResponse[paginator.Page[response.ListFormationLevelsResponse]]{
		HttpStatusCode: http.StatusOK,
		Result:         *resultPaginated,
	}, nil
}
