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

func CreateBuyer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	buyer := model.BuyerModel{}
	fmt.Println("yang login ")
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&buyer); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&buyer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, buyer)
}

func GetBuyer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	buyer := getBuyerOr404(db, id, w, r)
	if buyer == nil {
		return
	}
	respondJSON(w, http.StatusOK, buyer)
}

func UpdateBuyer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	buyer := getBuyerOr404(db, id, w, r)
	if buyer == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&buyer); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&buyer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, buyer)
}

func DeleteBuyer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	buyer := getBuyerOr404(db, id, w, r)
	if buyer == nil {
		return
	}
	if err := db.Delete(&buyer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAddressByUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID := vars["id"]
	user := getUserOr404(db, userID, w, r)
	if user == nil {
		return
	}
	address := model.BuyerModel{}
	if err := db.Where("account_id=?", userID).First(&address).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, address)
}

func getBuyerOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.BuyerModel {
	buyer := model.BuyerModel{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&buyer, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &buyer
}
