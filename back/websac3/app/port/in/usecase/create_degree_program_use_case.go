package usecase

import (
	"websac3/app/port/in/dto/command"
)

type CreateDegreeProgramUseCase interface {
	Execute(command command.CreateDegreeProgramCommand, lang string) error
}
