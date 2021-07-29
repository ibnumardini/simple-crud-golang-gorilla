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
		return
	}

	productModel := models.ProductModel{
		Db: db,
	}

	products, err := productModel.FindAll()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	help.ResponseWithJson(w, http.StatusOK, products)
}

func FindById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := config.GetDB()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productModel := models.ProductModel{
		Db: db,
	}

	product, err := productModel.FindById(id)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	help.ResponseWithJson(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {

	var product entities.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	db, err := config.GetDB()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productModel := models.ProductModel{
		Db: db,
	}

	err = productModel.Create(&product)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
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
		return
	}

	db, err := config.GetDB()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productModel := models.ProductModel{
		Db: db,
	}

	err = productModel.Update(&product)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	help.ResponseWithJson(w, http.StatusCreated, product)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := config.GetDB()
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productModel := models.ProductModel{
		Db: db,
	}

	result, err := productModel.Delete(id)
	if err != nil {
		help.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if result <= 0 {
		help.ResponseWithJson(w, http.StatusBadRequest, map[string]string{
			"message": "Product Not Found!",
		})
		return
	}

	help.ResponseWithJson(w, http.StatusCreated, map[string]string{
		"message": "Product Deleted!",
	})
}
