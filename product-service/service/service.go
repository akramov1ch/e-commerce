package service

import (
    "errors"
    md "product/model"
    rp "product/repository"

    "github.com/google/uuid"
)

type Service interface {
    AddProduct(name, description string, price float32, stock int32) (*md.Product, error)
    GetProduct(id string) (*md.Product, error)
    UpdateProduct(id, name, description string, price float32, stock int32) (*md.Product, error)
    DeleteProduct(id string) error
    ListProducts() ([]*md.Product, error)
}

var ErrUserNotFound = errors.New("user not found")

type ProductService struct {
    repo rp.Repository
}

func NewProductService(repo rp.Repository) Service {
    return &ProductService{repo: repo}
}

func (s *ProductService) AddProduct(name, description string, price float32, stock int32) (*md.Product, error) {
    product := &md.Product{
        ID:          uuid.New().String(),
        Name:        name,
        Description: description,
        Price:       price,
        Stock:       stock,
    }
    if err := s.repo.AddProduct(product); err != nil {
        return nil, err
    }
    return product, nil
}

func (s *ProductService) GetProduct(id string) (*md.Product, error) {
    return s.repo.GetProduct(id)
}

func (s *ProductService) UpdateProduct(id, name, description string, price float32, stock int32) (*md.Product, error) {
    product, err := s.repo.GetProduct(id)
    if err != nil {
        return nil, err
    }
    product.Name = name
    product.Description = description
    product.Price = price
    product.Stock = stock
    if err := s.repo.UpdateProduct(product); err != nil {
        return nil, err
    }
    return product, nil
}

func (s *ProductService) DeleteProduct(id string) error {
    return s.repo.DeleteProduct(id)
}

func (s *ProductService) ListProducts() ([]*md.Product, error) {
    return s.repo.ListProducts()
}
