package routes

import (
	"encoding/json"
	"fmt"
	"io"
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
		// log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not read the body"))
		return
	}

	defer r.Body.Close()

	// Getting data from json
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not parse json"))
		return
	}

	// Hashing password using bcrypt
	// Passwords to never be saved as plain text
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not register"))
		return
	}

	// Database query to insert user data
	var EmailExistsError = "pq: duplicate key value violates unique constraint \"UNIQUE_EMAIL\""
	var UsernameExistsError = "pq: duplicate key value violates unique constraint \"UNIQUE_USERNAME\""

	insert := `insert into "USER_DATA"("USERNAME", "PASSWORD", "EMAIL_ID", "LEVEL") values ($1, $2, $3, $4)`

	_, err = h.db.Exec(insert, user.Username, hash, user.Email_id, 1)
	if err != nil {

		if EmailExistsError == err.Error() {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("EMAIL ALREADY EXISTS"))
			return
		}

		if UsernameExistsError == err.Error() {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("USERNAME ALREADY EXISTS"))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not read the body"))
		return
	}

	defer r.Body.Close()

	// Getting data from json
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not parse json"))
		return
	}

	var pass string

	// Get only the password stored in the database
	q := fmt.Sprintf(`SELECT "PASSWORD" FROM "USER_DATA" WHERE "USERNAME" = '%s'`, user.Username)

	var UserNotExist = "sql: no rows in result set"

	// Getting hashed password of USER from DB using username
	err = h.db.QueryRow(q).Scan(&pass)
	if err != nil {
		if UserNotExist == err.Error() {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("USER DOES NOT EXIST"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Check if the password entered is the same as the password in the database
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("INCORRECT PASSWORD"))
		return
	}

	w.Write([]byte(user.Password + " " + pass))
}
