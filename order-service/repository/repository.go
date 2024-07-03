package repository

import (
    "database/sql"
    "errors"
    "order-service/models"
)

type Repository interface {
    CreateOrder(order *models.Order) error
    GetOrder(id string) (*models.Order, error)
    DeleteOrder(id string) error
    ListOrders() ([]*models.Order, error)
}

type PostgresRepository struct {
    db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
    return &PostgresRepository{db: db}
}

func (r *PostgresRepository) CreateOrder(order *models.Order) error {
    query := `INSERT INTO orders (id, user_id, product_id, quantity, status, total_price, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())`
    _, err := r.db.Exec(query, order.ID, order.UserID, order.ProductID, order.Quantity, order.Status, order.TotalPrice)
    return err
}

func (r *PostgresRepository) GetOrder(id string) (*models.Order, error) {
    query := `SELECT id, user_id, product_id, quantity, status, total_price, created_at, updated_at
              FROM orders WHERE id = $1`
    row := r.db.QueryRow(query, id)

    var order models.Order
    err := row.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("order not found")
        }
        return nil, err
    }
    return &order, nil
}

func (r *PostgresRepository) DeleteOrder(id string) error {
    query := `DELETE FROM orders WHERE id = $1`
    _, err := r.db.Exec(query, id)
    return err
}

func (r *PostgresRepository) ListOrders() ([]*models.Order, error) {
    query := `SELECT id, user_id, product_id, quantity, status, total_price, created_at, updated_at
              FROM orders`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []*models.Order
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt); err != nil {
            return nil, err
        }
        orders = append(orders, &order)
    }

    return orders, nil
}
