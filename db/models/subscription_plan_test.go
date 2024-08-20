package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscriptionPlan(t *testing.T) {
	t.Run("Validate that SubscriptionPlan.ToProto works", func(t *testing.T) {

		subscriptionPlan := SubscriptionPlan{
			ID:          12345,
			Name:        "Test product subscription plan",
			Description: "Test product subscription plan description",
			Price:       12.34,
			ProductID:   123456,
			Duration:    30,
		}

		proto := subscriptionPlan.ToProto()

		assert.Equal(t, subscriptionPlan.ID, proto.Id)
		assert.Equal(t, subscriptionPlan.Name, proto.Name)
		assert.Equal(t, subscriptionPlan.Description, proto.Description)
		assert.Equal(t, subscriptionPlan.ProductID, proto.ProductId)
		assert.Equal(t, subscriptionPlan.Duration, proto.Duration)
	})
}
