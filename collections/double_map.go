package collections

import (
	"fmt"
	"strings"
	"sync"
)

func NewDoubleMap[P comparable, K comparable, V any]() *DoubleMap[P, K, V] {
	counter := make(map[P]map[K]V)
	return &DoubleMap[P, K, V]{
		m: counter,
	}
}

type DoubleMap[P comparable, K comparable, V any] struct {
	mu sync.Mutex
	m  map[P]map[K]V
}

// Add saves item; creates path
func (km *DoubleMap[P, K, V]) Add(parent P, key K, value V) bool {
	km.mu.Lock()
	defer km.mu.Unlock()

	if inner, ok := km.m[parent]; ok {
		inner[key] = value
	} else {
		km.m[parent] = make(map[K]V)
		km.m[parent][key] = value
	}
	return true
}

// Get return Item by its id
func (km *DoubleMap[P, K, V]) Get(parent P, key K) (V, bool, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	if outer, ok := km.m[parent]; ok {
		inner, ok := outer[key]
		return inner, true, ok
	}
	var zero V
	return zero, false, false
}

// String debug
func (km *DoubleMap[P, K, V]) String() string {
	km.mu.Lock()
	defer km.mu.Unlock()

	items := make([]string, 0, len(km.m))
	for k, v := range km.m {
		items = append(items, fmt.Sprintf("%v=%v", k, v))
	}

	return fmt.Sprintf("DoubleMap{%s}", strings.Join(items, ", "))
}
