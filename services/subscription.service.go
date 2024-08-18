package services

import (
	"context"
	"lithium-test/auth"
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

	return nil, nil
}

func (s *SubscriptionService) ListSubscriptionPlans(ctx context.Context, in *emptypb.Empty) (*pb.SubscriptionPlanList, error) {
	// panic("no implemented")

	return nil, nil
}

func (s *SubscriptionService) CreateSubscriptionPlan(ctx context.Context, in *pb.CreateSubscriptionPlanInput) (*pb.SubscriptionPlan, error) {
	// panic("no implemented")

	return nil, nil
}

func (s *SubscriptionService) UpdateSubscriptionPlan(ctx context.Context, in *pb.SubscriptionPlan) (*pb.SubscriptionPlan, error) {
	// panic("no implemented")

	return nil, nil
}

func (s *SubscriptionService) DeleteSubscriptionPlan(ctx context.Context, in *pb.DeleteSubscriptionPlanInput) (*emptypb.Empty, error) {
	// panic("no implemented")

	return nil, nil
}

func NewSubscriptionService(db *gorm.DB) *SubscriptionService {
	service := SubscriptionService{
		db: db,
	}

	return &service
}
