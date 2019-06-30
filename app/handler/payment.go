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

func CreatePayment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	payment := model.PaymentModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&payment).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, payment)
}

func GetPayment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	payment := getPaymentOr404(db, id, w, r)
	if payment == nil {
		return
	}
	respondJSON(w, http.StatusOK, payment)
}

func UpdatePayment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	payment := getPaymentOr404(db, id, w, r)
	if payment == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&payment).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, payment)
}

func DeletePayment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	payment := getPaymentOr404(db, id, w, r)
	if payment == nil {
		return
	}
	if err := db.Delete(&payment).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAllPayments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	payments := []model.PaymentModel{}
	db.Find(&payments)
	respondJSON(w, http.StatusOK, payments)
}

func getPaymentOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.PaymentModel {
	payment := model.PaymentModel{}

	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&payment, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &payment
}
