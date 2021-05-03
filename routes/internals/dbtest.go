package internal

import (
	"database/sql"
	"log"
	"net/http"
)

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

// Http HandlerFunc for testing DB connection
func (h *DBHandler) DBTest(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		log.Fatal("DB connection failed")
	}

	w.Write([]byte("Hey there DB connected successfully 👍"))
}
