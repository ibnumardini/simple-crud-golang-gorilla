package entities

import "fmt"

type Product struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

func (product Product) ToString() string {
	return fmt.Sprintf("id: %d\n name: %s\n price: %0.1f\n quantity:%d", product.Id, product.Name, product.Price, product.Quantity)
}
