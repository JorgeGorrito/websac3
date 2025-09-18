package handler

import (
	"net/http"

	"websac3/adapter/in/web/handler"
	"websac3/adapter/in/web/response"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/paginator"
)

type ListUsersQueryHandler struct {
	handler.Authenticable
	listUsersUseCase usecase.ListUsersUseCase
	msgProvider      message.Provider
	logger           logging.Logger
}

func NewListUsersQueryHandler(
	listUsersUseCase usecase.ListUsersUseCase,
	msgProvider message.Provider,
	logger logging.Logger,
) *ListUsersQueryHandler {
	return &ListUsersQueryHandler{
		Authenticable:    handler.Authenticable{PermissionsRequired: []string{"list"}},
		listUsersUseCase: listUsersUseCase,
		msgProvider:      msgProvider,
		logger:           logger,
	}
}

func (h *ListUsersQueryHandler) Handle(req query.ListUsersQuery, lang string) (response.ListUsersApiResponse, error) {
	h.logger.Info("Inicio de consulta de usuarios con paginación para el usuario con ID: %d", req.UserID)

	usersPage, err := h.listUsersUseCase.Execute(req, lang)
	if err != nil {
		h.logger.Error("Error al consultar usuarios: %v", err)
		return response.ListUsersApiResponse{
			HttpStatusCode: http.StatusInternalServerError,
			Result:         nil,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
			},
		}, nil
	}

	if !h.ValidatePermissions(req.Permissions) {
		h.logger.Warn("El usuario con ID: %d no tiene permisos para listar usuarios", req.UserID)
		return response.ListUsersApiResponse{
			HttpStatusCode: http.StatusForbidden,
			Result:         nil,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("base_error", "forbidden"),
			},
		}, nil
	}

	// Verificar si no hay usuarios
	if len(usersPage.Data) == 0 {
		h.logger.Info("No se encontraron usuarios")
		return response.ListUsersApiResponse{
			HttpStatusCode: http.StatusNotFound,
			Result:         nil,
			Errors: []string{
				h.msgProvider.WithLang(lang).GetMessage("list_users", "not_found"),
			},
		}, nil
	}

	// Convert entity users to response DTOs using mapper
	var userResponses []response.ListUsersResponse
	for _, user := range usersPage.Data {
		userResponse, err := mapper.Map[entity.User, response.ListUsersResponse](&user)
		if err != nil {
			h.logger.Error("Error al mapear usuario a response: %v", err)
			return response.ListUsersApiResponse{
				HttpStatusCode: http.StatusInternalServerError,
				Result:         nil,
				Errors: []string{
					h.msgProvider.WithLang(lang).GetMessage("base_error", "internal_error"),
				},
			}, nil
		}
		userResponses = append(userResponses, userResponse)
	}

	// Create paginated response
	paginatedResponse := &paginator.Page[response.ListUsersResponse]{
		Data:         userResponses,
		TotalCount:   usersPage.TotalCount,
		Currentpage:  usersPage.Currentpage,
		ItemsPerpage: usersPage.ItemsPerpage,
	}

	h.logger.Info("Consulta de usuarios finalizada con éxito. Se listaron %d usuarios (página %d)", len(userResponses), usersPage.Currentpage)
	return response.ListUsersApiResponse{
		HttpStatusCode: http.StatusOK,
		Result:         paginatedResponse,
		Errors:         []string{},
	}, nil
}
