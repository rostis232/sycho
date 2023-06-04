package repository

import "github.com/jmoiron/sqlx"

type ActivityPostgres struct {
	db *sqlx.DB
}

func NewActivityPostgres (db *sqlx.DB) *ActivityPostgres {
	return &ActivityPostgres{
		db: db,
	}
}