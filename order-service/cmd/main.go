package main

import (
    "database/sql"
    "fmt"
    "log"
    "net"
    "order-service/config"
    "order-service/order"
    orderpb "order-service/order/proto/orderproto"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    "google.golang.org/grpc"
)

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    cfg := config.LoadConfig()
    db, err := sql.Open("postgres", "host=" + cfg.DBHost + " port=" + cfg.DBPort + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " sslmode=disable")
    if err != nil {
        log.Fatalf("failed to connect to the database: %v", err)
    }
    defer db.Close()

    repo := order.NewPostgresRepository(db)
    service := order.NewOrderService(repo)

    productClient, err := order.NewProductClient(":" + cfg.USER_SERVICE_PORT) 
    if err != nil {
        log.Fatalf("failed to create product client: %v", err)
    }

    service.SetProductClient(productClient) 

    server := order.NewServer(service)

    fmt.Println("Server is running on port :", cfg.PORT)
    lis, err := net.Listen("tcp", ":" + cfg.PORT)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    orderpb.RegisterOrderServiceServer(s, server)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
