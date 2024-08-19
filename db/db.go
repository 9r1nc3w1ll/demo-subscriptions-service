package db

import (
	"lithium-test/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.Product{})
	db.AutoMigrate(models.SubscriptionPlan{})

	return db, nil
}
