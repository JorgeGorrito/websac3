package mapper

import (
	"strings"
	"websac3/adapter/in/web/request"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/common/jwt"
)

func registerUserMappers() {
	RegisterMapFunc(func(registerUserRequest *request.LoginRequest) (command.LoginCommand, error) {
		return command.LoginCommand{
			Email:    registerUserRequest.Email,
			Password: registerUserRequest.Password,
		}, nil
	})

	RegisterMapFunc(func(loginCommand *command.LoginCommand) (entity.User, error) {
		return entity.User{
			Email:    loginCommand.Email,
			Password: loginCommand.Password,
		}, nil
	})

	RegisterMapFunc(func(user *model.User) (entity.User, error) {
		var err error

		var roleID uint = user.RoleID
		var role entity.Role
		role, err = Map[model.Role, entity.Role](&user.Role)
		if err != nil {
			return entity.User{}, err
		}

		var personID uint = user.PersonID
		var person entity.Person
		person, err = Map[model.Person, entity.Person](&user.Person)
		if err != nil {
			return entity.User{}, err
		}

		return entity.User{
			ID:            user.ID,
			Email:         user.Email,
			PasswordHash:  user.PasswordHash,
			RoleID:        roleID,
			Role:          &role,
			PersonID:      personID,
			Person:        &person,
			DeactivatedAt: user.DeactivatedAt,
			CreatedAt:     user.CreatedAt,
		}, nil
	})

	RegisterMapFunc(func(user *entity.User) (model.User, error) {
		var role model.Role
		if user.Role != nil {
			var err error
			role, err = Map[entity.Role, model.Role](user.Role)
			if err != nil {
				return model.User{}, err
			}
		}

		var person model.Person
		if user.Person != nil {
			var err error
			person, err = Map[entity.Person, model.Person](user.Person)
			if err != nil {
				return model.User{}, err
			}
		}

		return model.User{
			ID:            user.ID,
			Email:         user.Email,
			PasswordHash:  user.PasswordHash,
			RoleID:        user.RoleID,
			Role:          role,
			PersonID:      user.PersonID,
			Person:        person,
			DeactivatedAt: user.DeactivatedAt,
			CreatedAt:     user.CreatedAt,
		}, nil
	})

	RegisterMapFunc(
		func(user *entity.User) (jwt.AccessTokenClaims, error) {
			var username string = func() string {
				var firstName, firstLastname string
				if user.Person != nil {
					firstName = strings.Split(user.Person.Name, " ")[0]
					firstLastname = strings.Split(user.Person.Lastname, " ")[0]
				}
				return firstName + " " + firstLastname
			}()

			var permissions map[string][]string = func() map[string][]string {
				var permissions map[string][]string = make(map[string][]string)
				for _, permission := range user.Role.Permissions {
					if _, exists := permissions[permission.Module.Name]; !exists {
						permissions[permission.Module.Name] = []string{}
					}
					permissions[permission.Module.Name] = append(permissions[permission.Module.Name], permission.Action.Name)
				}
				return permissions
			}()

			return jwt.AccessTokenClaims{
				Sub:         user.ID,
				Username:    username,
				Role:        user.Role.Name,
				Permissions: permissions,
			}, nil
		},
	)

	RegisterMapFunc(
		func(user *entity.User) (jwt.RefreshTokenClaims, error) {
			return jwt.RefreshTokenClaims{
				Sub: user.ID,
			}, nil
		},
	)

	RegisterMapFunc(
		func(request *request.RefreshTokenRequest) (command.RefreshTokenCommand, error) {
			return command.RefreshTokenCommand{
				RefreshToken: request.RefreshToken,
			}, nil
		},
	)
}
