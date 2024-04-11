package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func MobileList(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var mobiles []utils.MobileDetail
	var mobile utils.MobileDetail
	uploadDir := "../uploads/"
	rows, err := db.Query("SELECT * FROM MobileDetails")
	if err != nil {
		http.Error(w, "error scanning row", http.StatusInternalServerError)
	}
	for rows.Next() {
		if err := rows.Scan(&mobile.ID, &mobile.Name, &mobile.Specs, &mobile.Price, &mobile.Image); err != nil {
			http.Error(w, "error scanning rows ", http.StatusInternalServerError)
		}
		mobile.Image = uploadDir + mobile.Image
		mobiles = append(mobiles, mobile)

	}
	rows.Close()
	jsondata, _ := json.Marshal(mobiles)

	w.Header().Set("Content-type", "application/json")
	w.Write(jsondata)

}
