package repository

import "github.com/jmoiron/sqlx"
type UserRep struct {
	db *sqlx.DB
}

func NewUserRep(db *sqlx.DB) *UserRep {
	return &UserRep{db: db}
}