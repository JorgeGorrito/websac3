package service

import (
	"websac3/adapter/in/web/response"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListExpertConsultationStatusService struct {
	getExpertConsultationStatusPort persistence.GetExpertConsultationStatusPort
	messageProvider                 message.Provider
	persistenceManager              db.Manager
}

func NewListExpertConsultationStatusService(
	getExpertConsultationStatusPort persistence.GetExpertConsultationStatusPort,
	messageProvider message.Provider,
	persistenceManager db.Manager,
) usecase.ListExpertConsultationStatusUseCase {
	return &ListExpertConsultationStatusService{
		getExpertConsultationStatusPort: getExpertConsultationStatusPort,
		messageProvider:                 messageProvider,
		persistenceManager:              persistenceManager,
	}
}

func (s *ListExpertConsultationStatusService) Execute(request query.ListExpertConsultationStatusQuery, lang string) (response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]], error) {
	var expertConsultationStatuses []entity.ExpertConsultationStatus
	var total uint
	var err error

	// Agregar filtro por idioma para optimizar la consulta
	// Esto asegura que solo se carguen los nombres en el idioma solicitado
	langFilter := filter.Params{
		"Names.lang": map[string]string{
			"eq": lang,
		},
	}

	// Combinar filtros existentes con el filtro de idioma
	combinedFilters := request.Filters
	for key, value := range langFilter {
		combinedFilters[key] = value
	}

	// Obtener estados de asesorías de experto
	err = s.persistenceManager.ExecuteNonTransactional(func(ctx db.Context) error {
		expertConsultationStatuses, total, err = s.getExpertConsultationStatusPort.GetAll(request.PaginationParams, combinedFilters, ctx)
		return err
	})

	if err != nil {
		return response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]]{}, err
	}

	// Si no hay datos, retornar error 404
	if len(expertConsultationStatuses) == 0 {
		notFoundMessage := s.messageProvider.WithLang(lang).GetMessage("expert_consultation_statuses", "not_found")
		emptyPage, _ := paginator.New[response.ListExpertConsultationStatusResponse]().
			SetData([]response.ListExpertConsultationStatusResponse{}).
			SetTotalCount(0).
			SetCurrentPage(request.PaginationParams.Currentpage).
			SetItemsPerPage(request.PaginationParams.ItemsPerpage).
			GetPage()
		return response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]]{
			HttpStatusCode: 404,
			Result:         *emptyPage,
			Errors:         []string{notFoundMessage},
		}, nil
	}

	// Mapear a response DTOs
	var responseList []response.ListExpertConsultationStatusResponse
	for _, status := range expertConsultationStatuses {
		// Obtener el nombre en el idioma correcto
		var name string
		for _, statusName := range status.Names {
			if statusName.Lang == lang {
				name = statusName.Name
				break
			}
		}
		// Si no se encuentra en el idioma solicitado, usar el primero disponible
		if name == "" && len(status.Names) > 0 {
			name = status.Names[0].Name
		}

		responseList = append(responseList, response.ListExpertConsultationStatusResponse{
			ID:   status.ID,
			Name: name,
		})
	}

	// Crear página de resultados usando builder pattern
	page, err := paginator.New[response.ListExpertConsultationStatusResponse]().
		SetData(responseList).
		SetTotalCount(int64(total)).
		SetCurrentPage(request.PaginationParams.Currentpage).
		SetItemsPerPage(request.PaginationParams.ItemsPerpage).
		GetPage()

	if err != nil {
		return response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]]{}, err
	}

	successMessage := s.messageProvider.WithLang(lang).GetMessage("expert_consultation_statuses", "retrieved_successfully")
	return response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]]{
		HttpStatusCode: 200,
		Result:         *page,
		Errors:         []string{successMessage},
	}, nil
}
