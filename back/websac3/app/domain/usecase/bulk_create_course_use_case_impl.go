package usecase

import (
	"websac3/adapter/in/web/response"
	"websac3/app/domain/service"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
)

type BulkCreateCourseUseCaseImpl struct {
	bulkCreateCourseService *service.BulkCreateCourseService
}

func NewBulkCreateCourseUseCaseImpl(
	bulkCreateCourseService *service.BulkCreateCourseService,
) usecase.BulkCreateCourseUseCase {
	return &BulkCreateCourseUseCaseImpl{
		bulkCreateCourseService: bulkCreateCourseService,
	}
}

func (u *BulkCreateCourseUseCaseImpl) Execute(
	command command.BulkCreateCourseCommand,
	lang string,
) (response.BulkCreateCourseResponse, error) {
	return u.bulkCreateCourseService.Execute(command, lang)
}




