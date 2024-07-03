package repository

import (
    "database/sql"
    "errors"
    md "product/model"
)

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
    AddProduct(product *md.Product) error
    GetProduct(id string) (*md.Product, error)
    UpdateProduct(product *md.Product) error
    DeleteProduct(id string) error
    ListProducts() ([]*md.Product, error)
}

type PostgresRepository struct {
    db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
    return &PostgresRepository{db: db}
}

func (r *PostgresRepository) AddProduct(product *md.Product) error {
    query := `INSERT INTO products (id, name, description, price, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING created_at, updated_at`
    err := r.db.QueryRow(query, product.ID, product.Name, product.Description, product.Price, product.Stock).Scan(&product.CreatedAt, &product.UpdatedAt)
    return err
}

func (r *PostgresRepository) GetProduct(id string) (*md.Product, error) {
    query := `SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = $1`
    row := r.db.QueryRow(query, id)

    var product md.Product
    if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("product not found")
        }
        return nil, err
    }
    return &product, nil
}

func (r *PostgresRepository) UpdateProduct(product *md.Product) error {
    query := `UPDATE products SET name = $2, description = $3, price = $4, stock = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING created_at, updated_at`
    err := r.db.QueryRow(query, product.ID, product.Name, product.Description, product.Price, product.Stock).Scan(&product.CreatedAt, &product.UpdatedAt)
    return err
}

func (r *PostgresRepository) DeleteProduct(id string) error {
    query := `DELETE FROM products WHERE id = $1`
    result, err := r.db.Exec(query, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return ErrUserNotFound
    }
    return nil
}

func (r *PostgresRepository) ListProducts() ([]*md.Product, error) {
    query := `SELECT id, name, description, price, stock, created_at, updated_at FROM products`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []*md.Product
    for rows.Next() {
        var product md.Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
            return nil, err
        }
        products = append(products, &product)
    }
    return products, nil
}
