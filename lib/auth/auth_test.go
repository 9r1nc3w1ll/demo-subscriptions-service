package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	t.Run("Should fail on invalid token", func(t *testing.T) {
		err := ValidateToken(context.Background(), "Bearer")
		assert.Error(t, err)
	})

	t.Run("Should pass with a valid token", func(t *testing.T) {
		err := ValidateToken(SetAuthContextForTest(context.Background(), "Bearer", "VALID_TEST_TOKEN"), "Bearer")
		assert.Nil(t, err)
	})
}
