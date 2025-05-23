package persistence

type Transaction any

type TransactionManager interface {
	ExecuteInTransaction(fn func(tx Transaction) error) error
}
