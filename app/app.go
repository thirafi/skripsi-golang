package app

import (
	"fmt"
	"log"
	"net/http"

	"../config"
	"./handler"
	"./model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()

}

func (a *App) setRouters() {
	//routing for handling user
	a.Get("/users", a.GetAllUsers)
	// a.Post("/users/post", a.CreateUser)
	a.Get("/user/{id}", a.GetUser)
	a.Put("/user/edit/{id}", a.UpdateUser)
	a.Delete("/user/delete/{id}", a.DeleteUser)
	a.Post("/login", a.Login)
	a.Post("/register", a.Register)
	a.Put("/users/confirmation/{id}", a.Confirmation)
	a.Get("/", a.HelloWorld)
	a.Post("/user/upload", a.UploadUser)

	// a.Get("/account/{email}", a.GetAllAccount)
	//routing for handling seller
	a.Post("/seller/post", a.CreateSeller)
	a.Get("/seller/{id}", a.GetSeller)
	a.Put("/seller/edit/{id}", a.UpdateSeller)
	a.Delete("/seller/delete/{id}", a.DeleteSeller)
	a.Post("/seller/upload", a.UploadSeller)
	//routing for handling buyer
	a.Post("/address/post", a.CreateBuyer)
	a.Get("/address/{id}", a.GetBuyer)
	a.Put("/address/edit/{id}", a.UpdateBuyer)
	a.Delete("/address/delete/{id}", a.DeleteBuyer)
	a.Get("/address/user/{id}", a.GetAddressByUser)
	//routing for handling Product
	a.Post("/product/post", a.CreateProduct)
	a.Get("/product/{id}", a.GetProduct)
	a.Put("/product/edit/{id}", a.UpdateProduct)
	a.Delete("/product/delete/{id}", a.DeleteProduct)
	a.Get("/products", a.GetAllProducts)
	a.Get("/product/seller/{id_seller}", a.GetProductBySeller)
	a.Post("/product/upload", a.UploadProduct)
	a.Get("/pencarian/{key}", a.Pencarian)
	//routing for handling Transaction
	a.Post("/transaction/post", a.CreateTransaction)
	a.Get("/transaction/{id}", a.GetTransaction)
	a.Put("/transaction/edit/{id}", a.UpdateTransaction)
	a.Delete("/transaction/delete/{id}", a.DeleteTransaction)
	a.Get("/transactions", a.GetAllTransactions)
	a.Put("/transaction/productarrived/{id}", a.ProductArrived)
	a.Get("/transaction/seller/{id_seller}", a.GetTransactionBySeller)
	//routing for handling Transaction detail
	a.Post("/transactiondetail/post", a.CreateTransactionDetail)
	a.Get("/transactiondetail/{id}", a.GetTransactionDetail)
	a.Put("/transactiondetail/edit/{id}", a.UpdateTransactionDetail)
	a.Delete("/transactiondetail/delete/{id}", a.DeleteTransactionDetail)
	a.Get("/transactiondetail/transaction/{id_transaction}", a.GetTransactionDetailByTransaction)
	a.Get("/transactiondetails", a.GetAllTransactionDetails)
	//routing for handling Invoice
	a.Post("/invoice/post", a.CreateInvoice)
	a.Get("/invoice/{id}", a.GetInvoice)
	a.Put("/invoice/edit/{id}", a.UpdateInvoice)
	a.Delete("/invoice/delete/{id}", a.DeleteInvoice)
	a.Get("/invoice/order/{id_buyer}", a.GetInvoiceByBuyer)
	a.Get("/invoices", a.GetAllInvoices)
	a.Post("/invoice/upload", a.UploadInvoice)
	//routing for handling Payment
	a.Post("/payment/post", a.CreatePayment)
	a.Get("/payment/{id}", a.GetPayment)
	a.Put("/payment/edit/{id}", a.UpdatePayment)
	a.Delete("/payment/delete/{id}", a.DeletePayment)
	a.Get("/payments", a.GetAllPayments)
	//routing for handling Boxes
	a.Post("/box/post", a.CreateBoxes)
	a.Get("/box/{id}", a.GetBoxes)
	a.Put("/box/edit/{id}", a.UpdateBoxes)
	a.Delete("/box/delete/{id}", a.DeleteBoxes)
	a.Get("/boxes", a.GetAllBoxes)
	a.Post("/box/upload", a.UploadBox)
	//routing for handling Boxes
	a.Post("/boxpaper/post", a.CreateBoxpaper)
	a.Get("/boxpaper/{id}", a.GetBoxpaper)
	a.Put("/boxpaper/edit/{id}", a.UpdateBoxpaper)
	a.Delete("/boxpaper/delete/{id}", a.DeleteBoxpaper)
	a.Get("/boxpapers", a.GetAllBoxpaper)
	a.Post("/boxpaper/upload", a.UploadPaper)
	//routing for handling Boxes
	a.Post("/ribbon/post", a.CreateRibbon)
	a.Get("/ribbon/{id}", a.GetRibbon)
	a.Put("/ribbon/edit/{id}", a.UpdateRibbon)
	a.Delete("/ribbon/delete/{id}", a.DeleteRibbon)
	a.Get("/ribbons", a.GetAllRibbon)
	a.Post("/ribbon/upload", a.UploadRibbon)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Use(JwtAuthentication)
	a.Router.HandleFunc(path, f).Methods("GET")

}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Use(JwtAuthentication)
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Use(JwtAuthentication)
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.Use(JwtAuthentication)
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

