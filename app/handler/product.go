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

func CreateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := model.ProductModel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, product)
}

func GetProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	product := getProductOr404(db, id, w, r)
	if product == nil {
		return
	}
	respondJSON(w, http.StatusOK, product)
}

func UpdateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	product := getProductOr404(db, id, w, r)
	if product == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, product)
}

func DeleteProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	product := getProductOr404(db, id, w, r)
	if product == nil {
		return
	}
	if err := db.Delete(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.ProductModel{}
	db.Find(&products)
	respondJSON(w, http.StatusOK, products)
}

func GetProductBySeller(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sellerID := vars["id_seller"]
	seller := getSellerOr404(db, sellerID, w, r)
	if seller == nil {
		return
	}
	fmt.Println("bisa ga nih,id = " + seller.Username)
	products := []model.ProductModel{}
	if err := db.Model(&seller).Association("Products").Find(&products).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, products)
}

func getProductOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.ProductModel {
	product := model.ProductModel{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&product, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &product
}
