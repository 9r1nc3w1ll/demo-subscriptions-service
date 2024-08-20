package models

import (
	"lithium-test/pb"
	"time"
)

var PhysicalProductType = "physical"
var DigitalProductType = "digital"
var SubscriptionProductType = "subscription"

type ProductType string

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float32
	Type        ProductType
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Digital product fields
	FileSize     *int64
	DownloadLink *string

	// Phhysical product fields
	Weight     *float32
	Dimensions *string

	// Subscription product fields
	SubscriptionPeriod *string
	RenewalPrice       *float32
	NextRenewalDate    *time.Time
}

func (pt ProductType) IsValid() bool {
	allowedTypes := []string{
		PhysicalProductType,
		DigitalProductType,
		SubscriptionProductType,
	}

	for _, typeName := range allowedTypes {
		if string(pt) == typeName {
			return true
		}
	}

	return false
}

func (p *Product) ToProto() *pb.Product {
	if p != nil {
		pbProduct := pb.Product{
			Id:        p.ID,
			Name:      p.Name,
			Type:      string(p.Type),
			Price:     float32(p.Price), // TODO: Sort this to avoid truncations
			CreatedAt: p.UpdatedAt.GoString(),
			UpdatedAt: p.UpdatedAt.GoString(),
		}

		if p.Dimensions != nil {
			pbProduct.Dimensions = *p.Dimensions
		}

		if p.Weight != nil {
			pbProduct.Weight = *p.Weight
		}

		if p.FileSize != nil {
			pbProduct.FileSize = *p.FileSize
		}

		if p.DownloadLink != nil {
			pbProduct.DownloadLink = *p.DownloadLink
		}

		if p.SubscriptionPeriod != nil {
			pbProduct.SubscriptionPeriod = *p.SubscriptionPeriod
		}

		if p.RenewalPrice != nil {
			pbProduct.RenewalPrice = *p.RenewalPrice
		}

		return &pbProduct
	}

	return nil
}
