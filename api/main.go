package main

import (
	"api/conn"
	"api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := conn.ConnFunc()
	defer db.Close()
	routes.Routes(r, db)
	http.ListenAndServe(":8080", r)

}
