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

func CreateInvoice(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	invoice := model.InvoiceModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&invoice); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&invoice).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, invoice)
}

func GetInvoice(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	invoice := getInvoiceOr404(db, id, w, r)
	if invoice == nil {
		return
	}
	respondJSON(w, http.StatusOK, invoice)
}

func UpdateInvoice(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	invoice := getInvoiceOr404(db, id, w, r)
	if invoice == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&invoice); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&invoice).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, invoice)
}

func DeleteInvoice(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	invoice := getInvoiceOr404(db, id, w, r)
	if invoice == nil {
		return
	}
	if err := db.Delete(&invoice).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAllInvoices(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	invoices := []model.InvoiceModel{}
	db.Find(&invoices)
	respondJSON(w, http.StatusOK, invoices)
}

func GetInvoiceByBuyer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	buyerID := vars["id_buyer"]
	buyer := getBuyerOr404(db, buyerID, w, r)
	if buyer == nil {
		return
	}
	invoices := []model.InvoiceModel{}
	if err := db.Model(&buyer).Association("Invoices").Find(&invoices).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, invoices)
}

func getInvoiceOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.InvoiceModel {
	invoice := model.InvoiceModel{}

	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&invoice, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &invoice
}
