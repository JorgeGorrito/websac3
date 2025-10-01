package usecase

import (
	"websac3/adapter/in/web/response"
	"websac3/app/domain/service"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
)

type bulkCreateDegreeProgramUseCaseImpl struct {
	bulkCreateDegreeProgramService *service.BulkCreateDegreeProgramService
}

func NewBulkCreateDegreeProgramUseCaseImpl(
	bulkCreateDegreeProgramService *service.BulkCreateDegreeProgramService,
) usecase.BulkCreateDegreeProgramUseCase {
	return &bulkCreateDegreeProgramUseCaseImpl{
		bulkCreateDegreeProgramService: bulkCreateDegreeProgramService,
	}
}

func (u *bulkCreateDegreeProgramUseCaseImpl) Execute(command command.BulkCreateDegreeProgramCommand, lang string) (response.BulkCreateDegreeProgramResponse, error) {
	return u.bulkCreateDegreeProgramService.Execute(command, lang)
}

