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

func CreateTransaction(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	transaction := model.TransactionModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transaction); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&transaction).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, transaction)
}

func GetTransaction(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	transaction := getTransactionOr404(db, id, w, r)
	if transaction == nil {
		return
	}
	respondJSON(w, http.StatusOK, transaction)
}

func UpdateTransaction(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	transaction := getTransactionOr404(db, id, w, r)
	if transaction == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transaction); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&transaction).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, transaction)
}

func DeleteTransaction(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	transaction := getTransactionOr404(db, id, w, r)
	if transaction == nil {
		return
	}
	if err := db.Delete(&transaction).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAllTransactions(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	transactions := []model.TransactionModel{}
	db.Find(&transactions)
	respondJSON(w, http.StatusOK, transactions)
}

func GetTransactionBySeller(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sellerID := vars["id_seller"]
	seller := getSellerOr404(db, sellerID, w, r)
	if seller == nil {
		return
	}
	transactions := []model.TransactionModel{}
	if err := db.Model(&seller).Association("Transactions").Find(&transactions).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, transactions)
}

func getTransactionOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.TransactionModel {
	transaction := model.TransactionModel{}

	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&transaction, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	fmt.Println("bisa ga nih,id = " + transaction.TransactionUuid)
	return &transaction
}
