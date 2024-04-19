package additional

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func UploadImage(w http.ResponseWriter, r *http.Request) (string, string) {
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "failed to get file from form", http.StatusBadRequest)
	}
	defer file.Close()
	fileExt := filepath.Ext(handler.Filename)
	fileName := uuid.New().String() + fileExt
	uploadDir := "../uploads"
	filePath := filepath.Join(uploadDir, fileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "failed to create file on server", http.StatusInternalServerError)
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "failed to save file on server", http.StatusInternalServerError)
	}
	return fileName, filePath

}
