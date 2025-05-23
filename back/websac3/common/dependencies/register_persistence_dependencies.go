package dependencies

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/app/port/out/persistence"

	"github.com/JorgeGorrito/anise-dependency-injection/andi"
	"gorm.io/gorm"
)

func (m *manager) registerPersistenceDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[persistence.TransactionManager](),
		func() any {
			return db.NewTransactionManager(
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
