package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Struct for decoding json
type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email_id string `json:"email_id"`
}

// Register method
func (h *DBHandler) Register(w http.ResponseWriter, r *http.Request) {

	var user UserRegister

	// Reading the request's body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	// Getting data from json
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		log.Fatal(err)
	}

	// Hashing password using bcrypt
	// Passwords to never be saved as plain text
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Fatal(err)
	}

	// Database query to insert user data
	insert := `insert into "USER_DATA"("USERNAME", "PASSWORD", "EMAIL_ID", "LEVEL") values ($1, $2, $3, $4)`
	_, err = h.db.Exec(insert, user.Username, hash, user.Email_id, 1)
	if err != nil {
		log.Fatal(err)
	}

	// Respone 200 and giving a respone of inserted into the database
	// Only if everything worked
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("INSERTED INTO THE DATABASE"))

}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

func (h *DBHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user UserLogin

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	// Getting data from json
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		log.Fatal(err)
	}

	var pass string

	// Get only the password stored in the database
	q := fmt.Sprintf(`SELECT "PASSWORD" FROM "USER_DATA" WHERE "USERNAME" = '%s'`, user.Username)

	// execute the query
	rows, err := h.db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Initialize the value returned from the database to the pass string
	for rows.Next() {

		err = rows.Scan(&pass)
		if err != nil {
			log.Fatal(err)
		}

	}

	// Check if the password entered is the same as the password in the database
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password))
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(user.Password + " " + pass))
}
