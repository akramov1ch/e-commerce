package service

import (
	"context"
	"errors"
	"order-service/models"
	prodpb "order-service/proto/productproto"
	"order-service/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service interface {
	CreateOrder(userID, productID string, quantity int32) (*models.Order, error)
	GetOrder(id string) (*models.Order, error)
	DeleteOrder(id string) error
	ListOrders() ([]*models.Order, error)
	SetProductClient(client prodpb.ProductServiceClient)
}

type OrderService struct {
	repo          repository.Repository
	productClient prodpb.ProductServiceClient
}

func NewOrderService(repo repository.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) SetProductClient(client prodpb.ProductServiceClient) {
	s.productClient = client
}

func (s *OrderService) CreateOrder(userID, productID string, quantity int32) (*models.Order, error) {
	productReq := &prodpb.GetProductRequest{Id: productID}
	productRes, err := s.productClient.GetProduct(context.Background(), productReq)
	if err != nil {
		return nil, err
	}
	if productRes.Stock < quantity {
		return nil, errors.New("insufficient stock")
	}

	totalPrice := float64(productRes.Price) * float64(quantity)

	order := &models.Order{
		ID:         generateID(), // generateID is a placeholder function to generate unique IDs
		UserID:     userID,
		ProductID:  productID,
		Quantity:   quantity,
		Status:     "Created",
		TotalPrice: totalPrice,
	}

	if err := s.repo.CreateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrder(id string) (*models.Order, error) {
	return s.repo.GetOrder(id)
}

func (s *OrderService) DeleteOrder(id string) error {
	return s.repo.DeleteOrder(id)
}

func (s *OrderService) ListOrders() ([]*models.Order, error) {
	return s.repo.ListOrders()
}

func generateID() string {
	return "unique-id"
}

func NewProductClient(address string) (prodpb.ProductServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return prodpb.NewProductServiceClient(conn), nil
}
