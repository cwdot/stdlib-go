package collections

import (
	"fmt"
	"strings"
	"sync"
)

func NewMapList[K comparable, T any]() *MapList[K, T] {
	counter := make(map[K][]T)
	return &MapList[K, T]{
		m: counter,
	}
}

type MapList[K comparable, T any] struct {
	mu sync.Mutex
	m  map[K][]T
}

// String debug
func (km *MapList[K, T]) String() string {
	km.mu.Lock()
	defer km.mu.Unlock()

	items := make([]string, 0, len(km.m))
	for k, v := range km.m {
		items = append(items, fmt.Sprintf("%v=%v", k, v))
	}

	return fmt.Sprintf("MapList: %s", strings.Join(items, ", "))
}

// Add saves item; creates path
func (km *MapList[K, T]) Add(key K, value T) bool {
	km.mu.Lock()
	defer km.mu.Unlock()

	if _, ok := km.m[key]; !ok {
		km.m[key] = make([]T, 0, 5)
	}
	km.m[key] = append(km.m[key], value)

	return true
}

// Get return Item by its id
func (km *MapList[K, T]) Get(parent K, index int) (T, bool, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	if outer, ok := km.m[parent]; ok {
		return outer[index], true, ok
	}
	var result T
	return result, false, false
}

// GetList return Item by its id
func (km *MapList[K, T]) GetList(key K) ([]T, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	if outer, ok := km.m[key]; ok {
		return outer, true
	}
	return nil, false
}

// Count the number of items in the list for provided key
func (km *MapList[K, T]) Count(key K) int {
	km.mu.Lock()
	defer km.mu.Unlock()

	if outer, ok := km.m[key]; ok {
		return len(outer)
	}
	return 0
}

// Size returns the total number of entries in the map
func (km *MapList[K, T]) Size() int {
	km.mu.Lock()
	defer km.mu.Unlock()

	return len(km.m)
}

// TotalSize returns the total number of tracked items across the map and sublists
func (km *MapList[K, T]) TotalSize() int {
	km.mu.Lock()
	defer km.mu.Unlock()

	count := 0
	for _, v := range km.m {
		count += len(v)
	}
	return count
}

// Copy creates a copy
func (km *MapList[K, T]) Copy() map[K][]T {
	km.mu.Lock()
	defer km.mu.Unlock()

	items := make(map[K][]T, len(km.m))
	for k, v := range km.m {
		items[k] = v
	}
	return items
}

// IterateLists iterate list keys
func (km *MapList[K, T]) IterateLists(fn func(K, []T) bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	for k, v := range km.m {
		if !fn(k, v) {
			break
		}
	}
}

// IterateItems iterate lists and their items
func (km *MapList[K, T]) IterateItems(fn func(K, T) bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	for k, v := range km.m {
		for _, item := range v {
			if !fn(k, item) {
				break
			}
		}
	}
}
