package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var mobiles []utils.MobileDetail
	var mobile utils.MobileDetail
	// uploadDir := "C:/Users/nidhi/Desktop/mobileApps/uploads/"
	query := r.URL.Query().Get("query")
	fmt.Println(query)
	queryStmt := "SELECT * FROM mobile_details WHERE LOWER(name) LIKE '%' || LOWER($1) || '%' OR LOWER(specs) LIKE '%' || LOWER($2) || '%'"
	rows, err := db.Query(queryStmt, "%"+query+"%", "%"+query+"%")
	if err != nil {
		http.Error(w, "error in executing query", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		if err := rows.Scan(&mobile.ID, &mobile.Name, &mobile.Specs, &mobile.Price, &mobile.Image); err != nil {

			return
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
