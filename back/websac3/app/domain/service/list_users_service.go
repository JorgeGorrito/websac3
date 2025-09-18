package service

import (
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/paginator"
)

type ListUsersService struct {
	getUserPort        persistence.GetUserPort
	persistenceManager db.Manager
}

func NewListUsersService(
	getUserPort persistence.GetUserPort,
	persistenceManager db.Manager,
) *ListUsersService {
	return &ListUsersService{
		getUserPort:        getUserPort,
		persistenceManager: persistenceManager,
	}
}

func (s *ListUsersService) Execute(query query.ListUsersQuery, lang string) (*paginator.Page[entity.User], error) {
	var users []entity.User
	var totalCount int64
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		var err error
		users, totalCount, err = s.getUserPort.ListAllWithPagination(ctx, query.PaginationParams, query.Filters)
		return err
	})

	if err != nil {
		return nil, err
	}

	page := paginator.New[entity.User]().
		SetData(users).
		SetTotalCount(totalCount).
		SetCurrentPage(query.PaginationParams.Currentpage).
		SetItemsPerPage(query.PaginationParams.ItemsPerpage)

	return page.GetPage()
}

var _ usecase.ListUsersUseCase = (*ListUsersService)(nil)
