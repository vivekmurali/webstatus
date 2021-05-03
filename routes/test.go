package routes

import (
	"fmt"
	"log"
	"net/http"

	"strings"
)

// A test query to return the column of a table
func (h *DBHandler) TestQuery(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT * FROM NEWTEST;")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	defer rows.Close()
	col, _ := rows.Columns()
	v := strings.Join(col, " ")
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(v))
}