//--------------------------------------------------------------------------------------//

func (a *App) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa ini dia ")
}

// Handlers to manage User Data
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}

func (a *App) Confirmation(w http.ResponseWriter, r *http.Request) {
	handler.Confirmation(a.DB, w, r)
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	handler.Login(a.DB, w, r)

}
func (a *App) Register(w http.ResponseWriter, r *http.Request) {
	handler.Register(a.DB, w, r)
}

func (a *App) UploadUser(w http.ResponseWriter, r *http.Request) {
	handler.UploadUser(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Seller Data
func (a *App) CreateSeller(w http.ResponseWriter, r *http.Request) {
	handler.CreateSeller(a.DB, w, r)
}

func (a *App) GetSeller(w http.ResponseWriter, r *http.Request) {
	handler.GetSeller(a.DB, w, r)
}

func (a *App) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	handler.UpdateSeller(a.DB, w, r)
}

func (a *App) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	handler.DeleteSeller(a.DB, w, r)
}

func (a *App) UploadSeller(w http.ResponseWriter, r *http.Request) {
	handler.UploadSeller(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Buyer Data
func (a *App) CreateBuyer(w http.ResponseWriter, r *http.Request) {
	handler.CreateBuyer(a.DB, w, r)
}

func (a *App) GetBuyer(w http.ResponseWriter, r *http.Request) {
	handler.GetBuyer(a.DB, w, r)
}

func (a *App) UpdateBuyer(w http.ResponseWriter, r *http.Request) {
	handler.UpdateBuyer(a.DB, w, r)
}

func (a *App) DeleteBuyer(w http.ResponseWriter, r *http.Request) {
	handler.DeleteBuyer(a.DB, w, r)
}

func (a *App) GetAddressByUser(w http.ResponseWriter, r *http.Request) {
	handler.GetAddressByUser(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Product Data
func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(a.DB, w, r)
}

func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	handler.GetProduct(a.DB, w, r)
}

func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	handler.UpdateProduct(a.DB, w, r)
}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	handler.DeleteProduct(a.DB, w, r)
}

func (a *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProducts(a.DB, w, r)
}

func (a *App) GetProductBySeller(w http.ResponseWriter, r *http.Request) {
	handler.GetProductBySeller(a.DB, w, r)
}

func (a *App) UploadProduct(w http.ResponseWriter, r *http.Request) {
	handler.UploadProduct(a.DB, w, r)
}
func (a *App) Pencarian(w http.ResponseWriter, r *http.Request) {
	handler.Pencarian(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Transaction Data
func (a *App) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	handler.CreateTransaction(a.DB, w, r)
}

func (a *App) GetTransaction(w http.ResponseWriter, r *http.Request) {
	handler.GetTransaction(a.DB, w, r)
}

func (a *App) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTransaction(a.DB, w, r)
}

func (a *App) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	handler.DeleteTransaction(a.DB, w, r)
}

func (a *App) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTransactions(a.DB, w, r)
}

func (a *App) GetTransactionBySeller(w http.ResponseWriter, r *http.Request) {
	handler.GetTransactionBySeller(a.DB, w, r)
}

