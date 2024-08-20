package models

import (
	"lithium-test/pb"
	"time"
)

type SubscriptionPlan struct {
	ID          int64
	ProductID   int64
	Duration    int64 // (days)
	Price       float32
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Product     Product
}

func (sp *SubscriptionPlan) ToProto() *pb.SubscriptionPlan {
	if sp != nil {
		return &pb.SubscriptionPlan{
			Id:          sp.ID,
			Name:        sp.Name,
			ProductId:   sp.ProductID,
			Duration:    sp.Duration,
			Description: sp.Description,
			Price:       sp.Price,
			CreatedAt:   sp.UpdatedAt.GoString(),
			UpdatedAt:   sp.UpdatedAt.GoString(),
		}
	}

	return nil
}
