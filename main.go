package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vivekmurali2k/webstatus/driver"
	"github.com/vivekmurali2k/webstatus/routes"
	internal "github.com/vivekmurali2k/webstatus/routes/internals"
)

func main() {

	r := chi.NewRouter()

	// Logs every request made
	r.Use(middleware.Logger)

	// Recovers in case of a panic and tries to send an internalservererror response
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	//Connecting databaase
	db := driver.ConnectDB()

	//Initializing internal DB
	iDB := internal.NewDBHandler(db)

	//Initializing routes DB Hander
	h := routes.NewDBHandler(db)

	r.Get("/ping", internal.Ping)
	r.Get("/panic", internal.Panic)
	r.Get("/db", iDB.DBTest)
	r.Get("/query", h.TestQuery)
	r.Post("/register", h.Register)
	r.Post("/login", h.Login)

	http.ListenAndServe(":3000", r)
}
