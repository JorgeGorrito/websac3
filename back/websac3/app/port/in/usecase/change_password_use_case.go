package usecase

import "websac3/app/port/in/dto/command"

type ChangePasswordUseCase interface {
	Execute(cmd command.ChangePasswordCommand, lang string) error
}
