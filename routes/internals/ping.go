package internal

import "net/http"

// Returns pong 😜
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}
