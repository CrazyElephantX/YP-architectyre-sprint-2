package db

import (
	"database/sql"
	"log"

	"cinemaabyss/events/internal/lib/config"
)

// Database connection
var db *sql.DB

func InitDB(cfg config.Config) *sql.DB {

	connStr := cfg.DSN

	if connStr == "" {
		connStr = "postgres://postgres:postgres@localhost/cinemaabyss?sslmode=disable"
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")

	return db
}
