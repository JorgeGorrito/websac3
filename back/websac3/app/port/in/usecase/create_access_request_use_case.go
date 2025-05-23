package usecase

import "websac3/app/port/in/dto/command"

type CreateAccessRequestUseCase interface {
	CreateAccessRequest(command.CreateAccessRequestCommand) error
}
