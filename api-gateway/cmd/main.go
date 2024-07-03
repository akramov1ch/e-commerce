package main

import (
	"api-gateway/api"
	connf "api-gateway/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := connf.LoadConfig()
	conn, err := grpc.NewClient(":" + cfg.UserServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	log.Println("User Service Connected")
	defer conn.Close()

	
	connP, err := grpc.NewClient(":" + cfg.ProductServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	log.Println("Product Service Connected")
	defer conn.Close()

	orderconn, err := grpc.NewClient(":" + cfg.OrderServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	log.Println("Order Service Connected")
	defer orderconn.Close()
	router := api.NewRouter(conn, connP, orderconn)
	router.Run(cfg.ApiGatewayPort)
}
