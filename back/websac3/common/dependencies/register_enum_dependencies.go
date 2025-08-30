package dependencies

import (
	"os"
	"strconv"
	"time"
	aenum "websac3/adapter/out/persistence/postgresql/enum"
	"websac3/adapter/out/persistence/postgresql/repository"
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/enum"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerEnumDependencies() {
	enumTTLMinutes, err := strconv.Atoi(os.Getenv("ENUM_TTL_MINUTES"))
	if err != nil {
		panic(err)
	}

	m.binder.Bind(
		andi.GetAbstractType[persistence.GetRolePort](),
		func() any { return repository.NewRoleRepository() },
	)

	m.binder.Bind(
		andi.GetAbstractType[enum.StatusEnum](),
		func() any {
			return aenum.New[entity.Status](
				container.Inject[db.Manager](),
				container.Inject[persistence.GetStatusPort](),
				time.Duration(enumTTLMinutes)*time.Minute,
			)
		},
	)

	m.binder.Bind(
		andi.GetAbstractType[enum.RoleEnum](),
		func() any {
			return aenum.New[entity.Role](
				container.Inject[db.Manager](),
				container.Inject[persistence.GetRolePort](),
				time.Duration(enumTTLMinutes)*time.Minute,
			)
		},
	)
}
