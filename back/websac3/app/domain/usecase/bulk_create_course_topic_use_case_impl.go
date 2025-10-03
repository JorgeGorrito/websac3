package usecase

import (
	"websac3/adapter/in/web/response"
	"websac3/app/domain/service"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
)

type BulkCreateCourseTopicUseCaseImpl struct {
	bulkCreateCourseTopicService *service.BulkCreateCourseTopicService
}

func NewBulkCreateCourseTopicUseCaseImpl(
	bulkCreateCourseTopicService *service.BulkCreateCourseTopicService,
) usecase.BulkCreateCourseTopicUseCase {
	return &BulkCreateCourseTopicUseCaseImpl{
		bulkCreateCourseTopicService: bulkCreateCourseTopicService,
	}
}

func (u *BulkCreateCourseTopicUseCaseImpl) Execute(
	command command.BulkCreateCourseTopicCommand,
	lang string,
) (response.BulkCreateCourseTopicResponse, error) {
	return u.bulkCreateCourseTopicService.Execute(command, lang)
}

