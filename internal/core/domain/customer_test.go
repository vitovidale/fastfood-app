package domain

import (
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomer_Activate(t *testing.T) {
	now := time.Now()
	t.Run("customer: inactive", func(t *testing.T) {
		c := Customer{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
			DeletedAt: &now,
		}
		err := c.Activate()
		require.NoError(t, err)
		require.True(t, c.IsActive())
	})
	t.Run("customer: active", func(t *testing.T) {
		c := Customer{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
		}
		err := c.Activate()
		require.Error(t, err)
		require.EqualError(t, err, "customer already active")
	})
}

func TestCustomer_Inactivate(t *testing.T) {
	now := time.Now()
	t.Run("customer: active", func(t *testing.T) {
		c := Customer{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
		}
		err := c.Deactivate()
		require.NoError(t, err)
		require.False(t, c.IsActive())
	})
	t.Run("customer: inactive", func(t *testing.T) {
		c := Customer{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
			DeletedAt: &now,
		}
		err := c.Deactivate()
		require.Error(t, err)
		require.EqualError(t, err, "customer already inactive")
	})
}
