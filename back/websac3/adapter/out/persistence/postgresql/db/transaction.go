package db

import (
	_db "websac3/app/port/out/persistence/db"

	"gorm.io/gorm"
)

type Context struct {
	db *gorm.DB
}

func (t *Context) DBSet(db *gorm.DB) {
	t.db = db
}

func (t *Context) DB() *gorm.DB {
	return t.db
}

type Manager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{db: db}
}

func (tm *Manager) ExecuteInTransaction(
	fn func(tx _db.Context) error,
) error {
	var tx *gorm.DB = tm.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	txWrapper := &Context{db: tx}
	if err := fn(txWrapper); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (tm *Manager) ExecuteNonTransactional(
	fn func(ctx _db.Context) error,
) error {
	var db *gorm.DB = tm.db
	if err := db.Error; err != nil {
		return err
	}

	var dbContext *Context = &Context{db: db}
	if err := fn(dbContext); err != nil {
		return err
	}
	return nil
}
