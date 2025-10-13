package usecase

import (
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
)

type BulkCreateCourseTopicUseCase interface {
	Execute(command command.BulkCreateCourseTopicCommand, lang string) (response.BulkCreateCourseTopicResponse, error)
}


