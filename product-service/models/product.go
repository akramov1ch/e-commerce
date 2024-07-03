package models

type Product struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
}
