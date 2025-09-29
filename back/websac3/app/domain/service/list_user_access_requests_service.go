package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	commonfilter "websac3/common/filter"
	"websac3/common/paginator"
)

type ListUserAccessRequestsService struct {
	getAccessRequestPort persistence.GetAccessRequestPort
	persistenceManager   db.Manager
	messageProvider      message.Provider
}

func NewListUserAccessRequestsService(
	getAccessRequestPort persistence.GetAccessRequestPort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListUserAccessRequestsService {
	return &ListUserAccessRequestsService{
		getAccessRequestPort: getAccessRequestPort,
		persistenceManager:   persistenceManager,
		messageProvider:      messageProvider,
	}
}

func (s *ListUserAccessRequestsService) Execute(userID uint, paginationParams paginator.PaginationParams, filters commonfilter.Params, lang string) ([]entity.AccessRequest, uint, error) {
	var accessRequests []entity.AccessRequest = make([]entity.AccessRequest, 0)
	var total uint
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Obtener solicitudes de acceso del usuario autenticado
		var err error
		accessRequests, total, err = s.getAccessRequestPort.GetByUserID(userID, paginationParams, filters, ctx)
		if err != nil {
			accessRequests = nil
			return err
		}

		// Verificar si no se encontraron solicitudes
		if len(accessRequests) == 0 {
			err = errs.NewNotFoundError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("list_user_access_requests", "not_found"),
			)
		}

		return err
	})
	return accessRequests, total, err
}
