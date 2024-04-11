package controllers

import (
	"api/utils"
	"database/sql"
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
	filename, _ := utils.UploadImage(w, r)
	_, err = db.Exec(`UPDATE MobileDetails SET name=$1, specs=$2, price=$3, image=$4 WHERE ID=$5`, s1.Name, s1.Specs, s1.Price, filename, id)
	if err != nil {
		http.Error(w, "Failed to update mobile details", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "plain/text")
	w.Write([]byte("updated"))

}
