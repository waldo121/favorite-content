package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/waldo121/brainrot-index/internal/database"
)

func main() {

	godotenv.Load()
	db := database.GetConnection()
	migrationDirectory := os.Getenv("MIGRATION_FILES")
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{
		MigrationsTable: "version",
	})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		migrationDirectory,
		"sqlite3", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}
