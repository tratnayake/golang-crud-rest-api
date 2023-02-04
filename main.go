package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}
func main() {
	//Load configs from config.json using Viper
	LoadAppConfig()

	//Initiatlize DB
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize Router
	router := mux.NewRouter().StrictSlash(true)

	//Register Routes
	RegisterProductRoutes(router)

	//Start the serer
	log.Println(fmt.Sprintf("Starting server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
