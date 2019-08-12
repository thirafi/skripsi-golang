package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func CreateBoxpaper(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	boxpaper := model.BoxpaperModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&boxpaper); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&boxpaper).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, boxpaper)
}

func GetBoxpaper(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	boxpaper := getBoxpaperOr404(db, id, w, r)
	if boxpaper == nil {
		return
	}
	respondJSON(w, http.StatusOK, boxpaper)
}

func UpdateBoxpaper(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	boxpaper := getBoxpaperOr404(db, id, w, r)
	if boxpaper == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&boxpaper); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&boxpaper).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, boxpaper)
}

func DeleteBoxpaper(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	boxpaper := getBoxpaperOr404(db, id, w, r)
	if boxpaper == nil {
		return
	}
	if err := db.Delete(&boxpaper).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAllBoxpaper(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	boxpaper := []model.BoxpaperModel{}

	if err := db.Find(&boxpaper).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}
	respondJSON(w, http.StatusOK, boxpaper)
}

func getBoxpaperOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.BoxpaperModel {
	boxpaper := model.BoxpaperModel{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&boxpaper, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &boxpaper
}

func UploadPaper(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := "paper"
	linkimage := UploadImage(w, r, product)

	respondJSON(w, http.StatusOK, linkimage)
}
