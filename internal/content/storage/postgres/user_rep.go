package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)
type UserRep struct {
	db *sqlx.DB
}

func NewUserRep(db *sqlx.DB) *UserRep {
	return &UserRep{db: db}
}

func (u *UserRep) Get(ctx context.Context) (int64, error)  {
	return 0, nil
}