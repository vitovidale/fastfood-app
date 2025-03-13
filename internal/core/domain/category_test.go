package domain

import (
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCategory_NewCategoryWithID(t *testing.T) {
	id := NewID()
	c := NewCategoryWithID(id, "category")
	require.Equal(t, id, c.ID)
	require.Equal(t, "category", c.Name)
	require.False(t, c.CreatedAt.IsZero())
	require.False(t, c.UpdatedAt.IsZero())
	require.True(t, c.DeletedAt.IsZero())
}

func TestCategory_NewCategory(t *testing.T) {
	c := NewCategory("category")
	require.NotEmpty(t, c.ID)
	require.Equal(t, "category", c.Name)
	require.False(t, c.CreatedAt.IsZero())
	require.False(t, c.UpdatedAt.IsZero())
	require.True(t, c.DeletedAt.IsZero())
}

func TestCategory_Activate(t *testing.T) {
	now := time.Now()
	t.Run("category is inactive", func(t *testing.T) {
		c := Category{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
			DeletedAt: &now,
		}
		err := c.Activate()
		require.NoError(t, err)
		require.True(t, c.IsActive())
	})
	t.Run("category is active", func(t *testing.T) {
		c := Category{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
		}
		err := c.Activate()
		require.Error(t, err)
		require.EqualError(t, err, "category already active")
	})
}

func TestCategory_Inactivate(t *testing.T) {
	now := time.Now()
	t.Run("category is active", func(t *testing.T) {
		c := Category{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
		}
		err := c.Inactivate()
		require.NoError(t, err)
		require.False(t, c.IsActive())
	})
	t.Run("category is inactive", func(t *testing.T) {
		c := Category{
			CreatedAt: time.Now(),
			UpdatedAt: &now,
			DeletedAt: &now,
		}
		err := c.Inactivate()
		require.Error(t, err)
		require.EqualError(t, err, "category already inactive")
	})
}
