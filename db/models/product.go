package models

import (
	"lithium-test/pb"
	"time"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Type        string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Digital product fields
	FileSize     *int
	DownloadLink *string

	// Phhysical product fields
	Weight     *float64
	Dimensions *string

	// Subscription product fields
	SubscriptionPeriod *string
	RenewalPrice       *float64
	NextRenewalDate    *time.Time
}

func (p *Product) ToProto() *pb.Product {
	if p != nil {
		return &pb.Product{
			Id:        p.ID,
			Name:      p.Name,
			Type:      p.Type,
			Price:     float32(p.Price), // TODO: Sort this to avoid truncations
			CreatedAt: p.UpdatedAt.GoString(),
			UpdatedAt: p.UpdatedAt.GoString(),
		}
	}

	return nil
}
