package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email_id string `json:"email_id"`
}

func (h *DBHandler) Register(w http.ResponseWriter, r *http.Request) {

	var user UserJson

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		log.Fatal(err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Fatal(err)
	}

	insert := `insert into "USER_DATA"("USERNAME", "PASSWORD", "EMAIL_ID", "LEVEL") values ($1, $2, $3, $4)`
	_, err = h.db.Exec(insert, user.Username, hash, user.Email_id, 1)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("INSERTED INTO THE DATABASE"))

}
