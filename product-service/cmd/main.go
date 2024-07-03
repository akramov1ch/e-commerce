package main

import (
	"log"
	"net"

	"product-service/config"
	"product-service/handlers"
	"product-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.LoadConfig()

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.Port, err)
	}

	grpcServer := grpc.NewServer()
	productService := handlers.NewProductService()
	proto.RegisterProductServiceServer(grpcServer, productService)

	reflection.Register(grpcServer)
	log.Printf("gRPC server listening on %s", cfg.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
