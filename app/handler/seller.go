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

func CreateSeller(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	seller := model.SellerModel{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&seller); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&seller).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, seller)
}

func UploadSeller(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := "seller"
	linkimage := UploadImage(w, r, product)

	respondJSON(w, http.StatusOK, linkimage)
}

func GetSeller(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	seller := getSellerOr404(db, id, w, r)
	if seller == nil {
		return
	}
	respondJSON(w, http.StatusOK, seller)
}

func UpdateSeller(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	seller := getSellerOr404(db, id, w, r)
	if seller == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&seller); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&seller).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, seller)
}

func DeleteSeller(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	seller := getSellerOr404(db, id, w, r)
	if seller == nil {
		return
	}
	if err := db.Delete(&seller).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getSellerOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.SellerModel {
	seller := model.SellerModel{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&seller, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &seller
}
