package handler

import (
	"context"
	"net/http"

	order "api-gateway/proto/order"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type OrderHandler struct {
	client order.OrderServiceClient
}

func NewOrderHandler(conn *grpc.ClientConn) *OrderHandler {
	client := order.NewOrderServiceClient(conn)
	return &OrderHandler{client: client}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req order.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.CreateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	req := order.GetOrderRequest{Id: id}
	resp, err := h.client.GetOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	req := order.DeleteOrderRequest{Id: id}
	resp, err := h.client.DeleteOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	req := order.ListOrdersRequest{}
	resp, err := h.client.ListOrders(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *OrderHandler) CreateOrders(c *gin.Context) {
	stream, err := h.client.CreateOrders(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req order.CreateOrderRequest
	for {
		if err := c.ShouldBindJSON(&req); err != nil {
			break
		}
		if err := stream.Send(&req); err != nil {
			break
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
