package services

import (
	"context"
	"lithium-test/auth"
	"lithium-test/db/models"
	"lithium-test/pb"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type SubscriptionService struct {
	pb.UnimplementedSubscriptionServiceServer
	db *gorm.DB
}

func (s *SubscriptionService) GetSubscriptionPlan(ctx context.Context, in *pb.GetSubscriptionPlanInput) (*pb.SubscriptionPlan, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	plan := models.SubscriptionPlan{}
	if result := s.db.Where(&models.SubscriptionPlan{ID: in.Id}).First(&plan); result.Error != nil {
		return nil, result.Error
	}

	return plan.ToProto(), nil
}

func (s *SubscriptionService) ListSubscriptionPlans(ctx context.Context, in *emptypb.Empty) (*pb.SubscriptionPlanList, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	plans := []models.SubscriptionPlan{}

	if result := s.db.Find(&plans); result.Error != nil {
		return nil, result.Error
	}

	data := []*pb.SubscriptionPlan{}
	for _, plan := range plans {
		data = append(data, plan.ToProto())
	}

	return &pb.SubscriptionPlanList{
		Data: data,
	}, nil
}

func (s *SubscriptionService) CreateSubscriptionPlan(ctx context.Context, in *pb.CreateSubscriptionPlanInput) (*pb.SubscriptionPlan, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	plan := models.SubscriptionPlan{
		Name:        in.Name,
		Description: in.Description,
		ProductID:   in.ProductId,
		Duration:    in.Duration,
		Price:       float64(in.Price),
	}

	if result := s.db.Create(&plan); result.Error != nil {
		return nil, result.Error
	}

	return plan.ToProto(), nil
}

func (s *SubscriptionService) UpdateSubscriptionPlan(ctx context.Context, in *pb.SubscriptionPlan) (*pb.SubscriptionPlan, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	plan := models.SubscriptionPlan{}
	if result := s.db.Where(&models.SubscriptionPlan{ID: in.Id}).First(&plan); result.Error != nil {
		return nil, result.Error
	}

	plan.Name = in.Name
	plan.Price = float64(in.Price)
	plan.Description = in.Description
	plan.ProductID = in.ProductId

	if result := s.db.Save(&plan); result.Error != nil {
		return nil, result.Error
	}

	return plan.ToProto(), nil
}

func (s *SubscriptionService) DeleteSubscriptionPlan(ctx context.Context, in *pb.DeleteSubscriptionPlanInput) (*emptypb.Empty, error) {
	if err := auth.ValidateToken(ctx, "Bearer"); err != nil {
		return nil, err
	}

	plan := models.SubscriptionPlan{}
	if result := s.db.Where(&models.SubscriptionPlan{ID: in.Id}).First(&plan); result.Error != nil {
		return nil, result.Error
	}

	if result := s.db.Delete(&plan); result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}

func NewSubscriptionService(db *gorm.DB) *SubscriptionService {
	service := SubscriptionService{
		db: db,
	}

	return &service
}
