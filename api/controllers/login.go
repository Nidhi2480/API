package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user1 utils.User
	var password string
	var Mesg utils.Message
	user1.Username = r.FormValue("username")
	password = r.FormValue("password")
	stmt := `SELECT id,password,role FROM users WHERE username=$1`
	err := db.QueryRow(stmt, user1.Username).Scan(&user1.Id, &user1.Password, &user1.Role)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		Mesg.Data = "No Such User"
		jsondata, _ := json.Marshal(Mesg)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsondata)
		return
	}

	if password != user1.Password {
		Mesg.Data = "Password Mismatch"
		jsondata, _ := json.Marshal(Mesg)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsondata)
		return
	}
	tokenString, err := CreateToken(user1.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error generating token: %v", err)
		return
	}

	response := map[string]string{"token": tokenString,
		"role": user1.Role}
	jsondata, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
