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
	a.Post("/users", a.CreateUser)
	a.Get("/users/{id}", a.GetUser)
	a.Put("/users/{id}", a.UpdateUser)
	a.Delete("/users/{id}", a.DeleteUser)
	a.Post("/login", a.Login)
	a.Post("/register", a.Register)
	a.Put("/users/confirmation/{id}", a.Confirmation)
	a.Get("/", a.HelloWorld)
	// a.Get("/account/{email}", a.GetAllAccount)
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

// Run the app on it's router
func (a *App) Run(host string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(a.Router)
	log.Fatal(http.ListenAndServe(host, handler))
}
