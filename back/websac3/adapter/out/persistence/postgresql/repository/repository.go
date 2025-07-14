package repository

import (
	"errors"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/app/port/out/persistence"
)

type Repository struct{}

func (r *Repository) CastDbContext(ctx persistence.Context) (*db.Context, error) {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return nil, errors.New("db context cast error")
	}
	return dbCtx, nil
}
