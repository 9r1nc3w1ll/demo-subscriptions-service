package models

import (
	"lithium-test/pb"
	"time"
)

type SubscriptionPlan struct {
	ID          int64
	ProductID   int64
	Duration    int64 // (days)
	Price       float64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Product     Product
}

func (sp *SubscriptionPlan) ToProto() *pb.SubscriptionPlan {
	if sp != nil {
		return &pb.SubscriptionPlan{
			Id:        sp.ID,
			Name:      sp.Name,
			ProductId: sp.ProductID,
			Duration:  int64(sp.Duration),
			Price:     float32(sp.Price), // TODO: Sort this to avoid truncations
			CreatedAt: sp.UpdatedAt.GoString(),
			UpdatedAt: sp.UpdatedAt.GoString(),
		}
	}

	return nil
}
