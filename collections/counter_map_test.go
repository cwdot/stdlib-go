package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCounterMap(t *testing.T) {
	cm := NewCounterMap[string]()
	assert.NotNil(t, cm)
	assert.Equal(t, 0, len(cm.counter))
}

func TestCounterMap_Add(t *testing.T) {
	cm := NewCounterMap[string]()
	cm.Add("key1", 1)
	cm.Add("key1", 1)
	cm.Add("key2", 4)

	assert.Equal(t, 2, cm.counter["key1"])
	assert.Equal(t, 4, cm.counter["key2"])
}

func TestCounterMap_Get(t *testing.T) {
	cm := NewCounterMap[string]()
	t.Run("ExistentKey", func(t *testing.T) {
		cm.Add("key1", 1)

		value, ok := cm.Get("key1")
		assert.True(t, ok)
		assert.Equal(t, 1, value)
	})
	t.Run("NonExistentKey", func(t *testing.T) {
		value, ok := cm.Get("key2")
		assert.False(t, ok)
		assert.Equal(t, 0, value)
	})
}

func TestCounterMap_AddAndGet(t *testing.T) {
	cm := NewCounterMap[string]()
	cm.Add("key1", 1)
	value, ok := cm.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, 1, value)
}

func TestCounterMap_AddMultipleTimes(t *testing.T) {
	cm := NewCounterMap[string]()
	cm.Add("key1", 1)
	cm.Add("key1", 2)
	value, ok := cm.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestCounterMap_Clone(t *testing.T) {
	cm := NewCounterMap[string]()
	cm.Add("key1", 1)
	cm.Add("key2", 2)
	clone := cm.Clone()
	assert.Equal(t, 2, len(clone))
	assert.Equal(t, 1, clone["key1"])
	assert.Equal(t, 2, clone["key2"])
}

func TestCounterMap_String(t *testing.T) {
	cm := NewCounterMap[string]()
	cm.Add("key1", 1)
	cm.Add("key2", 2)
	expected := "CounterMap{key1=1, key2=2}"
	assert.Equal(t, expected, cm.String())
}

func TestCounterMap_Keys(t *testing.T) {
	cm := NewCounterMap[string]()
	cm.Add("key1", 1)
	cm.Add("key2", 2)
	keys := cm.Keys()
	assert.Equal(t, 2, len(keys))
	assert.Contains(t, keys, "key1")
	assert.Contains(t, keys, "key2")
}

func TestCounterMap_AddAll(t *testing.T) {
	cm := NewCounterMap[string]()
	cm.Add("key1", 1)
	cm.Add("key2", 4)

	other := NewCounterMap[string]()
	other.Add("key1", 2)
	other.Add("key3", 5)

	cm.AddAll(other)
	assert.Equal(t, 3, cm.counter["key1"])
	assert.Equal(t, 4, cm.counter["key2"])
	assert.Equal(t, 5, cm.counter["key3"])
}
