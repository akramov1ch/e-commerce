package order

import (
    "github.com/google/uuid"
)

type Service interface {
    CreateOrder(userID, productID string, quantity int32) (*Order, error)
    GetOrder(id string) (*Order, error)
    DeleteOrder(id string) error
    ListOrders() ([]*Order, error)
    SetProductClient(*ProductClient)
}

type OrderService struct {
    repo          Repository
    productClient *ProductClient
}

func NewOrderService(repo Repository) Service {
    return &OrderService{repo: repo}
}

func (s *OrderService) SetProductClient(pc *ProductClient) {
    s.productClient = pc
}

func (s *OrderService) CreateOrder(userID, productID string, quantity int32) (*Order, error) {
    price, err := s.productClient.GetProductPrice(productID)
    if err != nil {
        return nil, err
    }

    totalPrice := 0.0
    if price > 0 {
        totalPrice = float64(price) * float64(quantity)
    } 

    
    order := &Order{
        ID:         uuid.New().String(),
        UserID:     userID,
        ProductID:  productID,
        Quantity:   quantity,
        Status:     "pending",
        TotalPrice: totalPrice,
    }

    if err := s.repo.CreateOrder(order); err != nil {
        return nil, err
    }
    return order, nil
}

func (s *OrderService) GetOrder(id string) (*Order, error) {
    return s.repo.GetOrder(id)
}


func (s *OrderService) DeleteOrder(id string) error {
    return s.repo.DeleteOrder(id)
}

func (s *OrderService) ListOrders() ([]*Order, error) {
    return s.repo.ListOrders()
}
