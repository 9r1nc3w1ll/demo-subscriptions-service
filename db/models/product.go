package models

import "time"

type Product struct {
	ID          uint
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
