package main

import (
    "database/sql"
    "fmt"
    "log"
    "net"
    "product/config"
    hand "product/handler"
    "product/repository"
    "product/service"
    prodpb "product/proto/productproto"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
    "google.golang.org/grpc"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    cfg := config.LoadConfig()
    db, err := sql.Open("postgres", "host="+cfg.DBHost+" port="+cfg.DBPort+" user="+cfg.DBUser+" dbname="+cfg.DBName+" password="+cfg.DBPassword+" sslmode=disable")
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer db.Close()

    repo := repository.NewPostgresRepository(db)
    service := service.NewProductService(repo)
    server := hand.NewServer(service)

    fmt.Println("Server is running on port 8003")
    lis, err := net.Listen("tcp", ":8003")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    prodpb.RegisterProductServiceServer(s, server)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
