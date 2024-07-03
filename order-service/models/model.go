package models

type Order struct {
    ID         string  `json:"id"`
    UserID     string  `json:"user_id"`
    ProductID  string  `json:"product_id"`
    Quantity   int32   `json:"quantity"`
    Status     string  `json:"status"`
    CreatedAt  string  `json:"created_at"`
    UpdatedAt  string  `json:"updated_at"`
    TotalPrice float64 `json:"total_price"`
}
