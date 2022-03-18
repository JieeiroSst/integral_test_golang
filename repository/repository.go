package repository

import "fmt"

type DatabaseInterface interface {
	Insert(number int) error
}

type Database struct{}

func New(db Database) Database {
	return db
}

func (db Database) Insert(value int) error {
	fmt.Println("Inserting to database ...")

	return nil
}

