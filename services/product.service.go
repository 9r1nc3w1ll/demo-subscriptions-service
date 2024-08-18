package services

import (
	"context"
	"lithium-test/pb"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	db *gorm.DB
}

func (s *ProductService) GetProduct(ctx context.Context, in *pb.GetProductInput) (*pb.Product, error) {
	// panic("no implemented")

	return nil, nil
}

func (s *ProductService) ListProducts(ctx context.Context, in *emptypb.Empty) (*pb.ProductList, error) {
	// panic("no implemented")

	return nil, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, in *pb.CreateProductInput) (*pb.Product, error) {
	// panic("no implemented")

	return nil, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {
	// panic("no implemented")

	return nil, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, in *pb.DeleteProductInput) (*emptypb.Empty, error) {
	// panic("no implemented")

	return nil, nil
}

func NewProductService(db *gorm.DB) *ProductService {
	service := ProductService{
		db: db,
	}

	return &service
}
