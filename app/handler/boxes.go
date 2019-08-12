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

func CreateBoxes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	boxes := model.BoxesModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&boxes); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&boxes).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, boxes)
}

func GetBoxes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	boxes := getBoxesOr404(db, id, w, r)
	if boxes == nil {
		return
	}
	respondJSON(w, http.StatusOK, boxes)
}

func UpdateBoxes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	boxes := getBoxesOr404(db, id, w, r)
	if boxes == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&boxes); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&boxes).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, boxes)
}

func DeleteBoxes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	boxes := getBoxesOr404(db, id, w, r)
	if boxes == nil {
		return
	}
	if err := db.Delete(&boxes).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAllBoxes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	boxes := []model.BoxesModel{}

	if err := db.Find(&boxes).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}
	respondJSON(w, http.StatusOK, boxes)
}

func getBoxesOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.BoxesModel {
	boxes := model.BoxesModel{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&boxes, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &boxes
}
func UploadBox(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := "box"
	linkimage := UploadImage(w, r, product)

	respondJSON(w, http.StatusOK, linkimage)
}
