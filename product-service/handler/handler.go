package handler

import (
	"context"
	"errors"
	prodpb "product/proto/productproto"

	sv "product/service"

	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUserNotFound = errors.New("user not found")

type Server struct {
	prodpb.UnimplementedProductServiceServer
	service sv.Service
}

func NewServer(service sv.Service) *Server {
	return &Server{service: service}
}

func (s *Server) AddProduct(ctx context.Context, req *prodpb.AddProductRequest) (*prodpb.AddProductResponse, error) {
	product, err := s.service.AddProduct(req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		return nil, err
	}
	return &prodpb.AddProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (s *Server) GetProduct(ctx context.Context, req *prodpb.GetProductRequest) (*prodpb.GetProductResponse, error) {
	product, err := s.service.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}
	return &prodpb.GetProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (s *Server) UpdateProduct(ctx context.Context, req *prodpb.UpdateProductRequest) (*prodpb.UpdateProductResponse, error) {
	product, err := s.service.UpdateProduct(req.Id, req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		return nil, err
	}
	return &prodpb.UpdateProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (s *Server) DeleteProduct(ctx context.Context, req *prodpb.DeleteProductRequest) (*prodpb.DeleteProductResponse, error) {
	if err := s.service.DeleteProduct(req.Id); err != nil {
		if err == ErrUserNotFound {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, err
	}
	return &prodpb.DeleteProductResponse{Message: "product deleted"}, nil
}

func (s *Server) ListProducts(req *prodpb.ListProductsRequest, stream prodpb.ProductService_ListProductsServer) error {
	products, err := s.service.ListProducts()
	if err != nil {
		return err
	}

	for _, product := range products {
		if err := stream.Send(&prodpb.ListProductsResponse{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}); err != nil {
			return err
		}
	}
	return nil
}
