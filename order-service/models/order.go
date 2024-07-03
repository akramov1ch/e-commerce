package models

type Order struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
	OrderedAt string `json:"ordered_at"`
}
