package main

import (
    "database/sql"
    "fmt"
    "log"
    "net"
    "user/config"
    "user/handlers"
    "user/repository"
    "user/service"
    upb "user/proto/uproto"

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

    repo := repository.NewPostgresRepository(db)
    svc := service.NewUserService(repo)
    server := handlers.NewServer(svc)

    fmt.Println("Server is running on port :", cfg.PORT)
    lis, err := net.Listen("tcp", ":" + cfg.PORT)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    upb.RegisterUserServiceServer(s, server)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
