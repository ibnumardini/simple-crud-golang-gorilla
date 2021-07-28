package models

import (
	"crud-restfull-api-1/app/entities"
	"database/sql"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (products []entities.Product, err error) {

	rows, err := productModel.Db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var id int64
		var name string
		var price float64
		var quantity int64

		err := rows.Scan(&id, &name, &price, &quantity)
		if err != nil {
			return nil, err
		}

		product := entities.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}

		products = append(products, product)
	}

	return products, nil
}

func (productModel ProductModel) FindById(id string) (products []entities.Product, err error) {

	rows, err := productModel.Db.Query("SELECT * FROM products WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var id int64
		var name string
		var price float64
		var quantity int64

		err := rows.Scan(&id, &name, &price, &quantity)
		if err != nil {
			return nil, err
		}

		product := entities.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}

		products = append(products, product)
	}

	return products, nil
}
