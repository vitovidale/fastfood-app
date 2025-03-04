package domain

import (
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct_NewProductWithID(t *testing.T) {
	id := NewID()
	p := NewProductWithID(id, "product", "description", 10.5, NewID())
	require.Equal(t, id, p.ID)
	require.Equal(t, "product", p.Name)
	require.Equal(t, "description", p.Description)
	require.Equal(t, 10.5, p.Price)
	require.False(t, p.CreatedAt.IsZero())
	require.False(t, p.UpdatedAt.IsZero())
	require.True(t, p.DeletedAt.IsZero())
}

func TestProduct_NewProduct(t *testing.T) {
	p := NewProduct("product", "description", 10.5, NewID())
	require.NotEmpty(t, p.ID)
	require.Equal(t, "product", p.Name)
	require.Equal(t, "description", p.Description)
	require.Equal(t, 10.5, p.Price)
	require.False(t, p.CreatedAt.IsZero())
	require.False(t, p.UpdatedAt.IsZero())
	require.True(t, p.DeletedAt.IsZero())
}

func TestProduct_Activate(t *testing.T) {
	now := time.Now()
	t.Run("product is inactive", func(t *testing.T) {
		p := Product{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
			DeletedAt: &now,
		}
		err := p.Activate()
		require.NoError(t, err)
		require.True(t, p.IsActive())
	})
	t.Run("product is active", func(t *testing.T) {
		p := Product{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
		}
		err := p.Activate()
		require.Error(t, err)
		require.EqualError(t, err, "product already active")
	})
}

func TestProduct_Inactivate(t *testing.T) {
	now := time.Now()

	t.Run("product is active", func(t *testing.T) {
		p := Product{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
		}
		err := p.Inactivate()
		require.NoError(t, err)
		require.False(t, p.IsActive())
	})
	t.Run("product is inactive", func(t *testing.T) {
		p := Product{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
			DeletedAt: &now,
		}
		err := p.Inactivate()
		require.Error(t, err)
		require.EqualError(t, err, "product already inactive")
	})
}
