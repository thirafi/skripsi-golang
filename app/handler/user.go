package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.Account{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Account{}

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

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	account := model.Account{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	account1 := getUserOr404email(db, account.Email, w, r)
	fmt.Println(account.Nama)
	arr := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(account1.Password))
	if arr != nil && arr == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		respondError(w, http.StatusNotFound, arr.Error())
	}
	//Worked! Logged In
	account1.Password = ""

	//Create JWT token
	tk := model.Token{Id_user: account1.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account1.Token = tokenString
	fmt.Println("yang login " + account1.Email)

	respondJSON(w, http.StatusCreated, account1)
}
func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	account := model.Account{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	if err := db.Save(&account).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	tk := model.Token{Id_user: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = "" //delete password
	respondJSON(w, http.StatusCreated, account)
}

func Confirmation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
