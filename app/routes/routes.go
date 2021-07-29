package routes

import (
	"crud-restfull-api-1/app/apis/product_api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	// products
	product := api.PathPrefix("/product").Subrouter()

	product.HandleFunc("/findall", product_api.FindAll).Methods("GET")
	product.HandleFunc("/find/{id}", product_api.FindById).Methods("GET")
	product.HandleFunc("/create", product_api.Create).Methods("POST")

	return router
}
