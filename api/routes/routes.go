package routes

import (
	"api/controllers"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(rou *mux.Router, db *sql.DB) {

	rou.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Login(w, r, db)
	})
	rou.HandleFunc("/mobiles", func(w http.ResponseWriter, r *http.Request) {
		controllers.MobileList(w, r, db)
	})
	rou.HandleFunc("/getmobile/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.Getmobilebyid(w, r, db)
	})
	rou.HandleFunc("/addmobile", func(w http.ResponseWriter, r *http.Request) {
		controllers.AddMobile(w, r, db)
	})
	rou.HandleFunc("/delmobile/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DelMobile(w, r, db)
	})
	rou.HandleFunc("/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateMobile(w, r, db)
	})
}
