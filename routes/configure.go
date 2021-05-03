package routes

import "database/sql"

// maintaining database state
type DBHandler struct {
	db *sql.DB
}

// Set a new DB for the whole package
func NewDBHandler(db *sql.DB) *DBHandler {
	return &DBHandler{
		db: db,
	}
}
