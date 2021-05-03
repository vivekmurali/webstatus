package internal

import "net/http"

// Panic the application ===> This is to test the recover
func Panic(w http.ResponseWriter, r *http.Request) {
	panic("Test")
}
