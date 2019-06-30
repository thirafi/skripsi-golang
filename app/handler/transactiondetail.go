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

func CreateTransactionDetail(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	transactiondetail := model.TransactionDetailModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transactiondetail); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&transactiondetail).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, transactiondetail)
}

func GetTransactionDetail(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	transactiondetail := getTransactionDetailOr404(db, id, w, r)
	if transactiondetail == nil {
		return
	}
	respondJSON(w, http.StatusOK, transactiondetail)
}

func UpdateTransactionDetail(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	transactiondetail := getTransactionDetailOr404(db, id, w, r)
	if transactiondetail == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transactiondetail); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&transactiondetail).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, transactiondetail)
}

func DeleteTransactionDetail(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	transactiondetail := getTransactionDetailOr404(db, id, w, r)
	if transactiondetail == nil {
		return
	}
	if err := db.Delete(&transactiondetail).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAllTransactionDetails(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	transactiondetails := []model.TransactionDetailModel{}
	db.Find(&transactiondetails)
	respondJSON(w, http.StatusOK, transactiondetails)
}

func GetTransactionDetailByTransaction(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	transactionID := vars["id_transaction"]
	fmt.Println("bisa ga  " + transactionID)
	transaction := getTransactionOr404(db, transactionID, w, r)
	fmt.Println("bisa ga nih,id = " + transaction.TransactionUuid)
	if transaction == nil {
		return
	}
	transactiondetails := []model.TransactionDetailModel{}
	if err := db.Model(&transaction).Association("TransactionDetails").Find(&transactiondetails).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, transactiondetails)
}

func getTransactionDetailOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.TransactionDetailModel {
	transactiondetail := model.TransactionDetailModel{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&transactiondetail, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &transactiondetail
}
