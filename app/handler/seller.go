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
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getUserOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Account {
	user := model.Account{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&user, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
func getUserOr404email(db *gorm.DB, email string, w http.ResponseWriter, r *http.Request) *model.Account {
	user := model.Account{}
	// ud, err := strconv.ParseUint(id, 10, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// ad := uint(ud)
	if err := db.Where("email=?", email).First(&user).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
