package db

import (
	"database/sql"
)

var DBStore Store
var DB *sql.DB

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	DB *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		DB:      db,
		Queries: New(db),
	}
}

func (s SQLStore) GetDB() *sql.DB {
	return s.GetDB()
}
