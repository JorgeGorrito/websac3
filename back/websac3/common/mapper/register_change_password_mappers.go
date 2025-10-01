package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/app/port/in/dto/command"
)

func registerChangePasswordMappers() {
	RegisterMapFunc(func(req *request.ChangePasswordRequest) (command.ChangePasswordCommand, error) {
		return command.ChangePasswordCommand{
			CurrentPassword:    req.CurrentPassword,
			NewPassword:        req.NewPassword,
			ConfirmNewPassword: req.ConfirmNewPassword,
		}, nil
	})
}
