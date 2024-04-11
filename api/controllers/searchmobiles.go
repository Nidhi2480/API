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
	uploadDir := "../uploads/"
	query := r.URL.Query().Get("query")
	fmt.Println(query)
	queryStmt := "SELECT * FROM MobileDetails WHERE LOWER(name) LIKE '%' || LOWER($1) || '%' OR LOWER(specs) LIKE '%' || LOWER($2) || '%'"
	rows, err := db.Query(queryStmt, "%"+query+"%", "%"+query+"%")
	if err != nil {
		http.Error(w, "error in executing query", http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		if err := rows.Scan(&mobile.ID, &mobile.Name, &mobile.Specs, &mobile.Price, &mobile.Image); err != nil {
			http.Error(w, "error scanning rows", http.StatusInternalServerError)
			return
		}
		mobile.Image = uploadDir + mobile.Image
		mobiles = append(mobiles, mobile)
	}
	rows.Close()
	jsondata, _ := json.Marshal(mobiles)
	w.Header().Set("Content-type", "application/json")
	w.Write(jsondata)
}
