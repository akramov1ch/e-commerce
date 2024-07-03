package repository

import (
	"database/sql"
	"log"
	"order-service/models"
	"time"

	"github.com/google/uuid"
)

type OrderRepository interface {
	CreateOrder(userID, productID string) (*models.Order, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(userID, productID string) (*models.Order, error) {
	id := uuid.New().String()
	orderedAt := time.Now().Format(time.RFC3339)

	_, err := r.db.Exec("INSERT INTO orders (id, user_id, product_id, ordered_at) VALUES ($1, $2, $3, $4)",
		id, userID, productID, orderedAt)
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return nil, err
	}

	return &models.Order{
		ID:        id,
		UserID:    userID,
		ProductID: productID,
		OrderedAt: orderedAt,
	}, nil
}
