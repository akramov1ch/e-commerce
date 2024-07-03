package handler

import (
	prod "api-gateway/proto/product"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ProductHandler struct {
	client prod.ProductServiceClient
}

func NewProductHandler(conn *grpc.ClientConn) *ProductHandler {
	client := prod.NewProductServiceClient(conn)
	return &ProductHandler{client: client}
}

func (h *ProductHandler) AddProduct(c *gin.Context) {
	var req prod.AddProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.AddProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	req := prod.GetProductRequest{Id: id}
	resp, err := h.client.GetProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var req prod.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.UpdateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	req := prod.DeleteProductRequest{Id: id}
	resp, err := h.client.DeleteProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	req := prod.ListProductsRequest{}
	stream, err := h.client.ListProducts(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var products []prod.ListProductsResponse
	for {
		product, err := stream.Recv()
		if err != nil {
			break
		}
		products = append(products, *product)
	}
	c.JSON(http.StatusOK, products)
}
