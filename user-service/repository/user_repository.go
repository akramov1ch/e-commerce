package repository

import (
	"context"
	"database/sql"
	"log"

	"user-service/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := `UPDATE users SET username=$1, email=$2 WHERE id=$3`
	_, err := r.DB.ExecContext(ctx, query, user.Username, user.Email, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, username, email FROM users WHERE id=$1`
	user := &models.User{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}
	return user, nil
}
