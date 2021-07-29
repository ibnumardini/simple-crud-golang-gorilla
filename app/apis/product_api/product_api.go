package product_api

import (
	"crud-restfull-api-1/app/config"
	"crud-restfull-api-1/app/entities"
	"crud-restfull-api-1/app/help"
	"crud-restfull-api-1/app/models"
	"encoding/json"
	"net/http"
	"strconv"

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

func Create(w http.ResponseWriter, r *http.Request) {

	var product entities.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	db, err := config.GetDB()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	productModel := models.ProductModel{
		Db: db,
	}

	err = productModel.Create(&product)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	help.ResponseWithJson(w, http.StatusCreated, product)
}

func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var product entities.Product

	product.Id, _ = strconv.ParseInt(id, 10, 64)

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	db, err := config.GetDB()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	productModel := models.ProductModel{
		Db: db,
	}

	err = productModel.Update(&product)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	help.ResponseWithJson(w, http.StatusCreated, product)
}
