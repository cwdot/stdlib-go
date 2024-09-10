package collections

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func NewCounterMap[K comparable]() *CounterMap[K] {
	counter := make(map[K]int)
	return &CounterMap[K]{
		counter: counter,
	}
}

type CounterMap[K comparable] struct {
	mu      sync.Mutex
	counter map[K]int
}

// Add saves item; creates path
// Returns the new value and a boolean indicating if the key was found
func (km *CounterMap[K]) Add(key K, increment int) (int, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	var found bool
	value, ok := km.counter[key]
	if ok {
		value += increment
		found = true
	} else {
		value = increment
	}
	km.counter[key] = value
	return value, found
}

// Remove saves item; creates path
// Returns the new value and a boolean indicating if the key was found
func (km *CounterMap[K]) Remove(key K, increment int) (int, bool) {
	return km.Add(key, -increment)
}

func (km *CounterMap[K]) AddAll(other *CounterMap[K]) {
	km.mu.Lock()
	defer km.mu.Unlock()

	other.mu.Lock()
	defer other.mu.Unlock()

	for k, v := range other.counter {
		km.counter[k] += v
	}
}

// Get return Item by its id
func (km *CounterMap[K]) Get(key K) (int, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	item, ok := km.counter[key]
	return item, ok
}

// Clone return a copy of the map
func (km *CounterMap[K]) Clone() map[K]int {
	km.mu.Lock()
	defer km.mu.Unlock()

	clone := make(map[K]int)
	for k, v := range km.counter {
		clone[k] = v
	}
	return clone
}

// Keys return all keys
func (km *CounterMap[K]) Keys() []K {
	km.mu.Lock()
	defer km.mu.Unlock()

	keys := make([]K, 0, len(km.counter))
	for k := range km.counter {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j])
	})
	return keys
}

// String debug
func (km *CounterMap[K]) String() string {
	keys := km.Keys()

	km.mu.Lock()
	defer km.mu.Unlock()

	items := make([]string, 0, len(km.counter))
	for _, k := range keys {
		v := km.counter[k]
		items = append(items, fmt.Sprintf("%v=%d", k, v))
	}

	return fmt.Sprintf("CounterMap{%s}", strings.Join(items, ", "))
}
