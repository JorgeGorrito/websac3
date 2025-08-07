package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/db"
	_db "websac3/app/port/out/persistence/db"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
	"gorm.io/gorm"
)

func (m *manager) registerPersistenceDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[_db.Manager](),
		func() any {
			return db.NewManager(
				func() *gorm.DB {
					if conn, err := db.GetConnection(); err != nil {
						panic(err)
					} else {
						return conn
					}
				}(),
			)
		},
	)
}
