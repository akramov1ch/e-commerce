package service

import (
    "user/models"
    "user/repository"

    "github.com/google/uuid"
)

type Service interface {
    CreateUser(name, email, password string) (*models.User, error)
    GetUser(id string) (*models.User, error)
    UpdateUser(id, name, email string) (*models.User, error)
    DeleteUser(id string) error
}

type UserService struct {
    repo repository.Repository
}

func NewUserService(repo repository.Repository) Service {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email, password string) (*models.User, error) {
    user := &models.User{
        ID:       uuid.New().String(),
        Name:     name,
        Email:    email,
        Password: password,
    }
    if err := s.repo.CreateUser(user); err != nil {
        return nil, err
    }
    return user, nil
}

func (s *UserService) GetUser(id string) (*models.User, error) {
    return s.repo.GetUser(id)
}

func (s *UserService) UpdateUser(id, name, email string) (*models.User, error) {
    user, err := s.repo.GetUser(id)
    if err != nil {
        return nil, err
    }
    user.Name = name
    user.Email = email
    if err := s.repo.UpdateUser(user); err != nil {
        return nil, err
    }
    return user, nil
}

func (s *UserService) DeleteUser(id string) error {
    return s.repo.DeleteUser(id)
}
