package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductType(t *testing.T) {
	t.Run("Should fail on invalid type", func(t *testing.T) {
		productType := ProductType("invalid_type")
		assert.False(t, productType.IsValid())
	})

	t.Run("Should pass on valid type", func(t *testing.T) {
		productType := ProductType("subscription")
		assert.True(t, productType.IsValid())
	})
}

func TestProduct(t *testing.T) {
	t.Run("Validate that Product.ToProto works", func(t *testing.T) {
		downloadLink := "https://download.test"
		weight := float32(11.5)
		fileSize := int64(10)
		dimension := "2x5"
		subscriptionPeriod := "1 month"
		renewalPrice := float32(30.10)

		product := Product{
			ID:                 12345,
			Name:               "Test product",
			Description:        "Test product description",
			Price:              12.34,
			Type:               "physical",
			DownloadLink:       &downloadLink,
			Weight:             &weight,
			FileSize:           &fileSize,
			Dimensions:         &dimension,
			SubscriptionPeriod: &subscriptionPeriod,
			RenewalPrice:       &renewalPrice,
		}

		proto := product.ToProto()

		assert.Equal(t, product.ID, proto.Id)
		assert.Equal(t, product.Name, proto.Name)
		assert.Equal(t, product.Price, proto.Price)
		assert.Equal(t, string(product.Type), proto.Type)
		assert.Equal(t, *product.Dimensions, proto.Dimensions)
		assert.Equal(t, *product.Weight, proto.Weight)
		assert.Equal(t, *product.FileSize, proto.FileSize)
		assert.Equal(t, *product.DownloadLink, proto.DownloadLink)
		assert.Equal(t, *product.SubscriptionPeriod, proto.SubscriptionPeriod)
		assert.Equal(t, *product.RenewalPrice, proto.RenewalPrice)
	})
}
