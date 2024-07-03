package user

import (
	"context"
	"user/user/proto/uproto"

	"google.golang.org/grpc/status"
)

type Server struct {
	uproto.UnimplementedUserServiceServer
	service Service
}

func NewServer(service Service) *Server {
	return &Server{service: service}
}

func (s *Server) CreateUser(ctx context.Context, req *uproto.CreateUserRequest) (*uproto.CreateUserResponse, error) {
	user, err := s.service.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &uproto.CreateUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req *uproto.GetUserRequest) (*uproto.GetUserResponse, error) {
	user, err := s.service.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &uproto.GetUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *uproto.UpdateUserRequest) (*uproto.UpdateUserResponse, error) {
	user, err := s.service.UpdateUser(req.Id, req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return &uproto.UpdateUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *uproto.DeleteUserRequest) (*uproto.DeleteUserResponse, error) {
	if err := s.service.DeleteUser(req.Id); err != nil {
		if err == ErrUserNotFound {
			return nil, status.Errorf(status.Code(err), "user not found")
		}
		return nil, err
	}
	return &uproto.DeleteUserResponse{Message: "User o'chirildi"}, nil
}
