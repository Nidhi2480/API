package main

import (
	"api/conn"
	"api/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := conn.ConnFunc()
	defer db.Close()
	corsHeader := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	routes.Routes(r, db)
	http.ListenAndServe(":8080", corsHeader(r))

}
