package main

import (
	"user-service/config"
	"user-service/handlers"
	"user-service/middleware"
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
	defer userConn.Close()

	r := gin.Default()

	r.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	api := r.Group("/api")
	{
		handlers.UserRoutes(api, userConn, cfg.JWTSecret)
	}

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
