package models

import "database/sql"

func Init() error {
	var err error

	db, err := sql.Open("postgres", "postgres://postgres:test@localhost:5432/akshakak?sslmode=disable")
	if err != nil {
		return err
	}

	return db.Ping()
}
