package routes

import (
	"crud-restfull-api-1/app/apis/product_api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	product := api.PathPrefix("/product").Subrouter()

	// products
	product.HandleFunc("/findall", product_api.FindAll).Methods("GET")
	product.HandleFunc("/find/{id}", product_api.FindById).Methods("GET")

	return router
}
