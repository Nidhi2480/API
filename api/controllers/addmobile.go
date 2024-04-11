package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func AddMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
	filename, filePath := utils.UploadImage(w, r)
	_, err = db.Exec("INSERT INTO MobileDetails(name, specs, price, image) VALUES($1, $2, $3, $4)", s1.Name, s1.Specs, s1.Price, filename)
	if err != nil {
		http.Error(w, "failed to insert data into database", http.StatusInternalServerError)
		return
	}
	s1.Image = filePath
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s1)
}
