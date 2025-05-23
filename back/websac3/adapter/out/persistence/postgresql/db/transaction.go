package db

import (
	"websac3/app/port/out/persistence"

	"gorm.io/gorm"
)

type Transaction struct {
	tx *gorm.DB
}

func (t *Transaction) Tx() *gorm.DB {
	return t.tx
}

type TransactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) ExecuteInTransaction(
	fn func(tx persistence.Transaction) error,
) error {
	var tx *gorm.DB = tm.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	txWrapper := &Transaction{tx: tx}
	if err := fn(txWrapper); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
