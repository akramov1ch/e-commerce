package handlers

import (
	"context"
	"database/sql"
	"io"

	"order-service/proto"
	"order-service/repository"
)

type OrderHandler struct {
	repo repository.OrderRepository
	proto.UnimplementedOrderServiceServer
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{
		repo: repository.NewOrderRepository(db),
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	order, err := h.repo.CreateOrder(req.UserId, req.ProductId)
	if err != nil {
		return nil, err
	}

	return &proto.CreateOrderResponse{
		Id:        order.ID,
		UserId:    order.UserID,
		ProductId: order.ProductID,
		OrderedAt: order.OrderedAt,
	}, nil
}

func (h *OrderHandler) CreateOrders(stream proto.OrderService_CreateOrdersServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		order, err := h.repo.CreateOrder(req.UserId, req.ProductId)
		if err != nil {
			return err
		}

		resp := &proto.CreateOrderResponse{
			Id:        order.ID,
			UserId:    order.UserID,
			ProductId: order.ProductID,
			OrderedAt: order.OrderedAt,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}
