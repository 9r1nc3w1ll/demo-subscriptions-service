package models

import "time"

type SubscriptionPlan struct {
	ID          uint
	ProductID   uint
	Duration    uint // (days)
	Price       float64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Product     Product
}
