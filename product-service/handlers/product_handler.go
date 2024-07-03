package handlers

import (
	"context"
	"log"

	"product-service/models"
	"product-service/proto"
	"product-service/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	repo repository.ProductRepository
	proto.UnimplementedProductServiceServer
}

func NewProductService() *ProductService {
	repo, err := repository.NewPostgresRepository()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	product := &models.Product{
		ID:          uuid.New().String(),
		ProductName: req.GetProductName(),
		Description: req.GetDescription(),
	}

	err := s.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &proto.CreateProductResponse{
		Id:          product.ID,
		ProductName: product.ProductName,
		Description: product.Description,
	}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.UpdateProductResponse, error) {
	product := &models.Product{
		ID:          req.GetId(),
		ProductName: req.GetProductName(),
		Description: req.GetDescription(),
	}

	err := s.repo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateProductResponse{
		Id:          product.ID,
		ProductName: product.ProductName,
		Description: product.Description,
	}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.DeleteProductResponse, error) {
	err := s.repo.DeleteProduct(req.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.DeleteProductResponse{
		Status: "Deleted",
	}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	product, err := s.repo.GetProductByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.GetProductResponse{
		Id:          product.ID,
		ProductName: product.ProductName,
		Description: product.Description,
	}, nil
}

func (s *ProductService) GetProducts(req *proto.GetProductsRequest, stream proto.ProductService_GetProductsServer) error {
	products, err := s.repo.GetProducts()
	if err != nil {
		return err
	}

	for _, product := range products {
		resp := &proto.GetProductResponse{
			Id:          product.ID,
			ProductName: product.ProductName,
			Description: product.Description,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}

	return nil
}
