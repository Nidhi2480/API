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
	// uploadDir := "C:/Users/nidhi/Desktop/mobileApps/uploads/"
	rows, err := db.Query("SELECT * FROM mobile_details")
	if err != nil {
		http.Error(w, "error scanning row", http.StatusInternalServerError)
	}
	for rows.Next() {
		if err := rows.Scan(&mobile.ID, &mobile.Name, &mobile.Specs, &mobile.Price, &mobile.Image); err != nil {
			http.Error(w, "error scanning rows ", http.StatusInternalServerError)
		}
		// mobile.Image = uploadDir + mobile.Image
		mobiles = append(mobiles, mobile)

	}
	rows.Close()
	jsondata, _ := json.Marshal(mobiles)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-type", "application/json")
	w.Write(jsondata)

}
