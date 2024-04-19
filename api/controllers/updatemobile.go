package controllers

import (
	"api/additional"
	"api/utils"
	"database/sql"
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	s1 := utils.MobileDetail{
		Name:  r.FormValue("name"),
		Specs: r.FormValue("specs"),
	}
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "invalid price value", http.StatusBadRequest)
		return
	}
	s1.Price = price
	filename := ""
	file, _, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		filename, _ = additional.UploadImage(w, r)
	} else {

		filename = r.FormValue("image1")

	}
	_, err = db.Exec(`UPDATE MobileDetails SET name=$1, specs=$2, price=$3, image=$4 WHERE ID=$5`, s1.Name, s1.Specs, s1.Price, filename, id)
	if err != nil {
		http.Error(w, "Failed to update mobile details", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s1)

}
