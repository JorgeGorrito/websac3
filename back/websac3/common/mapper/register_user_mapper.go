package mapper

import (
	"strings"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/common/jwt"
)

// Tipo para los mappers de activate/deactivate
type UserActionInput struct {
	UserID      uint
	Permissions []string
}

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
			// UpdatedAt y DeletedAt se inicializan como nil ya que no existen en el modelo
			UpdatedAt: nil,
			DeletedAt: nil,
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
			// UpdatedAt y DeletedAt no existen en el modelo, se manejan automáticamente por GORM
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

	RegisterMapFunc(
		func(user *entity.User) (response.ListUsersResponse, error) {
			fullName := ""
			if user.Person != nil {
				fullName = user.Person.Name + " " + user.Person.Lastname
			}

			isActive := user.IsActive()

			return response.ListUsersResponse{
				ID:       user.ID,
				FullName: fullName,
				Email:    user.Email,
				IsActive: isActive,
			}, nil
		},
	)

	RegisterMapFunc(
		func(user *entity.User) (response.GetUserProfileResponse, error) {
			profileResponse := response.GetUserProfileResponse{
				ID:       user.ID,
				Email:    user.Email,
				IsActive: user.IsActive(),
			}

			if user.Person != nil {
				profileResponse.Name = user.Person.Name
				profileResponse.Lastname = user.Person.Lastname
				profileResponse.IdentificationNumber = user.Person.IdentificationNumber
				profileResponse.JobPosition = user.Person.JobPosition

				if user.Person.IdentificationType != nil {
					profileResponse.IdentificationTypeID = user.Person.IdentificationType.ID
					profileResponse.IdentificationTypeName = user.Person.IdentificationType.Name
				}

				if user.Person.HigherEducationInstitution != nil {
					profileResponse.InstitutionSnies = user.Person.HigherEducationInstitution.Snies
					profileResponse.InstitutionName = user.Person.HigherEducationInstitution.Name
				}
			}

			if user.Role != nil {
				profileResponse.RoleID = user.Role.ID
				profileResponse.RoleName = user.Role.Name
			}

			return profileResponse, nil
		},
	)

	// Mappers para activate/deactivate user commands
	RegisterMapFunc(func(input *UserActionInput) (command.DeactivateUserCommand, error) {
		return command.DeactivateUserCommand{
			UserID:      input.UserID,
			Permissions: input.Permissions,
		}, nil
	})

	RegisterMapFunc(func(input *UserActionInput) (command.ActivateUserCommand, error) {
		return command.ActivateUserCommand{
			UserID:      input.UserID,
			Permissions: input.Permissions,
		}, nil
	})
}
