package usecase

import "websac3/app/port/in/dto/command"

type CreateAccessRequestUseCase interface {
	Execute(command.CreateAccessRequestCommand) error
}
