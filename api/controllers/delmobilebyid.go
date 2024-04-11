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
	_, err := db.Exec("DELETE FROM MobileDetails WHERE id=$1", id)
	if err != nil {
		http.Error(w, "cant select", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte("deleted"))

}