func (a *App) ProductArrived(w http.ResponseWriter, r *http.Request) {
	handler.ProductArrived(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Transaction Data
func (a *App) CreateTransactionDetail(w http.ResponseWriter, r *http.Request) {
	handler.CreateTransactionDetail(a.DB, w, r)
}

func (a *App) GetTransactionDetail(w http.ResponseWriter, r *http.Request) {
	handler.GetTransactionDetail(a.DB, w, r)
}

func (a *App) UpdateTransactionDetail(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTransactionDetail(a.DB, w, r)
}

func (a *App) DeleteTransactionDetail(w http.ResponseWriter, r *http.Request) {
	handler.DeleteTransactionDetail(a.DB, w, r)
}

func (a *App) GetAllTransactionDetails(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTransactionDetails(a.DB, w, r)
}

func (a *App) GetTransactionDetailByTransaction(w http.ResponseWriter, r *http.Request) {
	handler.GetTransactionDetailByTransaction(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Invoice Data
func (a *App) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	handler.CreateInvoice(a.DB, w, r)
}

func (a *App) GetInvoice(w http.ResponseWriter, r *http.Request) {
	handler.GetInvoice(a.DB, w, r)
}

func (a *App) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	handler.UpdateInvoice(a.DB, w, r)
}

func (a *App) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	handler.DeleteInvoice(a.DB, w, r)
}

func (a *App) GetAllInvoices(w http.ResponseWriter, r *http.Request) {
	handler.GetAllInvoices(a.DB, w, r)
}

func (a *App) GetInvoiceByBuyer(w http.ResponseWriter, r *http.Request) {
	handler.GetInvoiceByBuyer(a.DB, w, r)
}

func (a *App) UploadInvoice(w http.ResponseWriter, r *http.Request) {
	handler.UploadInvoice(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Payment Data
func (a *App) CreatePayment(w http.ResponseWriter, r *http.Request) {
	handler.CreatePayment(a.DB, w, r)
}

func (a *App) GetPayment(w http.ResponseWriter, r *http.Request) {
	handler.GetPayment(a.DB, w, r)
}

func (a *App) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	handler.UpdatePayment(a.DB, w, r)
}

func (a *App) DeletePayment(w http.ResponseWriter, r *http.Request) {
	handler.DeletePayment(a.DB, w, r)
}

func (a *App) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPayments(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Payment Data
func (a *App) CreateBoxes(w http.ResponseWriter, r *http.Request) {
	handler.CreateBoxes(a.DB, w, r)
}

func (a *App) GetBoxes(w http.ResponseWriter, r *http.Request) {
	handler.GetBoxes(a.DB, w, r)
}

func (a *App) UpdateBoxes(w http.ResponseWriter, r *http.Request) {
	handler.UpdateBoxes(a.DB, w, r)
}

func (a *App) DeleteBoxes(w http.ResponseWriter, r *http.Request) {
	handler.DeleteBoxes(a.DB, w, r)
}

func (a *App) GetAllBoxes(w http.ResponseWriter, r *http.Request) {
	handler.GetAllBoxes(a.DB, w, r)
}

func (a *App) UploadBox(w http.ResponseWriter, r *http.Request) {
	handler.UploadBox(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Payment Data
func (a *App) CreateBoxpaper(w http.ResponseWriter, r *http.Request) {
	handler.CreateBoxpaper(a.DB, w, r)
}

func (a *App) GetBoxpaper(w http.ResponseWriter, r *http.Request) {
	handler.GetBoxpaper(a.DB, w, r)
}

func (a *App) UpdateBoxpaper(w http.ResponseWriter, r *http.Request) {
	handler.UpdateBoxpaper(a.DB, w, r)
}

func (a *App) DeleteBoxpaper(w http.ResponseWriter, r *http.Request) {
	handler.DeleteBoxpaper(a.DB, w, r)
}

func (a *App) GetAllBoxpaper(w http.ResponseWriter, r *http.Request) {
	handler.GetAllBoxpaper(a.DB, w, r)
}

func (a *App) UploadPaper(w http.ResponseWriter, r *http.Request) {
	handler.UploadPaper(a.DB, w, r)
}

//////////////////////////////////////////////////////////////////////////////
// Handlers to manage Payment Data
func (a *App) CreateRibbon(w http.ResponseWriter, r *http.Request) {
	handler.CreateRibbon(a.DB, w, r)
}

func (a *App) GetRibbon(w http.ResponseWriter, r *http.Request) {
	handler.GetRibbon(a.DB, w, r)
}

func (a *App) UpdateRibbon(w http.ResponseWriter, r *http.Request) {
	handler.UpdateRibbon(a.DB, w, r)
}

func (a *App) DeleteRibbon(w http.ResponseWriter, r *http.Request) {
	handler.DeleteRibbon(a.DB, w, r)
}

func (a *App) GetAllRibbon(w http.ResponseWriter, r *http.Request) {
	handler.GetAllRibbon(a.DB, w, r)
}
func (a *App) UploadRibbon(w http.ResponseWriter, r *http.Request) {
	handler.UploadRibbon(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(a.Router)
	log.Fatal(http.ListenAndServe(host, handler))
}
