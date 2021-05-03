package routes

import "net/http"

func Panic(w http.ResponseWriter, r *http.Request) {
	panic("Test")
}
