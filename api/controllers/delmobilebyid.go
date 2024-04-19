package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DelMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var Mesg utils.Message
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, err := db.Exec("DELETE FROM MobileDetails WHERE id=$1", id)
	if err != nil {
		http.Error(w, "cant select", http.StatusInternalServerError)
		return
	}
	Mesg.Data = "Mobile Deleted"
	jsondata, _ := json.Marshal(Mesg)
	w.Header().Set("Content-type", "application/json")
	w.Write(jsondata)
}
