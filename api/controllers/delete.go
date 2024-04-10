package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DelMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, err := db.Exec("DELETE FROM mobile_details WHERE id=$1", id)
	if err != nil {
		http.Error(w, "cant select", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte("deleted"))

}
