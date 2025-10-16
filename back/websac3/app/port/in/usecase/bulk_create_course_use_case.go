package usecase

import (
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
)

type BulkCreateCourseUseCase interface {
	Execute(command command.BulkCreateCourseCommand, lang string) (response.BulkCreateCourseResponse, error)
}





