package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Getmobilebyid(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	uploadDir := "../uploads/"
	var mobile utils.MobileDetail
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := db.QueryRow("SELECT * FROM MobileDetails WHERE id=$1", id).Scan(&mobile.ID, &mobile.Name, &mobile.Specs, &mobile.Price, &mobile.Image)
	if err != nil {
		http.Error(w, "error scanning rows ", http.StatusInternalServerError)

	}
	mobile.Image = uploadDir + mobile.Image
	jsondata, _ := json.Marshal(mobile)
	w.Header().Set("Content-type", "application/json")
	w.Write(jsondata)
}
