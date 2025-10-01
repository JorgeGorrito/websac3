package usecase

import (
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
)

type BulkCreateDegreeProgramUseCase interface {
	Execute(command command.BulkCreateDegreeProgramCommand, lang string) (response.BulkCreateDegreeProgramResponse, error)
}

