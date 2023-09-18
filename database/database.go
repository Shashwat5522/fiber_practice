package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "root", "secret", "fiber_practice"))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	CreateProductTable()
	fmt.Println("Connection opened to database")
	return nil
}
