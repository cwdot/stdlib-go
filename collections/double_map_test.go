package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoubleMap_AddAndGet(t *testing.T) {
	dm := NewDoubleMap[string, string, int]()
	dm.Add("parent1", "key1", 100)

	value, parentExists, keyExists := dm.Get("parent1", "key1")
	require.True(t, parentExists)
	require.True(t, keyExists)
	require.Equal(t, 100, value)
}

func TestDoubleMap_AddOverwritesExistingValue(t *testing.T) {
	dm := NewDoubleMap[string, string, int]()
	dm.Add("parent1", "key1", 100)
	dm.Add("parent1", "key1", 200)

	value, parentExists, keyExists := dm.Get("parent1", "key1")
	require.True(t, parentExists)
	require.True(t, keyExists)
	require.Equal(t, 200, value)
}

func TestDoubleMap_GetNonExistentParent(t *testing.T) {
	dm := NewDoubleMap[string, string, int]()

	_, parentExists, _ := dm.Get("nonexistent", "key1")
	require.False(t, parentExists)
}

func TestDoubleMap_GetNonExistentKey(t *testing.T) {
	dm := NewDoubleMap[string, string, int]()
	dm.Add("parent1", "key1", 100)

	_, _, keyExists := dm.Get("parent1", "nonexistent")
	require.False(t, keyExists)
}

func TestDoubleMap_String(t *testing.T) {
	dm := NewDoubleMap[string, string, int]()
	dm.Add("parent1", "key1", 100)
	dm.Add("parent2", "key2", 200)

	expected := "DoubleMap{parent1=map[key1:100], parent2=map[key2:200]}"
	require.Equal(t, expected, dm.String())
}
