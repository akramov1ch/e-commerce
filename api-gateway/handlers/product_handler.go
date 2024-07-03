package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"api-gateway/proto"
)

func ProductRoutes(router *gin.RouterGroup, productConn *grpc.ClientConn) {
	client := proto.NewProductServiceClient(productConn)

	router.POST("/createproduct", func(c *gin.Context) {
		var req proto.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.CreateProduct(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.PUT("/updateproduct", func(c *gin.Context) {
		var req proto.UpdateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.UpdateProduct(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.DELETE("/deleteproduct/:id", func(c *gin.Context) {
		req := proto.DeleteProductRequest{
			Id: c.Param("id"),
		}

		resp, err := client.DeleteProduct(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.GET("/getproduct/:id", func(c *gin.Context) {
		req := proto.GetProductRequest{
			Id: c.Param("id"),
		}

		resp, err := client.GetProduct(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.GET("/getproducts", func(c *gin.Context) {
		req := proto.GetProductsRequest{}

		stream, err := client.GetProducts(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var products []*proto.GetProductResponse
		for {
			product, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			products = append(products, product.GetProducts()...)
		}
		c.JSON(http.StatusOK, gin.H{"products": products})
	})
}
