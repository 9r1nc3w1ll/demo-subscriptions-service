package db

import (
	"lithium-test/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=root dbname=lithium_test_db port=5490 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.Product{})
	db.AutoMigrate(models.SubscriptionPlan{})

	return db, nil
}
