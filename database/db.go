package database 

import (
	"database/sql"
)

type Store interface {
	Get(ID int) (int, error)
}

func NewStore(db *sql.DB) Store {
	return &store{db}
}

type store struct {
	db *sql.DB
}

func (d *store) Get(ID int) (int, error) {
	return 0, nil
}