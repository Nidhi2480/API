package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the image file name from the URL path
	fileName := strings.TrimPrefix(r.URL.Path, "/images/")

	// Construct the file path to the image file
	filePath := "C:/Users/nidhi/Desktop/mobileApps/uploads/" + fileName

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "image not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate content type header
	w.Header().Set("Content-Type", "image/jpeg") // Adjust content type based on the image format

	// Set CORS headers to allow requests from any origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Copy the image data to the response writer
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "error serving image", http.StatusInternalServerError)
		return
	}
}
