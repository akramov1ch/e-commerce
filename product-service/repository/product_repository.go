package repository

import (
	"database/sql"

	"product-service/config"
	"product-service/models"

	_ "github.com/lib/pq"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id string) error
	GetProductByID(id string) (*models.Product, error)
	GetProducts() ([]*models.Product, error)
}

type PostgresRepository struct {
	db *sql.DB
}


func NewPostgresRepository() (ProductRepository, error) {
	var config config.Config
	connStr := config.DatabaseURL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) CreateProduct(product *models.Product) error {
	query := `INSERT INTO products (id, product_name, description) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, product.ID, product.ProductName, product.Description)
	return err
}

func (r *PostgresRepository) UpdateProduct(product *models.Product) error {
	query := `UPDATE products SET product_name = $1, description = $2 WHERE id = $3`
	_, err := r.db.Exec(query, product.ProductName, product.Description, product.ID)
	return err
}

func (r *PostgresRepository) DeleteProduct(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) GetProductByID(id string) (*models.Product, error) {
	query := `SELECT id, product_name, description FROM products WHERE id = $1`
	row := r.db.QueryRow(query, id)
	product := &models.Product{}
	err := row.Scan(&product.ID, &product.ProductName, &product.Description)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *PostgresRepository) GetProducts() ([]*models.Product, error) {
	query := `SELECT id, product_name, description FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(&product.ID, &product.ProductName, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
