package repository

import (
	"errors"
	"websac3/adapter/out/persistence/postgresql/db"
	_db "websac3/app/port/out/persistence/db"
)

type Repository struct{}

func (r *Repository) CastDbContext(ctx _db.Context) (*db.Context, error) {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return nil, errors.New("db context cast error")
	}
	return dbCtx, nil
}
