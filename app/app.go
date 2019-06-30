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
	a.Get("/users/{id}", a.GetUser)
	a.Put("/users/edit/{id}", a.UpdateUser)
	a.Delete("/users/delete/{id}", a.DeleteUser)
	a.Post("/login", a.Login)
	a.Post("/register", a.Register)
	a.Put("/users/confirmation/{id}", a.Confirmation)
	a.Get("/", a.HelloWorld)
	// a.Get("/account/{email}", a.GetAllAccount)
	//routing for handling seller
	a.Post("/seller/post", a.CreateSeller)
	a.Get("/seller/{id}", a.GetSeller)
	a.Put("/seller/edit/{id}", a.UpdateSeller)
	a.Delete("/seller/delete/{id}", a.DeleteSeller)
	//routing for handling buyer
	a.Post("/buyer/post", a.CreateBuyer)
	a.Get("/buyer/{id}", a.GetBuyer)
	a.Put("/buyer/edit/{id}", a.UpdateBuyer)
	a.Delete("/buyer/delete/{id}", a.DeleteBuyer)
	//routing for handling Product
	a.Post("/product/post", a.CreateProduct)
	a.Get("/product/{id}", a.GetProduct)
	a.Put("/product/edit/{id}", a.UpdateProduct)
	a.Delete("/product/delete/{id}", a.DeleteProduct)
	a.Get("/products", a.GetAllProducts)
	a.Get("/product/seller/{id_seller}", a.GetProductBySeller)
	//routing for handling Transaction
	a.Post("/transaction/post", a.CreateTransaction)
	a.Get("/transaction/{id}", a.GetTransaction)
	a.Put("/transaction/edit/{id}", a.UpdateTransaction)
	a.Delete("/transaction/delete/{id}", a.DeleteTransaction)
	a.Get("/transactions", a.GetAllTransactions)
	a.Get("/transaction/seller/{id_seller}", a.GetTransactionBySeller)
	//routing for handling Transaction
	a.Post("/transactiondetail/post", a.CreateTransactionDetail)
	a.Get("/transactiondetail/{id}", a.GetTransactionDetail)
	a.Put("/transactiondetail/edit/{id}", a.UpdateTransactionDetail)
	a.Delete("/transactiondetail/delete/{id}", a.DeleteTransactionDetail)
	a.Get("/transactiondetail/transaction/{id_transaction}", a.GetTransactionDetailByTransaction)
	a.Get("/transactiondetails", a.GetAllTransactionDetails)

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

// Run the app on it's router
func (a *App) Run(host string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(a.Router)
	log.Fatal(http.ListenAndServe(host, handler))
}
