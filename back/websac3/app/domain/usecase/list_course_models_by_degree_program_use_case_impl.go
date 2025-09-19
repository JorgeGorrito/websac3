package usecase

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/service"
)

type ListCourseModelsByDegreeProgramUseCaseImpl struct {
	listCourseModelsByDegreeProgramService *service.ListCourseModelsByDegreeProgramService
}

func NewListCourseModelsByDegreeProgramUseCaseImpl(
	listCourseModelsByDegreeProgramService *service.ListCourseModelsByDegreeProgramService,
) *ListCourseModelsByDegreeProgramUseCaseImpl {
	return &ListCourseModelsByDegreeProgramUseCaseImpl{
		listCourseModelsByDegreeProgramService: listCourseModelsByDegreeProgramService,
	}
}

func (u *ListCourseModelsByDegreeProgramUseCaseImpl) Execute(
	degreeProgramID uint,
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]model.Course, int64, error) {
	return u.listCourseModelsByDegreeProgramService.Execute(
		degreeProgramID,
		page,
		perPage,
		name,
		lang,
	)
}
