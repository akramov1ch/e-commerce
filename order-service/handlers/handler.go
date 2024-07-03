package handlers

import (
    "context"
    "io"
    "log"
    "order-service/service"
    orderpb "order-service/proto/orderproto"
)

type Server struct {
    orderpb.UnimplementedOrderServiceServer
    service service.Service
}

func NewServer(service service.Service) *Server {
    return &Server{service: service}
}

func (s *Server) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
    log.Printf("Received CreateOrder request: %+v", req)
    order, err := s.service.CreateOrder(req.UserId, req.ProductId, req.Quantity)
    if err != nil {
        log.Printf("Error creating order: %v", err)
        return nil, err
    }
    return &orderpb.CreateOrderResponse{
        Id:         order.ID,
        TotalPrice: float32(order.TotalPrice),
    }, nil
}

func (s *Server) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
    log.Printf("Received GetOrder request: %+v", req)
    order, err := s.service.GetOrder(req.Id)
    if err != nil {
        log.Printf("Error getting order: %v", err)
        return nil, err
    }
    return &orderpb.GetOrderResponse{
        Id:         order.ID,
        UserId:     order.UserID,
        ProductId:  order.ProductID,
        Quantity:   order.Quantity,
        Status:     order.Status,
        CreatedAt:  order.CreatedAt,
        UpdatedAt:  order.UpdatedAt,
        TotalPrice: float32(order.TotalPrice),
    }, nil
}

func (s *Server) DeleteOrder(ctx context.Context, req *orderpb.DeleteOrderRequest) (*orderpb.DeleteOrderResponse, error) {
    log.Printf("Received DeleteOrder request: %+v", req)
    if err := s.service.DeleteOrder(req.Id); err != nil {
        log.Printf("Error deleting order: %v", err)
        return nil, err
    }
    return &orderpb.DeleteOrderResponse{Message: "Order deleted"}, nil
}

func (s *Server) ListOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
    log.Printf("Received ListOrders request")
    orders, err := s.service.ListOrders()
    if err != nil {
        log.Printf("Error listing orders: %v", err)
        return nil, err
    }

    var ordersResponses []*orderpb.GetOrderResponse
    for _, order := range orders {
        ordersResponses = append(ordersResponses, &orderpb.GetOrderResponse{
            Id:         order.ID,
            UserId:     order.UserID,
            ProductId:  order.ProductID,
            Quantity:   order.Quantity,
            Status:     order.Status,
            TotalPrice: float32(order.TotalPrice),
            CreatedAt:  order.CreatedAt,
            UpdatedAt:  order.UpdatedAt,
        })
    }

    return &orderpb.ListOrdersResponse{
        Orders: ordersResponses,
    }, nil
}

func (s *Server) CreateOrders(stream orderpb.OrderService_CreateOrdersServer) error {
    var orderResponses []*orderpb.CreateOrderResponse

    for {
        req, err := stream.Recv()
        if err != nil {
            if err == io.EOF {
                return stream.SendAndClose(&orderpb.CreateOrdersResponse{Orders: orderResponses})
            }
            log.Printf("Error receiving stream: %v", err)
            return err
        }

        log.Printf("Received CreateOrders stream request: %+v", req)
        order, err := s.service.CreateOrder(req.UserId, req.ProductId, req.Quantity)
        if err != nil {
            log.Printf("Error creating order: %v", err)
            return err
        }

        orderResponses = append(orderResponses, &orderpb.CreateOrderResponse{
            Id:         order.ID,
            TotalPrice: float32(order.TotalPrice),
        })
    }
}
