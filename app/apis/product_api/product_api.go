package product_api

import (
	"crud-restfull-api-1/app/config"
	"crud-restfull-api-1/app/help"
	"crud-restfull-api-1/app/models"
	"net/http"

	"github.com/gorilla/mux"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	productModel := models.ProductModel{
		Db: db,
	}

	products, err := productModel.FindAll()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	help.ResponseWithJson(w, http.StatusOK, products)
}

func FindById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := config.GetDB()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	productModel := models.ProductModel{
		Db: db,
	}

	product, err := productModel.FindById(id)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	help.ResponseWithJson(w, http.StatusOK, product)
}
