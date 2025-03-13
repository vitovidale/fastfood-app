package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestID_NewID(t *testing.T) {
	t.Run("new id", func(t *testing.T) {
		id := NewID()
		require.NotEmpty(t, id)
	})
}

func TestID_ParseID(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		id := NewID()
		s := id.String()
		parsedID, err := ParseID(s)
		require.NoError(t, err)
		require.Equal(t, id, parsedID)
	})
	t.Run("invalid id", func(t *testing.T) {
		_, err := ParseID("invalid")
		require.Error(t, err)
	})
}
