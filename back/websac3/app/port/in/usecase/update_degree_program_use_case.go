package usecase

import (
	"websac3/app/port/in/dto/command"
)

type UpdateDegreeProgramUseCase interface {
	Execute(command command.UpdateDegreeProgramCommand, lang string) error
}
