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

func UploadRibbon(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := "ribbon"
	linkimage := UploadImage(w, r, product)

	respondJSON(w, http.StatusOK, linkimage)
}

func CreateRibbon(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	ribbon := model.RibbonModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ribbon); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&ribbon).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, ribbon)
}

func GetRibbon(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	ribbon := getRibbonOr404(db, id, w, r)
	if ribbon == nil {
		return
	}
	respondJSON(w, http.StatusOK, ribbon)
}

func UpdateRibbon(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	ribbon := getRibbonOr404(db, id, w, r)
	if ribbon == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ribbon); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&ribbon).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, ribbon)
}

func DeleteRibbon(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ribbon := getRibbonOr404(db, id, w, r)
	if ribbon == nil {
		return
	}
	if err := db.Delete(&ribbon).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAllRibbon(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	ribbon := []model.RibbonModel{}

	if err := db.Find(&ribbon).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
	}
	respondJSON(w, http.StatusOK, ribbon)
}

func getRibbonOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.RibbonModel {
	ribbon := model.RibbonModel{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&ribbon, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &ribbon
}
