package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"user/config"
	"user/user"
	"user/user/proto/uproto"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading.env file")
	}

	cfg := config.LoadConfig()
	db, err := sql.Open("postgres", "host="+cfg.DBHost+" port="+cfg.DBPort+" user="+cfg.DBUser+" password="+cfg.DBPassword+" dbname="+cfg.DBName+" sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	repo := user.NewPostgresRepository(db)
	service := user.NewUserService(repo)
	server := user.NewServer(service)

	fmt.Println("Server is running on port :8002")
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	uproto.RegisterUserServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
