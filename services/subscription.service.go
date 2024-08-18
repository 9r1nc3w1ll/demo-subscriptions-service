package services

import (
	"context"
	"lithium-test/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SubscriptionService struct {
	pb.UnimplementedSubscriptionServiceServer
}

func (s *SubscriptionService) GetSubscriptionPlan(ctx context.Context, in *pb.GetSubscriptionPlanInput) (*pb.SubscriptionPlan, error) {
	// panic("no implemented")

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

func NewSubscriptionService() *SubscriptionService {
	service := SubscriptionService{}

	return &service
}
