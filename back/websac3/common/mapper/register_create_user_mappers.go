package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/app/port/in/dto/command"
)

func registerCreateUserMappers() {
	RegisterMapFunc(
		func(request *request.CreateUserFromTokenRequest) (command.CreateUserFromTokenCommand, error) {
			return command.CreateUserFromTokenCommand{
				CreateUserToken: request.CreateUserToken,
				Password:        request.Password,
				ConfirmPassword: request.ConfirmPassword,
			}, nil
		},
	)
}
