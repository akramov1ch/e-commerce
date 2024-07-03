package user

import (
	"database/sql"
	"errors"
	"time"
)
var ErrUserNotFound = errors.New("user not found")
type Repository interface {
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) CreateUser(user *User) error {
	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING created_at, updated_at`
	err := r.db.QueryRow(query, user.ID, user.Name, user.Email, user.Password).Scan(&user.CreatedAt, &user.UpdatedAt)
	return err
}

func (r *PostgresRepository) GetUser(id string) (*User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresRepository) UpdateUser(user *User) error {
	newtime := time.Now()
	query := `UPDATE users SET name = $2, email = $3, password=$4, updated_at=$5  WHERE id = $1 RETURNING created_at, updated_at`
	err := r.db.QueryRow(query, user.ID, user.Name, user.Email, user.Password, newtime).Scan(&user.CreatedAt, &user.UpdatedAt)
	return err
}

func (r *PostgresRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err!= nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
