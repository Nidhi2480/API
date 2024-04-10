package controllers

import (
	"api/utils"
	"database/sql"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func UpdateMobile(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
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
		uploadDir := "C:/Users/nidhi/Desktop/mobileApps/uploads/"
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

		db.Exec(`UPDATE mobile_details SET name=$1, specs=$2, price=$3, image=$4 WHERE ID=$5`, s1.Name, s1.Specs, s1.Price, fileName, id)
		s1.Image = filePath
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-type", "plain/text")
		w.Write([]byte("updated"))
	}
}
