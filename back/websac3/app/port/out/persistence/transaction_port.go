package persistence

type Context any

type Manager interface {
	ExecuteInTransaction(fn func(tx Context) error) error
	ExecuteNonTransactional(fn func(db Context) error) error
}
