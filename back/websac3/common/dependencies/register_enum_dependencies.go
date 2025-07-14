package dependencies

import (
	"os"
	"strconv"
	"time"
	aenum "websac3/adapter/out/persistence/postgresql/enum"
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/enum"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
)

func (m *manager) registerEnumDependencies() {
	accessRequestStatusTTLMinutes, err := strconv.Atoi(os.Getenv("ENUM_ACCESS_REQUEST_STATUS_TTL_MINUTES"))
	if err != nil {
		panic(err)
	}

	m.binder.Bind(
		andi.GetAbstractType[enum.StatusEnum](),
		func() any {
			return aenum.New[entity.Status](
				container.Inject[persistence.Manager](),
				container.Inject[persistence.GetStatusPort](),
				time.Duration(accessRequestStatusTTLMinutes)*time.Minute,
			)
		},
	)
}
