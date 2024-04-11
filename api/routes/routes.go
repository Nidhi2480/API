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
	}).Methods("POST")
	rou.HandleFunc("/mobiles", func(w http.ResponseWriter, r *http.Request) {
		controllers.MobileList(w, r, db)
	}).Methods("GET")
	rou.HandleFunc("/getmobile/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.Getmobilebyid(w, r, db)
	}).Methods("GET")
	rou.HandleFunc("/addmobile", func(w http.ResponseWriter, r *http.Request) {
		controllers.AddMobile(w, r, db)
	}).Methods("POST")
	rou.HandleFunc("/delmobile/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DelMobile(w, r, db)
	}).Methods("DELETE")
	rou.HandleFunc("/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateMobile(w, r, db)
	}).Methods("PUT")
	rou.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		controllers.SearchMobile(w, r, db)
<<<<<<< HEAD
	}).Methods("GET")
=======
	})
	rou.HandleFunc("/images/{fileName}", controllers.ImageHandler)
>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa
}
