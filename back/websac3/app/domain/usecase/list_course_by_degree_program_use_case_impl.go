package usecase

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/service"
)

type ListCourseByDegreeProgramUseCaseImpl struct {
	listCourseByDegreeProgramService *service.ListCourseByDegreeProgramService
}

func NewListCourseByDegreeProgramUseCaseImpl(
	listCourseByDegreeProgramService *service.ListCourseByDegreeProgramService,
) *ListCourseByDegreeProgramUseCaseImpl {
	return &ListCourseByDegreeProgramUseCaseImpl{
		listCourseByDegreeProgramService: listCourseByDegreeProgramService,
	}
}

func (u *ListCourseByDegreeProgramUseCaseImpl) Execute(
	degreeProgramID uint,
	page uint,
	perPage uint,
	name string,
	lang string,
) ([]entity.Course, int64, error) {
	return u.listCourseByDegreeProgramService.Execute(
		degreeProgramID,
		page,
		perPage,
		name,
		lang,
	)
}
