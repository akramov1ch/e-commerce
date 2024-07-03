package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    cfg := config.LoadConfig()

    userConn, err := grpc.NewClient(cfg.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Could not connect to User Service: %v", err)
    }
    log.Println("connected to User Service")
    defer userConn.Close()

    productConn, err := grpc.NewClient(cfg.ProductServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Could not connect to Product Service: %v", err)
    }
    defer productConn.Close()
    log.Println("Connected to Product Service")

    orderConn, err := grpc.NewClient(cfg.OrderServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Could not connect to Order Service: %v", err)
    }
    defer orderConn.Close()
    log.Println("Connected to Order Service")

    r := gin.Default()

    r.Use(middleware.AuthMiddleware(cfg.JWTSecret))

    api := r.Group("/api")
    {
        handlers.UserRoutes(api, userConn, cfg.JWTSecret)
        handlers.ProductRoutes(api, productConn)
        handlers.OrderRoutes(api, orderConn)
    }

    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
