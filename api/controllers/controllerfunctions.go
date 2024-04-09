package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user1 utils.User
	var user2 utils.User
	user1.Username = r.FormValue("username")
	user1.Password = r.FormValue("password")
	stmt := `SELECT * FROM users WHERE username=$1`
	err := db.QueryRow(stmt, user1.Username).Scan(&user2.Id, &user2.Username, &user2.Password, &user2.Email, &user2.Role)
	fmt.Println(user2.Email)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-type", "plain/text")
	if err != nil {
		w.Write([]byte("there is no such user"))
		return
	}
	if user1.Password == user2.Password {
		w.Write([]byte("successful"))
	} else {
		w.Write([]byte("password is incorrect"))

	}

}
func MobileList(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var mobiles []utils.MobileDetail
	var mobile utils.MobileDetail
	uploadDir := "/home/nidhinsajeev/Desktop/mobileapp/uploads/"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-type", "application/json")
	w.Write(jsondata)

}

func Getmobilebyid(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	uploadDir := "/home/nidhinsajeev/Desktop/mobileapp/uploads/"
	var mobile utils.MobileDetail
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := db.QueryRow("SELECT * FROM MobileDetails WHERE id=$1", id).Scan(&mobile.ID, &mobile.Name, &mobile.Specs, &mobile.Price, &mobile.Image)
	if err != nil {
		http.Error(w, "error scanning rows ", http.StatusInternalServerError)

	}
	mobile.Image = uploadDir + mobile.Image
	jsondata, _ := json.Marshal(mobile)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-type", "application/json")
	w.Write(jsondata)
}

func AddMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

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
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique file name using UUID and original file extension
	fileExt := filepath.Ext(handler.Filename)
	fileName := uuid.New().String() + fileExt

	// Save the file to a designated directory on the server
	uploadDir := "/home/nidhinsajeev/Desktop/mobileapp/uploads"
	filePath := filepath.Join(uploadDir, fileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "failed to create file on server", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "failed to save file on server", http.StatusInternalServerError)
		return
	}

	// Save the file name (or file path) in the database
	_, err = db.Exec("INSERT INTO MobileDetails(name, specs, price, image) VALUES($1, $2, $3, $4)", s1.Name, s1.Specs, s1.Price, fileName)
	if err != nil {
		http.Error(w, "failed to insert data into database", http.StatusInternalServerError)
		return
	}
	s1.Image = filePath
	// Set response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send response with the inserted data as JSON
	json.NewEncoder(w).Encode(s1)
}

func DelMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, err := db.Exec("DELETE FROM MobileDetails WHERE id=$1", id)
	if err != nil {
		http.Error(w, "cant select", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte("deleted"))

}

func UpdateMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
			return
		}
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
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "failed to get file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Generate a unique file name using UUID and original file extension
		fileExt := filepath.Ext(handler.Filename)
		fileName := uuid.New().String() + fileExt

		// Save the file to a designated directory on the server
		uploadDir := "/home/nidhinsajeev/Desktop/mobileapp/uploads"
		filePath := filepath.Join(uploadDir, fileName)
		outFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "failed to create file on server", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()
		_, err = io.Copy(outFile, file)
		if err != nil {
			http.Error(w, "failed to save file on server", http.StatusInternalServerError)
			return
		}

		db.Exec(`UPDATE MobileDetails SET name=$1, specs=$2, price=$3, image=$4 WHERE ID=$5`, s1.Name, s1.Specs, s1.Price, fileName, id)
		s1.Image = filePath
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-type", "plain/text")
		w.Write([]byte("updated"))
	}
}
