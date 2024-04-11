package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user1 utils.User
	var user2 utils.User
	user1.Username = r.FormValue("username")
	user1.Password = r.FormValue("password")
<<<<<<< HEAD
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
=======
	// 	stmt := `SELECT * FROM users WHERE username=$1`
	// 	err := db.QueryRow(stmt, user1.Username).Scan(&user2.Id, &user2.Username, &user2.Password, &user2.Email, &user2.Role)
	// 	fmt.Println(user2.Email)
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 	w.Header().Set("Content-type", "plain/text")
	// 	if err != nil {
	// 		w.Write([]byte("there is no such user"))
	// 		return
	// 	}
	// 	if user1.Password == user2.Password {
	// 		w.Write([]byte("successful"))
	// 	} else {
	// 		w.Write([]byte("password is incorrect"))

	// 	}
	// }

	stmt := "SELECT id,password,role FROM users WHERE username=$1 AND password=$2"
	err := db.QueryRow(stmt, user1.Username, user1.Password).Scan(&user2.Id, &user2.Password, &user2.Role)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.Write([]byte("there is no such user"))
		return
	}
	fmt.Println(user2)
	jsondata, _ := json.Marshal(user2)
	w.Write([]byte(jsondata))
>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa

}
