package handlers

import (
	"context"
	"errors"
	upb "user/proto/uproto"
	"user/service" 
	rp "user/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	upb.UnimplementedUserServiceServer
	service service.Service
}

func NewServer(service service.Service) *Server {
	return &Server{service: service}
}

func (s *Server) CreateUser(ctx context.Context, req *upb.CreateUserRequest) (*upb.CreateUserResponse, error) {
	user, err := s.service.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &upb.CreateUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req *upb.GetUserRequest) (*upb.GetUserResponse, error) {
	user, err := s.service.GetUser(req.Id)
	if err != nil {
		if errors.Is(err, rp.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, err
	}
	return &upb.GetUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *upb.UpdateUserRequest) (*upb.UpdateUserResponse, error) {
	user, err := s.service.UpdateUser(req.Id, req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return &upb.UpdateUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *upb.DeleteUserRequest) (*upb.DeleteUserResponse, error) {
	if err := s.service.DeleteUser(req.Id); err != nil {
		if errors.Is(err, rp.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, err
	}
	return &upb.DeleteUserResponse{Message: "User deleted"}, nil
}
