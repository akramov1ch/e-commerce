package handlers

import (
	"context"
	"net/http"
	"io"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"api-gateway/proto"
)

func OrderRoutes(router *gin.RouterGroup, orderConn *grpc.ClientConn) {
	client := proto.NewOrderServiceClient(orderConn)

	router.POST("/createorder", func(c *gin.Context) {
		var req proto.CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.CreateOrder(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.POST("/createorders", func(c *gin.Context) {
		stream, err := client.CreateOrders(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var reqs []*proto.CreateOrderRequest
		if err := c.ShouldBindJSON(&reqs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		responses := []*proto.CreateOrderResponse{}
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			responses = append(responses, resp)
		}
		c.JSON(http.StatusOK, gin.H{"orders": responses})
	})
}
