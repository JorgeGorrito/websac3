package dependencies

import (
	"os"
	"strconv"
	"time"
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/service"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"
	"websac3/common/jwt"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerUserDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.UserPort](),
		func() any {
			return repository.NewUserRepository(
				container.Inject[message.Provider](),
			)
		},
	)
	m.binder.Bind(
		andi.GetAbstractType[persistence.CreateUserPort](),
		func() any { return container.Inject[persistence.UserPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.UpdateUserPort](),
		func() any { return container.Inject[persistence.UserPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetUserPort](),
		func() any { return container.Inject[persistence.UserPort]() },
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.LoginUseCase](),
		func() any {
			return service.NewLoginService(
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
				container.Inject[persistence.GetUserPort](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.GetUserByIDUseCase](),
		func() any {
			return service.NewGetUserByIDService(
				container.Inject[persistence.GetUserPort](),
				container.Inject[message.Provider](),
				container.Inject[db.Manager](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[usecase.ListUsersUseCase](),
		func() any {
			return service.NewListUsersService(
				container.Inject[persistence.GetUserPort](),
				container.Inject[db.Manager](),
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[jwt.Generator](),
		func() any {
			return jwt.NewGenerator(
				func() string {
					jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
					if jwtSecretKey == "" {
						panic("JWT_SECRET_KEY environment variable is not set")
					}
					return jwtSecretKey
				}(),
				func() string {
					jwtRefreshSecretKey := os.Getenv("JWT_REFRESH_SECRET_KEY")
					if jwtRefreshSecretKey == "" {
						panic("JWT_REFRESH_SECRET_KEY environment variable is not set")
					}
					return jwtRefreshSecretKey
				}(),
				func() time.Duration {
					durationStr := os.Getenv("JWT_EXPIRATION_TIME")
					if durationStr == "" {
						panic("JWT_EXPIRATION_TIME environment variable is not set")
					}
					duration, err := strconv.Atoi(durationStr)
					if err != nil {
						panic("Invalid JWT_EXPIRATION_TIME value: " + err.Error())
					}
					return time.Duration(duration) * time.Second
				}(),
				func() time.Duration {
					durationStr := os.Getenv("JWT_REFRFESH_EXPIRATION_TIME")
					if durationStr == "" {
						panic("JWT_REFRFESH_EXPIRATION_TIME environment variable is not set")
					}
					duration, err := strconv.Atoi(durationStr)
					if err != nil {
						panic("Invalid JWT_REFRFESH_EXPIRATION_TIME value: " + err.Error())
					}
					return time.Duration(duration) * time.Second
				}(),
			)
		},
	)
}
