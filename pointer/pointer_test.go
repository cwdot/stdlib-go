package pointer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRef(t *testing.T) {
	t.Run("Value", func(t *testing.T) {
		value := 42
		ptr := Ref(value)
		require.NotNil(t, ptr)
		require.Equal(t, value, *ptr)
	})
	t.Run("Nil", func(t *testing.T) {
		var value *int
		ptr := Ref(value)
		require.NotNil(t, ptr)
		require.Nil(t, *ptr)
	})
}

func TestDeref(t *testing.T) {
	t.Run("Value", func(t *testing.T) {
		value := 42
		ptr := &value
		derefValue := Deref(ptr)
		require.Equal(t, value, derefValue)
	})
	t.Run("Nil", func(t *testing.T) {
		var ptr *int
		require.Zero(t, Deref(ptr))
	})
}

func TestCast(t *testing.T) {
	t.Run("Value", func(t *testing.T) {
		value := 42
		valueP := &value
		ptr := Cast[int](valueP)
		require.NotNil(t, ptr)
		require.Equal(t, valueP, ptr)
	})
	t.Run("Nil", func(t *testing.T) {
		var value *int
		ptr := Cast[int](value)
		require.Nil(t, ptr)
	})
}
