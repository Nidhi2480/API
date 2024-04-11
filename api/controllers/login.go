package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user1 utils.User
	var user2 utils.User
	user1.Username = r.FormValue("username")
	user1.Password = r.FormValue("password")
	stmt := `SELECT id,password,role FROM users WHERE username=$1`
	err := db.QueryRow(stmt, user1.Username).Scan(&user2.Id, &user2.Password, &user2.Role)

	w.Header().Set("Content-type", "application/json")
	if err != nil {
		http.Error(w, "no such user", http.StatusInternalServerError)
		return
	}
	if user1.Password == user2.Password {
		user2.Password = "0"
		jsondata, _ := json.Marshal(user2)
		w.Write([]byte(jsondata))
	} else {
		http.Error(w, "password incorrect", http.StatusInternalServerError)

	}

}
