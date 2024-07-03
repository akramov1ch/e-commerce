package order

import (
    "database/sql"
    "errors"
)

type Repository interface {
    CreateOrder(order *Order) error
    GetOrder(id string) (*Order, error)
    DeleteOrder(id string) error
    ListOrders() ([]*Order, error)
}

type PostgresRepository struct {
    db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
    return &PostgresRepository{db: db}
}

func (r *PostgresRepository) CreateOrder(order *Order) error {
    query := `INSERT INTO orders (id, user_id, product_id, quantity, status, total_price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING created_at, updated_at`
    err := r.db.QueryRow(query, order.ID, order.UserID, order.ProductID, order.Quantity, order.Status, order.TotalPrice).Scan(&order.CreatedAt, &order.UpdatedAt)
    return err
}

func (r *PostgresRepository) GetOrder(id string) (*Order, error) {
    query := `SELECT id, user_id, product_id, quantity, status, total_price, created_at, updated_at FROM orders WHERE id=$1`
    row := r.db.QueryRow(query, id)

    var order Order
    if err := row.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("order not found")
        }
        return nil, err
    }
    return &order, nil
}


func (r *PostgresRepository) DeleteOrder(id string) error {
    query := `DELETE FROM orders WHERE id=$1`
    _, err := r.db.Exec(query, id)
    return err
}

func (r *PostgresRepository) ListOrders() ([]*Order, error) {
    query := `SELECT id, user_id, product_id, quantity, status, total_price, created_at, updated_at FROM orders`
    rows, err := r.db.Query(query)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var orders []*Order
    for rows.Next() {
        var order Order
        if err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt); err != nil {
            return nil, err
        }
        orders = append(orders, &order)
    }
    return orders, nil
}
