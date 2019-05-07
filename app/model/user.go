package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Token struct {
	Id_user uint
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	gorm.Model
	Email        string `gorm:"unique" json:"email"`
	Nama         string `json:"nama"`
	Password     string `json:"password"`
	Code         string
	NoHp         int
	Birtday      time.Time
	JenisKelamin bool
	PathPhoto    string
	Role         int
	Token        string `json:"token"`
	VerifyEmail  bool
}

type SellerModel struct {
	gorm.Model
	Account     Account
	Username    string `gorm:"unique"`
	Alamat      string
	Logo        string
	Kodepos     int
	NamaBank    string
	PemilikBank string
	NoRekening  string
	Wallet      int
	LastLogin   time.Time
	Deskripsi   string
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Account{}, SellerModel{})
	db.Model(&Account{}).Related(&SellerModel{})
	return db
}
