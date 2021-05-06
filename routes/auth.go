package routes

import (
	"encoding/json"
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
