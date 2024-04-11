package controllers

import (
	"api/utils"
	"database/sql"
	"encoding/json"
<<<<<<< HEAD
	"net/http"
	"strconv"
)

func AddMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
=======
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
)

func AddMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa
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
<<<<<<< HEAD
	filename, filePath := utils.UploadImage(w, r)
	_, err = db.Exec("INSERT INTO MobileDetails(name, specs, price, image) VALUES($1, $2, $3, $4)", s1.Name, s1.Specs, s1.Price, filename)
=======
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
	uploadDir := "C:/Users/nidhi/Desktop/mobileApps/uploads"
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
	_, err = db.Exec("INSERT INTO mobile_details(name, specs, price, image) VALUES($1, $2, $3, $4)", s1.Name, s1.Specs, s1.Price, fileName)
>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa
	if err != nil {
		http.Error(w, "failed to insert data into database", http.StatusInternalServerError)
		return
	}
	s1.Image = filePath
<<<<<<< HEAD
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
=======
	// Set response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send response with the inserted data as JSON
>>>>>>> 5b80a44f5c6aa687719752b3d40f36b88def21aa
	json.NewEncoder(w).Encode(s1)
}
