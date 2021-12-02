package models

import "database/sql"

func Init() error {
	var err error

	db, err := sql.Open("postgres", "postgres://test:test@localhost:5432/akshakak")
	if err != nil {
		return err
	}

	return db.Ping()
}
