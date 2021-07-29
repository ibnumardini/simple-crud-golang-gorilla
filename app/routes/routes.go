package routes

import (
	"crud-restfull-api-1/app/apis/product_api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	// products
	api.HandleFunc("/product", product_api.FindAll).Methods("GET")
	api.HandleFunc("/product/{id}", product_api.FindById).Methods("GET")
	api.HandleFunc("/product", product_api.Create).Methods("POST")
	api.HandleFunc("/product/{id}", product_api.Update).Methods("PUT")
	api.HandleFunc("/product/{id}", product_api.Delete).Methods("DELETE")

	return router
}
