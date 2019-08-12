package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func UploadImage(w http.ResponseWriter, r *http.Request, asal string) string {

	if err := r.ParseMultipartForm(1024); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return ""
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return ""
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return ""
	}
	filename := handler.Filename
	t := time.Now()

	filename = fmt.Sprintf("%s%d%d%d%s", asal, t.Second(), t.Minute(), t.Hour(), filepath.Ext(handler.Filename))
	fileLocation := filepath.Join(dir, "public", filename)
	fmt.Println(fileLocation)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return ""
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return ""
	}

	return fileLocation
}

// db.Last(&user) untuk ngambil id terakhir dari record
