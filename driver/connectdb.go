// Package driver provides a function connect DB that connects to the database to maintain state of the DB across
// all packages.
package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Database connection
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// Check if database is connected
	if err = db.Ping(); err != nil {
		log.Fatal("Rip in peace DB")
	}
	log.Println("Connected to the database")
	return db
}
