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
	NoHp         string
	Birtday      time.Time
	JenisKelamin bool
	PathPhoto    string
	Role         int
	Token        string `json:"token"`
	VerifyEmail  bool
}

type SellerModel struct {
	gorm.Model
	AccountID    uint   `gorm:"unique" json:"account_id"`
	Username     string `gorm:"unique"`
	Alamat       string
	Logo         string
	Kodepos      int
	NamaBank     string
	PemilikBank  string
	NoRekening   string
	Wallet       int
	LastLogin    time.Time
	Deskripsi    string
	Products     []ProductModel     `gorm:"ForeignKey:SellerID" json:"products"`
	Transactions []TransactionModel `gorm:"ForeignKey:SellerID" json:"transactions"`
}

type BuyerModel struct {
	gorm.Model
	AccountID     uint `gorm:"unique" json:"account_id"`
	DefaultAlamat int
	NoHpPenerima  string
	NamaAlamat    string
	Alamat        string
	KodePos       int
	Penerima      string
}
type ProductModel struct {
	gorm.Model
	SellerID uint `json:"seller_id"`
	Kategori int
	Nama     string
	Harga    float32 `sql:"type:decimal(10,2);"`
	Stok     int
	Terjual  int
	Dilihat  int
	Berat    float32 `sql:"type:decimal(10,2);"`
	Length   float32 `sql:"type:decimal(10,2);"`
	Width    float32 `sql:"type:decimal(10,2);"`
	Height   float32 `sql:"type:decimal(10,2);"`
	PathFoto string
}

type TransactionModel struct {
	gorm.Model
	SellerID           uint `json:"seller_id"`
	TransactionStatus  int
	TransactionUuid    string
	TransactionDetails []TransactionDetailModel `gorm:"ForeignKey:TransactionID" json:"transactiondetails"`
}

type TransactionDetailModel struct {
	gorm.Model
	TransactionID uint `json:"transaction_id"`
	ProductID     uint `json:"product_id"`
	Quantity      int
	ItemPrice     int
	TextCard      string
	Note          string
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Account{}, SellerModel{}, BuyerModel{}, ProductModel{}, TransactionModel{}, TransactionDetailModel{})
	db.Model(&SellerModel{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&BuyerModel{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&ProductModel{}).AddForeignKey("seller_id", "seller_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionModel{}).AddForeignKey("seller_id", "seller_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionDetailModel{}).AddForeignKey("transaction_id", "transaction_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionDetailModel{}).AddForeignKey("product_id", "product_models(id)", "RESTRICT", "RESTRICT")

	return db
}
