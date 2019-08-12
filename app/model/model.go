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
	Username     string `gorm:"unique"` //wajib
	Alamat       string //wajib
	Logo         string
	Kodepos      int //wajib
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
	Invoices      []InvoiceModel `gorm:"ForeignKey:BuyyerID" json:"invoices"`
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
	BoxesID       uint `json:"boxes_id"`
	BoxpaperID    uint `json:"boxpaper_id"`
	RibbonID      uint `json:"ribbon_id"`
	Quantity      int
	ItemPrice     int
	TextCard      string
	Note          string
}

type InvoiceModel struct {
	gorm.Model
	TransactionID  uint `json:"transaction_id"`
	BuyyerID       uint `json:"buyyer_id"`
	PaymentID      uint `json:"payment_id"`
	InvoiceStatus  int
	InvoiceNo      string
	Discount       int
	PaymentMethod  string
	PaymentReceipt string
	UniqueCode     int
}

type PaymentModel struct {
	gorm.Model
	NamaBank   string
	NoRekening int
	AtasNama   string
}

type BoxesModel struct {
	gorm.Model
	Code       string
	Name       string
	BoxPrice   int
	Length     int
	Width      int
	Height     int
	BoxPicture string
}

type BoxpaperModel struct {
	gorm.Model
	Code             string
	Name             string
	BoxpapperPrice   int
	BoxpapperPicture string
}

type RibbonModel struct {
	gorm.Model
	Code          string
	Name          string
	RibbonPrice   int
	RibbonPicture string
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Account{}, SellerModel{}, BuyerModel{}, ProductModel{}, TransactionModel{}, TransactionDetailModel{}, InvoiceModel{}, PaymentModel{}, BoxesModel{}, BoxpaperModel{}, RibbonModel{})
	db.Model(&SellerModel{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&BuyerModel{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&ProductModel{}).AddForeignKey("seller_id", "seller_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionModel{}).AddForeignKey("seller_id", "seller_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionDetailModel{}).AddForeignKey("transaction_id", "transaction_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionDetailModel{}).AddForeignKey("product_id", "product_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionDetailModel{}).AddForeignKey("boxes_id", "boxes_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionDetailModel{}).AddForeignKey("boxpaper_id", "boxpaper_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&TransactionDetailModel{}).AddForeignKey("ribbon_id", "ribbon_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&InvoiceModel{}).AddForeignKey("transaction_id", "transaction_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&InvoiceModel{}).AddForeignKey("buyyer_id", "buyer_models(id)", "RESTRICT", "RESTRICT")
	db.Model(&InvoiceModel{}).AddForeignKey("payment_id", "payment_models(id)", "RESTRICT", "RESTRICT")
	return db
}
