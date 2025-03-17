package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
	connectionString := os.Getenv("DATABASE_CONNECTION_STRING")
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
