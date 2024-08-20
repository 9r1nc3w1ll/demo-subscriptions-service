package services

import (
	"context"
	"fmt"
	"lithium-test/db/models"
	"lithium-test/lib/auth"
	"lithium-test/pb"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	db *gorm.DB
}

func (s *ProductService) GetProduct(ctx context.Context, in *pb.GetProductInput) (*pb.Product, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	product := models.Product{}

	if result := s.db.Where(&models.Product{ID: in.Id}).First(&product); result.Error != nil {
		return nil, result.Error
	}

	return product.ToProto(), nil
}

func (s *ProductService) ListProducts(ctx context.Context, in *emptypb.Empty) (*pb.ProductList, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	products := []models.Product{}

	if result := s.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}

	data := []*pb.Product{}
	for _, product := range products {
		data = append(data, product.ToProto())
	}

	return &pb.ProductList{
		Data: data,
	}, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, in *pb.CreateProductInput) (*pb.Product, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	if !models.ProductType(in.Type).IsValid() {
		return nil, fmt.Errorf("invalid product type %s", in.Type)
	}

	product := models.Product{
		Name:        in.Name,
		Description: in.Description,
		Type:        models.ProductType(in.Type),
		Price:       in.Price,
	}

	if in.Type == models.PhysicalProductType {
		product.Dimensions = &in.Dimensions
		product.Weight = &in.Weight
	}

	if in.Type == models.DigitalProductType {
		product.FileSize = &in.FileSize
		product.DownloadLink = &in.DownloadLink
	}

	if in.Type == models.SubscriptionProductType {
		product.SubscriptionPeriod = &in.SubscriptionPeriod
		product.RenewalPrice = &in.RenewalPrice
	}

	if result := s.db.Create(&product); result.Error != nil {
		return nil, result.Error
	}

	return product.ToProto(), nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	product := models.Product{
		Name:  in.Name,
		Type:  models.ProductType(in.Type),
		Price: in.Price,
	}

	if result := s.db.Create(&product); result.Error != nil {
		return nil, result.Error
	}

	return product.ToProto(), nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, in *pb.DeleteProductInput) (*emptypb.Empty, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	product := models.Product{}
	if result := s.db.Where(&models.Product{ID: in.Id}).First(&product); result.Error != nil {
		return nil, result.Error
	}

	if result := s.db.Delete(&product); result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}

func NewProductService(db *gorm.DB) *ProductService {
	service := ProductService{
		db: db,
	}

	return &service
}
