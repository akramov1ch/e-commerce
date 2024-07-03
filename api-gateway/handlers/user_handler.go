package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"api-gateway/proto"
	"api-gateway/utils"
)

func UserRoutes(router *gin.RouterGroup, userConn *grpc.ClientConn, jwtSecret string) {
	client := proto.NewUserServiceClient(userConn)

	router.POST("/user/create", func(c *gin.Context) {
		var req proto.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error3": err.Error()})
			return
		}

		resp, err := client.Register(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error4": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.POST("/login", func(c *gin.Context) {
		var req proto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.Login(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, err := utils.GenerateToken(resp.GetToken(), jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	router.PUT("/updateuser", func(c *gin.Context) {
		var req proto.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.UpdateUser(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.DELETE("/deleteuser/:id", func(c *gin.Context) {
		req := proto.DeleteUserRequest{
			Id: c.Param("id"),
		}

		resp, err := client.DeleteUser(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.GET("/getuser/:id", func(c *gin.Context) {
		req := proto.GetUserRequest{
			Id: c.Param("id"),
		}

		resp, err := client.GetUser(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})
}
